package statetracker

import (
	"context"
	"fmt"
	"sync"

	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/lavanet/lava/protocol/rpcprovider/reliabilitymanager"
	"github.com/lavanet/lava/protocol/rpcprovider/rewardserver"
	"github.com/lavanet/lava/utils"
	conflicttypes "github.com/lavanet/lava/x/conflict/types"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

const (
	debug = false
)

type EventTracker struct {
	lock               sync.RWMutex
	clientCtx          client.Context
	blockResults       *ctypes.ResultBlockResults
	latestUpdatedBlock int64
}

func (et *EventTracker) updateBlockResults(latestBlock int64) (err error) {
	ctx := context.Background()
	var blockResults *ctypes.ResultBlockResults
	if latestBlock == 0 {
		res, err := et.clientCtx.Client.Status(ctx)
		if err != nil {
			return utils.LavaFormatWarning("could not get latest block height and requested latestBlock = 0", err)
		}
		latestBlock = res.SyncInfo.LatestBlockHeight
	}
	brp, err := tryIntoTendermintRPC(et.clientCtx.Client)
	if err != nil {
		return utils.LavaFormatError("could not get block result provider", err)
	}
	blockResults, err = brp.BlockResults(ctx, &latestBlock)
	if err != nil {
		return err
	}
	// lock for update after successful block result query
	et.lock.Lock()
	defer et.lock.Unlock()
	et.latestUpdatedBlock = latestBlock
	et.blockResults = blockResults
	return nil
}

func (et *EventTracker) getLatestPaymentEvents() (payments []*rewardserver.PaymentRequest, err error) {
	et.lock.RLock()
	defer et.lock.RUnlock()
	transactionResults := et.blockResults.TxsResults
	for _, tx := range transactionResults {
		events := tx.Events
		for _, event := range events {
			if event.Type == utils.EventPrefix+pairingtypes.RelayPaymentEventName {
				paymentList, err := rewardserver.BuildPaymentFromRelayPaymentEvent(event, et.latestUpdatedBlock)
				if err != nil {
					return nil, utils.LavaFormatError("failed relay_payment_event parsing", err, utils.Attribute{Key: "event", Value: event})
				}
				if debug {
					utils.LavaFormatDebug("relay_payment_event", utils.Attribute{Key: "payment", Value: paymentList})
				}
				payments = append(payments, paymentList...)
			}
		}
	}
	return payments, nil
}

func (et *EventTracker) getLatestVersionEvents() (updated bool) {
	et.lock.RLock()
	defer et.lock.RUnlock()
	for _, event := range et.blockResults.EndBlockEvents {
		if event.Type == utils.EventPrefix+"param_change" {
			for _, attribute := range event.Attributes {
				if attribute.Key == "param" && attribute.Value == "Version" {
					return true
				}
			}
		}
	}
	return false
}

func (et *EventTracker) getLatestSpecModifyEvents() (updated bool) {
	// SpecModifyEventName
	et.lock.RLock()
	defer et.lock.RUnlock()
	for _, event := range et.blockResults.EndBlockEvents {
		if event.Type == utils.EventPrefix+spectypes.SpecModifyEventName {
			return true
		}
	}
	return
}

func (et *EventTracker) getLatestVoteEvents() (votes []*reliabilitymanager.VoteParams, err error) {
	et.lock.RLock()
	defer et.lock.RUnlock()

	transactionResults := et.blockResults.TxsResults
	for _, tx := range transactionResults {
		events := tx.Events
		for _, event := range events {
			if event.Type == utils.EventPrefix+conflicttypes.ConflictVoteDetectionEventName {
				vote, err := reliabilitymanager.BuildVoteParamsFromDetectionEvent(event)
				if err != nil {
					return nil, utils.LavaFormatError("failed conflict_vote_detection_event parsing", err, utils.Attribute{Key: "event", Value: event})
				}
				utils.LavaFormatDebug("conflict_vote_detection_event", utils.Attribute{Key: "voteID", Value: vote.VoteID})
				votes = append(votes, vote)
			}
		}
	}

	beginBlockEvents := et.blockResults.BeginBlockEvents
	for _, event := range beginBlockEvents {
		if event.Type == utils.EventPrefix+conflicttypes.ConflictVoteRevealEventName {
			voteID, voteDeadline, err := reliabilitymanager.BuildBaseVoteDataFromEvent(event)
			if err != nil {
				return nil, utils.LavaFormatError("failed conflict_vote_reveal_event parsing", err, utils.Attribute{Key: "event", Value: event})
			}
			vote_reveal := &reliabilitymanager.VoteParams{VoteID: voteID, VoteDeadline: voteDeadline, ParamsType: reliabilitymanager.RevealVoteType}
			utils.LavaFormatDebug("conflict_vote_reveal_event", utils.Attribute{Key: "voteID", Value: voteID})
			votes = append(votes, vote_reveal)
		}
		if event.Type == utils.EventPrefix+conflicttypes.ConflictVoteResolvedEventName {
			voteID, _, err := reliabilitymanager.BuildBaseVoteDataFromEvent(event)
			if err != nil {
				if !reliabilitymanager.NoVoteDeadline.Is(err) {
					return nil, utils.LavaFormatError("failed conflict_vote_resolved_event parsing", err, utils.Attribute{Key: "event", Value: event})
				}
			}
			vote_resolved := &reliabilitymanager.VoteParams{VoteID: voteID, VoteDeadline: 0, ParamsType: reliabilitymanager.CloseVoteType, CloseVote: true}
			votes = append(votes, vote_resolved)
			utils.LavaFormatDebug("conflict_vote_resolved_event", utils.Attribute{Key: "voteID", Value: voteID})
		}
	}

	return votes, err
}

type tendermintRPC interface {
	BlockResults(
		ctx context.Context,
		height *int64,
	) (*ctypes.ResultBlockResults, error)
	ConsensusParams(
		ctx context.Context,
		height *int64,
	) (*ctypes.ResultConsensusParams, error)
}

func tryIntoTendermintRPC(cl client.TendermintRPC) (tendermintRPC, error) {
	brp, ok := cl.(tendermintRPC)
	if !ok {
		return nil, fmt.Errorf("client does not implement tendermintRPC: %T", cl)
	}
	return brp, nil
}
