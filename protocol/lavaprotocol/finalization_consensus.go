package lavaprotocol

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lavanet/lava/protocol/chainlib"
	"github.com/lavanet/lava/utils"
	conflicttypes "github.com/lavanet/lava/x/conflict/types"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	"golang.org/x/exp/slices"
)

type FinalizationConsensus struct {
	currentProviderHashesConsensus   []ProviderHashesConsensus
	prevEpochProviderHashesConsensus []ProviderHashesConsensus
	providerDataContainersMu         sync.RWMutex
	currentEpoch                     uint64
	latestBlock                      uint64 // for caching
}

type ProviderHashesConsensus struct {
	FinalizedBlocksHashes map[int64]string
	agreeingProviders     map[string]providerDataContainer
}

type providerDataContainer struct {
	LatestFinalizedBlock  int64
	LatestBlockTime       time.Time
	FinalizedBlocksHashes map[int64]string
	SigBlocks             []byte
	SessionId             uint64
	BlockHeight           int64
	RelayNum              uint64
	LatestBlock           int64
	// TODO:: keep relay request for conflict reporting
}

func GetLatestFinalizedBlock(latestBlock, blockDistanceForFinalizedData int64) int64 {
	finalization_criteria := blockDistanceForFinalizedData
	return latestBlock - finalization_criteria
}

func (fc *FinalizationConsensus) newProviderHashesConsensus(blockDistanceForFinalizedData int64, providerAcc string, latestBlock int64, finalizedBlocks map[int64]string, reply *pairingtypes.RelayReply, req *pairingtypes.RelaySession) ProviderHashesConsensus {
	newProviderDataContainer := providerDataContainer{
		LatestFinalizedBlock:  GetLatestFinalizedBlock(latestBlock, blockDistanceForFinalizedData),
		LatestBlockTime:       time.Now(),
		FinalizedBlocksHashes: finalizedBlocks,
		SigBlocks:             reply.SigBlocks,
		SessionId:             req.SessionId,
		RelayNum:              req.RelayNum,
		BlockHeight:           req.Epoch,
		LatestBlock:           latestBlock,
	}
	providerDataContainers := map[string]providerDataContainer{}
	providerDataContainers[providerAcc] = newProviderDataContainer
	return ProviderHashesConsensus{
		FinalizedBlocksHashes: finalizedBlocks,
		agreeingProviders:     providerDataContainers,
	}
}

func (fc *FinalizationConsensus) insertProviderToConsensus(blockDistanceForFinalizedData int64, consensus *ProviderHashesConsensus, finalizedBlocks map[int64]string, latestBlock int64, reply *pairingtypes.RelayReply, req *pairingtypes.RelaySession, providerAcc string) {
	newProviderDataContainer := providerDataContainer{
		LatestFinalizedBlock:  GetLatestFinalizedBlock(latestBlock, blockDistanceForFinalizedData),
		LatestBlockTime:       time.Now(),
		FinalizedBlocksHashes: finalizedBlocks,
		SigBlocks:             reply.SigBlocks,
		SessionId:             req.SessionId,
		RelayNum:              req.RelayNum,
		BlockHeight:           req.Epoch,
		LatestBlock:           latestBlock,
	}
	consensus.agreeingProviders[providerAcc] = newProviderDataContainer

	for blockNum, blockHash := range finalizedBlocks {
		consensus.FinalizedBlocksHashes[blockNum] = blockHash
	}
}

