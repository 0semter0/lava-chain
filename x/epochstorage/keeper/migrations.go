package keeper

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/epochstorage/types"
	v3 "github.com/lavanet/lava/x/epochstorage/types/migrations/v3"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate7to8 transfers all the stake entries from the old stake storage to the new stake entries store
// StakeStorage is set to the stake entries store
// StakeStorageCurrent is set to the stake entries current store
// StakeStorageUnstake is deleted
func (m Migrator) Migrate7to8(ctx sdk.Context) error {
	utils.LavaFormatDebug("migrate: epochstorage to move stake entries from stake storage")
	k := m.keeper

	store := prefix.NewStore(ctx.KVStore(k.storeKey), v3.KeyPrefix(v3.StakeStorageKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())

		// identify stake storage type: regular, current or unstake
		if key == v3.StakeStorageKeyUnstakeConst {
			store.Delete(iterator.Key())
			continue
		}

		epoch, err := extractEpochFromStakeStorageKey(key)
		isCurrentStakeStorage := m.isCurrentStakeStorageKey(ctx, key)
		if err != nil && !isCurrentStakeStorage {
			panic(fmt.Errorf("stake storage with unidentified index %s", key))
		}

		var stakeStorage types.StakeStorage
		k.cdc.MustUnmarshal(iterator.Value(), &stakeStorage)

		for _, entry := range stakeStorage.StakeEntries {
			if isCurrentStakeStorage {
				k.SetStakeEntryCurrent(ctx, entry)
			} else {
				k.SetStakeEntry(ctx, epoch, entry)
			}
		}

		if !isCurrentStakeStorage {
			m.SetEpochHashForMigrator(ctx, epoch, stakeStorage.EpochBlockHash)
		}

		store.Delete(iterator.Key())
	}

	return nil
}

// the legacy StakeStorage store used keys that were built like this:
// strconv.FormatUint(block, 10) + chainID
func extractEpochFromStakeStorageKey(key string) (uint64, error) {
	number := ""
	for _, char := range key {
		if !unicode.IsDigit(char) {
			break
		}
		number += string(char)
	}
	// Attempt conversion to uint64, return 0 and false if fails
	parsedUint, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsedUint, nil
}

func (m Migrator) isCurrentStakeStorageKey(ctx sdk.Context, key string) bool {
	// the legacy StakeStorage key (both regular and current) had a "/" which should be cut off
	key, found := strings.CutSuffix(key, "/")
	if !found {
		return false
	}
	_, found, _ = m.keeper.specKeeper.IsSpecFoundAndActive(ctx, key)
	return found
}

func (m Migrator) SetEpochHashForMigrator(ctx sdk.Context, epoch uint64, hash []byte) {
	store := prefix.NewStore(ctx.KVStore(m.keeper.storeKey), types.KeyPrefix(types.EpochHashPrefix))
	store.Set(utils.SerializeBigEndian(epoch), hash)
}
