package keeper

import (
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/project-anatha/x/staking/types"
)

// Return all validators that a delegator is bonded to. If maxRetrieve is supplied, the respective amount will be returned.
func (k Keeper) GetDelegatorValidators(ctx sdk.Context, delegatorAddr sdk.AccAddress,
	maxRetrieve uint16) (validators []types.Validator) {
	validators = make([]types.Validator, maxRetrieve)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetDelegationsKey(delegatorAddr)
	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey) // smallest to largest
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())

		validator, found := k.GetValidator(ctx, delegation.ValidatorAddress)
		if !found {
			panic(types.ErrNoValidatorFound)
		}
		validators[i] = validator
		i++
	}
	return validators[:i] // trim
}

// return a validator that a delegator is bonded to
func (k Keeper) GetDelegatorValidator(ctx sdk.Context, delegatorAddr sdk.AccAddress,
	validatorAddr sdk.ValAddress) (validator types.Validator, err error) {

	delegation, found := k.GetDelegation(ctx, delegatorAddr, validatorAddr)
	if !found {
		return validator, types.ErrNoDelegation
	}

	validator, found = k.GetValidator(ctx, delegation.ValidatorAddress)
	if !found {
		panic(types.ErrNoValidatorFound)
	}
	return
}

//_____________________________________________________________________________________

// return all delegations for a delegator
func (k Keeper) GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []types.Delegation {
	delegations := make([]types.Delegation, 0)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetDelegationsKey(delegator)
	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey) //smallest to largest
	defer iterator.Close()

	i := 0
	for ; iterator.Valid(); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		delegations = append(delegations, delegation)
		i++
	}

	return delegations
}

// return all unbonding-delegations for a delegator
func (k Keeper) GetAllUnbondingDelegations(ctx sdk.Context, delegator sdk.AccAddress) []types.UnbondingDelegation {
	unbondingDelegations := make([]types.UnbondingDelegation, 0)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetUBDsKey(delegator)
	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey) // smallest to largest
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		unbondingDelegation := types.MustUnmarshalUBD(k.cdc, iterator.Value())
		unbondingDelegations = append(unbondingDelegations, unbondingDelegation)
		i++
	}

	return unbondingDelegations
}