// Compare finalized block hashes with previous providers
// Looks for discrepancy with current epoch providers
// if no conflicts, insert into consensus and break
// create new consensus group if no consensus matched
// check for discrepancy with old epoch
// checks if there is a consensus mismatch between hashes provided by different providers
func (fc *FinalizationConsensus) UpdateFinalizedHashes(blockDistanceForFinalizedData int64, providerAddress string, finalizedBlocks map[int64]string, req *pairingtypes.RelaySession, reply *pairingtypes.RelayReply) (finalizationConflict *conflicttypes.FinalizationConflict, err error) {
	latestBlock := reply.LatestBlock
	fc.providerDataContainersMu.Lock()
	defer fc.providerDataContainersMu.Unlock()

	if len(fc.currentProviderHashesConsensus) == 0 && len(fc.prevEpochProviderHashesConsensus) == 0 {
		newHashConsensus := fc.newProviderHashesConsensus(blockDistanceForFinalizedData, providerAddress, latestBlock, finalizedBlocks, reply, req)
		fc.currentProviderHashesConsensus = append(make([]ProviderHashesConsensus, 0), newHashConsensus)
	} else {
		inserted := false
		// Looks for discrepancy with current epoch providers
		// go over all consensus groups, if there is a mismatch add it as a consensus group and send a conflict
		for _, consensus := range fc.currentProviderHashesConsensus {
			err := fc.discrepancyChecker(finalizedBlocks, consensus)
			if err != nil {
				// TODO: bring the other data as proof
				finalizationConflict = &conflicttypes.FinalizationConflict{RelayReply0: reply}
				// we need to insert into a new consensus group before returning
				// or create new consensus group if no consensus matched
				continue
			}

			if !inserted {
				// if no discrepency with this group and not inserted yet -> insert into consensus
				fc.insertProviderToConsensus(blockDistanceForFinalizedData, &consensus, finalizedBlocks, latestBlock, reply, req, providerAddress)
				inserted = true
			}
			// keep comparing with other groups, if there is a new message with a conflict we need to report it too
		}
		if !inserted {
			// means there was a consensus mismatch with everything, so it wasn't inserted and we add it here
			newHashConsensus := fc.newProviderHashesConsensus(blockDistanceForFinalizedData, providerAddress, latestBlock, finalizedBlocks, reply, req)
			fc.currentProviderHashesConsensus = append(fc.currentProviderHashesConsensus, newHashConsensus)
		}
		if finalizationConflict != nil {
			// means there was a conflict and we need to report
			return finalizationConflict, utils.LavaFormatError("Simulation: Conflict found in discrepancyChecker", err)
		}

		// check for discrepancy with old epoch
		for idx, consensus := range fc.prevEpochProviderHashesConsensus {
			err := fc.discrepancyChecker(finalizedBlocks, consensus)
			if err != nil {
				// TODO: bring the other data as proof
				finalizationConflict = &conflicttypes.FinalizationConflict{RelayReply0: reply}
				return finalizationConflict, utils.LavaFormatError("Simulation: prev epoch Conflict found in discrepancyChecker", err, utils.Attribute{Key: "Consensus idx", Value: strconv.Itoa(idx)}, utils.Attribute{Key: "provider", Value: providerAddress})
			}
		}
	}
	if debug {
		utils.LavaFormatDebug("finalization information update successfully", utils.Attribute{Key: "finalization data", Value: finalizedBlocks}, utils.Attribute{Key: "currentProviderHashesConsensus", Value: fc.currentProviderHashesConsensus}, utils.Attribute{Key: "currentProviderHashesConsensus", Value: fc.currentProviderHashesConsensus})
	}
	return finalizationConflict, nil
}

func (fc *FinalizationConsensus) discrepancyChecker(finalizedBlocksA map[int64]string, consensus ProviderHashesConsensus) (errRet error) {
	var toIterate map[int64]string   // the smaller map between the two to compare
	var otherBlocks map[int64]string // the other map

	if len(finalizedBlocksA) < len(consensus.FinalizedBlocksHashes) {
		toIterate = finalizedBlocksA
		otherBlocks = consensus.FinalizedBlocksHashes
	} else {
		toIterate = consensus.FinalizedBlocksHashes
		otherBlocks = finalizedBlocksA
	}
	// Iterate over smaller array, looks for mismatching hashes between the inputs
	for blockNum, blockHash := range toIterate {
		if otherHash, ok := otherBlocks[blockNum]; ok {
			if blockHash != otherHash {
				// TODO: gather discrepancy data
				return utils.LavaFormatError("Simulation: reliability discrepancy, different hashes detected for block", HashesConsunsusError, utils.Attribute{Key: "blockNum", Value: blockNum}, utils.Attribute{Key: "Hashes", Value: fmt.Sprintf("%s vs %s", blockHash, otherHash)}, utils.Attribute{Key: "toIterate", Value: toIterate}, utils.Attribute{Key: "otherBlocks", Value: otherBlocks})
			}
		}
	}

	return nil
}

func (fc *FinalizationConsensus) NewEpoch(epoch uint64) {
	fc.providerDataContainersMu.Lock()
	defer fc.providerDataContainersMu.Unlock()

	if fc.currentEpoch < epoch {
		// means it's time to refresh the epoch
		fc.prevEpochProviderHashesConsensus = fc.currentProviderHashesConsensus
		fc.currentProviderHashesConsensus = []ProviderHashesConsensus{}
		fc.currentEpoch = epoch
	}
}

func (s *FinalizationConsensus) LatestBlock() uint64 {
	s.providerDataContainersMu.RLock()
	defer s.providerDataContainersMu.RUnlock()
	return s.latestBlock
}

