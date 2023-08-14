package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lavanet/lava/utils"
	planstypes "github.com/lavanet/lava/x/plans/types"
)

func (k Keeper) EnforceClientCUsUsageInEpoch(ctx sdk.Context, allowedCU, totalCUInEpochForUserProvider uint64, clientAddr sdk.AccAddress, chainID string, epoch uint64) error {
	project, err := k.GetProjectData(ctx, clientAddr, chainID, epoch)
	// if client is not legacy (works through a project), the CU verification is different
	if err == nil {
		plan, err := k.subscriptionKeeper.GetPlanFromSubscription(ctx, project.GetSubscription())
		if err != nil {
			return err
		}

		planPolicy := plan.GetPlanPolicy()
		policies := []*planstypes.Policy{&planPolicy, project.AdminPolicy, project.SubscriptionPolicy}

		sub, found := k.subscriptionKeeper.GetSubscription(ctx, project.GetSubscription())
		if !found {
			return utils.LavaFormatError("can't find subscription", fmt.Errorf("EnforceClientCUsUsageInEpoch_cant_find_subscription"), utils.Attribute{Key: "subscriptionKey", Value: project.GetSubscription()})
		}

		if sub.GetMonthCuLeft() == 0 {
			return utils.LavaFormatError("total cu in epoch for consumer exceeded the amount of CU left in the subscription", fmt.Errorf("consumer CU limit exceeded for subscription"), []utils.Attribute{{Key: "subscriptionCuLeft", Value: sub.GetMonthCuLeft()}}...)
		}

		_, effectiveTotalCu := k.CalculateEffectiveAllowedCuPerEpochFromPolicies(policies, project.UsedCu, sub.MonthCuLeft)
		if !planstypes.VerifyTotalCuUsage(effectiveTotalCu, project.GetUsedCu()) {
			return utils.LavaFormatError("total cu in epoch for consumer exceeded the allowed amount for the project", fmt.Errorf("consumer CU limit exceeded for project"), []utils.Attribute{{Key: "projectUsedCu", Value: project.GetUsedCu()}}...)
		}
	}

	if totalCUInEpochForUserProvider > allowedCU {
		// if cu limit reached we return an error.
		return utils.LavaFormatError("total cu in epoch for consumer exceeded the allowed amount", fmt.Errorf("consumer CU limit exceeded"), []utils.Attribute{{Key: "totalCUInEpochForUserProvider", Value: totalCUInEpochForUserProvider}, {Key: "allowedCUProvider", Value: allowedCU}}...)
	}
	return nil
}
