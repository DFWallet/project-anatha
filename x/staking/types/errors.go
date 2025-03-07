package types

import (
	sdkerrors "github.com/DFWallet/anatha/types/errors"
)

// x/staking module sentinel errors
//
// TODO: Many of these errors are redundant. They should be removed and replaced
// by sdkerrors.ErrInvalidRequest.
//
// REF: https://github.com/DFWallet/anatha/issues/5450
var (
	ErrEmptyValidatorAddr              = sdkerrors.Register(ModuleName, 1, "empty validator address")
	ErrBadValidatorAddr                = sdkerrors.Register(ModuleName, 2, "validator address is invalid")
	ErrNoValidatorFound                = sdkerrors.Register(ModuleName, 3, "validator does not exist")
	ErrValidatorOwnerExists            = sdkerrors.Register(ModuleName, 4, "validator already exist for this operator address; must use new validator operator address")
	ErrValidatorPubKeyExists           = sdkerrors.Register(ModuleName, 5, "validator already exist for this pubkey; must use new validator pubkey")
	ErrValidatorPubKeyTypeNotSupported = sdkerrors.Register(ModuleName, 6, "validator pubkey type is not supported")
	ErrValidatorJailed                 = sdkerrors.Register(ModuleName, 7, "validator for this address is currently jailed")
	ErrBadRemoveValidator              = sdkerrors.Register(ModuleName, 8, "failed to remove validator")
	ErrSelfDelegationBelowMinimum      = sdkerrors.Register(ModuleName, 16, "validator's self delegation must be greater than their minimum self delegation")
	ErrMinSelfDelegationInvalid        = sdkerrors.Register(ModuleName, 17, "minimum self delegation must be a positive integer")
	ErrMinSelfDelegationDecreased      = sdkerrors.Register(ModuleName, 18, "minimum self delegation cannot be decrease")
	ErrEmptyDelegatorAddr              = sdkerrors.Register(ModuleName, 19, "empty delegator address")
	ErrBadDenom                        = sdkerrors.Register(ModuleName, 20, "invalid coin denomination")
	ErrBadDelegationAddr               = sdkerrors.Register(ModuleName, 21, "invalid address for (address, validator) tuple")
	ErrBadDelegationAmount             = sdkerrors.Register(ModuleName, 22, "invalid delegation amount")
	ErrNoDelegation                    = sdkerrors.Register(ModuleName, 23, "no delegation for (address, validator) tuple")
	ErrBadDelegatorAddr                = sdkerrors.Register(ModuleName, 24, "delegator does not exist with address")
	ErrNoDelegatorForAddress           = sdkerrors.Register(ModuleName, 25, "delegator does not contain delegation")
	ErrInsufficientShares              = sdkerrors.Register(ModuleName, 26, "insufficient delegation shares")
	ErrDelegationValidatorEmpty        = sdkerrors.Register(ModuleName, 27, "cannot delegate to an empty validator")
	ErrNotEnoughDelegationShares       = sdkerrors.Register(ModuleName, 28, "not enough delegation shares")
	ErrBadSharesAmount                 = sdkerrors.Register(ModuleName, 29, "invalid shares amount")
	ErrBadSharesPercent                = sdkerrors.Register(ModuleName, 30, "Invalid shares percent")
	ErrNotMature                       = sdkerrors.Register(ModuleName, 31, "entry not mature")
	ErrNoUnbondingDelegation           = sdkerrors.Register(ModuleName, 32, "no unbonding delegation found")
	ErrMaxUnbondingDelegationEntries   = sdkerrors.Register(ModuleName, 33, "too many unbonding delegation entries for (delegator, validator) tuple")
	ErrDelegatorShareExRateInvalid     = sdkerrors.Register(ModuleName, 41, "cannot delegate to validators with invalid (zero) ex-rate")
	ErrBothShareMsgsGiven              = sdkerrors.Register(ModuleName, 42, "both shares amount and shares percent provided")
	ErrNeitherShareMsgsGiven           = sdkerrors.Register(ModuleName, 43, "neither shares amount nor shares percent provided")
	ErrInvalidHistoricalInfo           = sdkerrors.Register(ModuleName, 44, "invalid historical info")
	ErrNoHistoricalInfo                = sdkerrors.Register(ModuleName, 45, "no historical info found")
	ErrInvalidGenesis                  = sdkerrors.Register(ModuleName, 46, "invalid genesis state")
)
