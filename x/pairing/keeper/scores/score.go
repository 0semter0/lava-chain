// Package scores implements the scoring mechanism used for picking providers in the pairing process.
//
// The pairing process involves the following steps:
// 1. Collect pairing requirements and strategy from the policy.
// 2. Generate pairing slots with requirements (one slot per provider).
// 3. Compute the pairing score of each provider with respect to each slot.
// 4. Pick a provider for each slot with a pseudo-random weighted choice.
//
// Pairing requirements describe the policy-imposed requirements for paired providers. Examples include
// geolocation constraints, ability to service archive requests, and expectations regarding QoS ranking
// of selected providers. Pairing requirements must satisfy the ScoreReq interface, whose methods can
// identify the ScoreReq (by name), and compute a score for a provider with respect to that requirement.
//
// A pairing slot represents a single provider slot in the pairing list (The number of slots for pairing
// is defined by the policy). Each pairing slot holds a set of pairing requirements (a pairing slot may
// repeat). For example, a policy may state that the pairing list has 6 slots, and providers should be
// located in Asia and Europe. This can be satisfied with a pairing list that has 3 (identical) slots
// that require providers in Asia and 3 (identical) slots that require providers in Europe.
//
// A pairing score describes the suitability of a provider for a pairing slot (under a given strategy).
// The score depends on the slot's requirements: for example, given a slot which requires geolocation in
// Asia, a provider in Asia will generally get higher score than one in Europe. The score is calculated for
// each <provider, slot> combination.
//
// A pairing score strategy defines the weight of each score requirement in the final score calculation
// for a <provider, slot> combination. For example, given a slot with several requirements, then the
// overall pairing score would be calculated as score1^w1 + score2^w2 + ... (where score1 is the score
// of the provider with respect to the first requirement, score2 with respect to the second requirement
// and so on).
//
//
// To add a new requirement, create an object implementing the ScoreReq interface and add the new requirement in GetAllReqs().

package scores

import (
	"bytes"
	"fmt"
	"strconv"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/rand"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	planstypes "github.com/lavanet/lava/x/plans/types"
)

var uniformStrategy ScoreStrategy

// TODO: currently we'll use weight=1 for all reqs. In the future, we'll get it from policy
func init() {
	reqs := GetAllReqs()

	// init strategy
	uniformStrategy = make(ScoreStrategy)
	for _, req := range reqs {
		uniformStrategy[req.GetName()] = 1
	}
}

func GetAllReqs() []ScoreReq {
	return []ScoreReq{
		&StakeReq{},
		&GeoReq{},
		&QosReq{},
	}
}

// get the overall requirements from the policy and assign slots that'll fulfil them
// TODO: this function should be changed in the future since it only supports stake reqs
func CalcSlots(policy planstypes.Policy) []*PairingSlot {
	// init slot array (should be as the number of providers to pair)
	slots := make([]*PairingSlot, policy.MaxProvidersToPair)

	reqs := GetAllReqs()
	for i := range slots {
		reqMap := make(map[string]ScoreReq)
		for _, req := range reqs {
			active := req.Init(policy)
			if active {
				reqMap[req.GetName()] = req.GetReqForSlot(policy, i)
			}
		}

		slots[i] = NewPairingSlot()
		slots[i].Reqs = reqMap
	}

	return slots
}

// group the slots
func GroupSlots(slots []*PairingSlot) []*PairingSlot {
	uniqueSlots := []*PairingSlot{}

	if len(slots) == 0 {
		panic("no pairing slots available")
	}

	for k := 0; k < len(slots); k++ {
		isUnique := true

		for i := range uniqueSlots {
			if slots[k].Equal(uniqueSlots[i]) {
				uniqueSlots[i].Count += 1
				isUnique = false
				break
			}
		}

		if isUnique {
			uniqueSlot := *slots[k]
			uniqueSlots = append(uniqueSlots, &uniqueSlot)
		}
	}

	return uniqueSlots
}

// TODO: currently we'll use weight=1 for all reqs. In the future, we'll get it from policy
func GetStrategy() ScoreStrategy {
	return uniformStrategy
}

// CalcPairingScore calculates the final pairing score for a pairing slot (with strategy)
// For efficiency purposes, we calculate the score on a diff slot which represents the diff reqs of the current slot
// and the previous slot
func CalcPairingScore(scores []*PairingScore, strategy ScoreStrategy, diffSlot *PairingSlot) error {
	// calculate the score for each req for each provider
	for _, score := range scores {
		for _, req := range diffSlot.Reqs {
			reqName := req.GetName()
			weight, ok := strategy[reqName]
			if !ok {
				return utils.LavaFormatError("req not found in strategy", fmt.Errorf("cannot calculate pairing score"),
					utils.Attribute{Key: "req", Value: reqName},
				)
			}

			newScoreComp := req.Score(*score.Provider)
			if newScoreComp == math.ZeroUint() {
				return utils.LavaFormatError("new score component is zero", fmt.Errorf("cannot calculate pairing score"),
					utils.Attribute{Key: "score component", Value: reqName},
					utils.Attribute{Key: "provider", Value: score.Provider.Address},
				)
			}
			newScoreCompDec := sdk.NewDecFromInt(math.Int(newScoreComp))
			newScoreCompDec = newScoreCompDec.Power(weight)
			newScoreComp = math.Uint(newScoreCompDec.TruncateInt())

			// update the score component map
			score.ScoreComponents[reqName] = newScoreComp
		}

		// calc new score
		newScore := math.OneUint()
		for _, scoreComp := range score.ScoreComponents {
			newScore = newScore.Mul(scoreComp)
		}
		score.Score = newScore
	}

	return nil
}

// PrepareHashData prepares the hash needed in the pseudo-random choice of providers
func PrepareHashData(projectIndex, chainID string, epochHash []byte, idx int) []byte {
	return bytes.Join([][]byte{epochHash, []byte(chainID), []byte(projectIndex), []byte(strconv.Itoa(idx))}, nil)
}

// PickProviders pick a <group-count> providers set with a pseudo-random weighted choice
// (using the providers' score list and hashData)
func PickProviders(ctx sdk.Context, scores []*PairingScore, groupCount int, hashData []byte) (returnedProviders []epochstoragetypes.StakeEntry) {
	if len(scores) == 0 {
		return returnedProviders
	}

	scoreSum := math.ZeroUint()
	for _, providerScore := range scores {
		if providerScore.SkipForSelection {
			// skip index of providers already selected
			continue
		}
		scoreSum = scoreSum.Add(providerScore.Score)
	}
	if scoreSum == math.ZeroUint() {
		utils.LavaFormatError("score sum is zero", fmt.Errorf("cannot pick providers for pairing"))
		return returnedProviders
	}

	rng := rand.New(hashData)

	for it := 0; it < groupCount; it++ {
		randomValue := uint64(rng.Int63n(scoreSum.BigInt().Int64())) + 1
		newScoreSum := math.ZeroUint()

		for idx := len(scores) - 1; idx >= 0; idx-- {
			if scores[idx].SkipForSelection {
				// skip index of providers already selected
				continue
			}
			providerScore := scores[idx]
			newScoreSum = newScoreSum.Add(providerScore.Score)
			if randomValue <= newScoreSum.Uint64() {
				// we hit our chosen provider
				// remove this provider from the random pool, so the sum is lower now
				returnedProviders = append(returnedProviders, *providerScore.Provider)
				scoreSum = scoreSum.Sub(providerScore.Score)
				scores[idx].SkipForSelection = true
				break
			}
		}
	}

	return returnedProviders
}
