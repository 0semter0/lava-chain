package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/pairing/types"
)

// SetUniquePaymentStorageClientProvider set a specific uniquePaymentStorageClientProvider in the store from its index
func (k Keeper) SetUniquePaymentStorageClientProvider(ctx sdk.Context, uniquePaymentStorageClientProvider types.UniquePaymentStorageClientProvider) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniquePaymentStorageClientProviderKeyPrefix))
	b := k.cdc.MustMarshal(&uniquePaymentStorageClientProvider)
	store.Set(types.UniquePaymentStorageClientProviderKey(
		uniquePaymentStorageClientProvider.Index,
	), b)
}

// GetUniquePaymentStorageClientProvider returns a uniquePaymentStorageClientProvider from its index
func (k Keeper) GetUniquePaymentStorageClientProvider(
	ctx sdk.Context,
	index string,
) (val types.UniquePaymentStorageClientProvider, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniquePaymentStorageClientProviderKeyPrefix))

	b := store.Get(types.UniquePaymentStorageClientProviderKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniquePaymentStorageClientProvider removes a uniquePaymentStorageClientProvider from the store
func (k Keeper) RemoveUniquePaymentStorageClientProvider(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniquePaymentStorageClientProviderKeyPrefix))
	store.Delete(types.UniquePaymentStorageClientProviderKey(
		index,
	))
}

// GetAllUniquePaymentStorageClientProvider returns all uniquePaymentStorageClientProvider
func (k Keeper) GetAllUniquePaymentStorageClientProvider(ctx sdk.Context) (list []types.UniquePaymentStorageClientProvider) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniquePaymentStorageClientProviderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UniquePaymentStorageClientProvider
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) AddUniquePaymentStorageClientProvider(ctx sdk.Context, chainID string, block uint64, projectID string, providerAddress sdk.AccAddress, uniqueIdentifier string, usedCU uint64) *types.UniquePaymentStorageClientProvider {
	key := k.EncodeUniquePaymentKey(ctx, projectID, providerAddress, uniqueIdentifier, chainID)
	entry := types.UniquePaymentStorageClientProvider{Index: key, Block: block, UsedCU: usedCU}
	k.SetUniquePaymentStorageClientProvider(ctx, entry)
	return &entry
}

func (k Keeper) IsDoubleSpend(ctx sdk.Context, chainID string, block uint64, projectID string, providerAddress sdk.AccAddress, uniqueIdentifier string) bool {
	key := k.EncodeUniquePaymentKey(ctx, projectID, providerAddress, uniqueIdentifier, chainID)
	_, found := k.GetUniquePaymentStorageClientProvider(ctx, key)
	return found
}

func (k Keeper) GetConsumerFromUniquePayment(uniquePaymentStorageClientProvider string) string {
	key := uniquePaymentStorageClientProvider
	providerAdrLengh := charToAsciiNumber(rune(key[0]))
	provider := key[1 : providerAdrLengh+1]
	return provider
}

func (k Keeper) EncodeUniquePaymentKey(ctx sdk.Context, projectID string, providerAddress sdk.AccAddress, uniqueIdentifier, chainID string) string {
	leadingChar := asciiNumberToChar(len(projectID))
	key := string(leadingChar) + projectID + providerAddress.String() + uniqueIdentifier + chainID
	return key
}

func charToAsciiNumber(char rune) int {
	return int(char)
}

func asciiNumberToChar(asciiNum int) rune {
	return rune(asciiNum)
}
