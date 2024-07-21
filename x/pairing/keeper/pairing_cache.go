package keeper

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
)

func (k Keeper) SetPairingCached(ctx sdk.Context, project string, chainID string, epoch uint64, pairedProviders []epochstoragetypes.StakeEntry) {
	if k.pairingCache == nil {
		// pairing cache is not initialized, will be in next epoch so simply skip
		return
	}
	k.pairingCache[pairingCacheKey(project, chainID, epoch)] = pairedProviders
}

func (k *Keeper) GetPairingCached(ctx sdk.Context, project string, chainID string, epoch uint64) ([]epochstoragetypes.StakeEntry, bool) {
	if k.pairingCache == nil {
		// pairing cache is not initialized, will be in next epoch so simply skip
		return nil, false
	}

	if providers, ok := k.pairingCache[pairingCacheKey(project, chainID, epoch)]; ok {
		return providers, true
	}

	return nil, false
}

func pairingCacheKey(project string, chainID string, epoch uint64) string {
	return strings.Join([]string{project, chainID, strconv.FormatUint(epoch, 10)}, " ")
}
