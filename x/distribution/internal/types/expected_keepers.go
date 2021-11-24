package types

import sdk "github.com/DFWallet/anatha/types"

type NameHooks interface {
	AfterFirstNameCreated(ctx sdk.Context, address sdk.AccAddress) error
	AfterLastNameRemoved(ctx sdk.Context, address sdk.AccAddress) error
}