// returns the expected latest block to be at based on the current finalization data, and the number of providers we have information for
// does the calculation on finalized entries then extrapolates the ending based on blockDistance
func (s *FinalizationConsensus) ExpectedBlockHeight(chainParser chainlib.ChainParser) (expectedBlockHeight int64, numOfProviders int) {
	s.providerDataContainersMu.RLock()
	defer s.providerDataContainersMu.RUnlock()
	allowedBlockLagForQosSync, averageBlockTime, blockDistanceForFinalizedData, _ := chainParser.ChainBlockStats()
	averageBlockTime_ms := averageBlockTime

	var highestBlockNumber int64 = 0
	FindHighestBlockNumber := func(listProviderHashesConsensus []ProviderHashesConsensus) int64 {
		for _, providerHashesConsensus := range listProviderHashesConsensus {
			for _, providerDataContainer := range providerHashesConsensus.agreeingProviders {
				if highestBlockNumber < providerDataContainer.LatestFinalizedBlock {
					highestBlockNumber = providerDataContainer.LatestFinalizedBlock
				}
			}
		}
		return highestBlockNumber
	}
	highestBlockNumber = FindHighestBlockNumber(s.prevEpochProviderHashesConsensus) // update the highest in place
	highestBlockNumber = FindHighestBlockNumber(s.currentProviderHashesConsensus)

	now := time.Now()
	calcExpectedBlocks := func(mapExpectedBlockHeights map[string]int64, listProviderHashesConsensus []ProviderHashesConsensus) map[string]int64 {
		for _, providerHashesConsensus := range listProviderHashesConsensus {
			for providerAddress, providerDataContainer := range providerHashesConsensus.agreeingProviders {
				interpolation := InterpolateBlocks(now, providerDataContainer.LatestBlockTime, averageBlockTime_ms)
				expected := providerDataContainer.LatestFinalizedBlock + interpolation
				// limit the interpolation to the highest seen block height
				if expected > highestBlockNumber {
					expected = highestBlockNumber
				}
				mapExpectedBlockHeights[providerAddress] = expected
			}
		}
		return mapExpectedBlockHeights
	}
	mapExpectedBlockHeights := map[string]int64{}
	// prev must be before current because we overwrite
	mapExpectedBlockHeights = calcExpectedBlocks(mapExpectedBlockHeights, s.prevEpochProviderHashesConsensus)
	mapExpectedBlockHeights = calcExpectedBlocks(mapExpectedBlockHeights, s.currentProviderHashesConsensus)

	median := func(dataMap map[string]int64) int64 {
		data := make([]int64, len(dataMap))
		i := 0
		for _, latestBlock := range dataMap {
			data[i] = latestBlock
			i++
		}
		slices.Sort(data)

		var median int64
		data_len := len(data)
		if data_len == 0 {
			return 0
		} else if data_len%2 == 0 {
			median = ((data[data_len/2-1] + data[data_len/2]) / 2.0)
		} else {
			median = data[data_len/2]
		}
		return median
	}
	medianOfExpectedBlocks := median(mapExpectedBlockHeights)
	providersMedianOfLatestBlock := medianOfExpectedBlocks + int64(blockDistanceForFinalizedData)
	if debug {
		utils.LavaFormatDebug("finalization information", utils.Attribute{Key: "mapExpectedBlockHeights", Value: mapExpectedBlockHeights}, utils.Attribute{Key: "medianOfExpectedBlocks", Value: medianOfExpectedBlocks}, utils.Attribute{Key: "latestBlock", Value: providersMedianOfLatestBlock}, utils.Attribute{Key: "providersMedianOfLatestBlock", Value: providersMedianOfLatestBlock})
	}
	if medianOfExpectedBlocks > 0 && uint64(providersMedianOfLatestBlock) > s.latestBlock {
		if uint64(providersMedianOfLatestBlock) > s.latestBlock+1000 && s.latestBlock > 0 {
			utils.LavaFormatError("uncontinuous jump in finalization data", nil, utils.Attribute{Key: "latestBlock", Value: s.latestBlock}, utils.Attribute{Key: "providersMedianOfLatestBlock", Value: providersMedianOfLatestBlock})
		}
		atomic.StoreUint64(&s.latestBlock, uint64(providersMedianOfLatestBlock)) // we can only set conflict to "reported".
	}

	// median of all latest blocks after interpolation minus allowedBlockLagForQosSync is the lowest block in the finalization proof
	// then we move forward blockDistanceForFinalizedData to get the expected latest block
	return providersMedianOfLatestBlock - allowedBlockLagForQosSync, len(mapExpectedBlockHeights)
}

func InterpolateBlocks(timeNow, latestBlockTime time.Time, averageBlockTime time.Duration) int64 {
	if timeNow.Before(latestBlockTime) {
		return 0
	}
	return int64(timeNow.Sub(latestBlockTime) / averageBlockTime) // interpolation
}
