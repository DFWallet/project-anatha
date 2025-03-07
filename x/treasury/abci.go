package treasury

import (
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/project-anatha/x/treasury/internal/types"
)


func BeginBlocker(ctx sdk.Context, k Keeper) {}

func EndBlocker(ctx sdk.Context, k Keeper) {
	 k.IterateScheduledDisbursementQueue(ctx, ctx.BlockTime(), func(disbursement types.Disbursement) (stop bool) {

	 	_, fromBuyBack, fromTreasury := k.CalculatePinAmountExtended(ctx, disbursement.Amount)

		err := k.DisburseFunds(ctx, disbursement.Operator, disbursement.Recipient, disbursement.Amount, fromBuyBack, fromTreasury)
		if err != nil {
			k.Logger(ctx).Info(err.Error())
		}

		k.RemoveFromDisbursementQueue(ctx,disbursement.Recipient, disbursement.ScheduledFor)

	 	return false
	 })
}
