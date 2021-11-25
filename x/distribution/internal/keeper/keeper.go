package keeper

import (
	"fmt"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/params"
	"github.com/DFWallet/project-anatha/x/hra"
	"github.com/DFWallet/project-anatha/x/staking"
	"github.com/DFWallet/anatha/x/supply"
	"github.com/DFWallet/project-anatha/x/distribution/internal/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc          *codec.Codec
	storeKey     sdk.StoreKey
	paramSpace   params.Subspace
	supplyKeeper supply.Keeper
	stakingKeeper *staking.Keeper
	HraKeeper    *hra.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace, supplyKeeper supply.Keeper, stakingKeeper *staking.Keeper, hraKeeper *hra.Keeper) Keeper {
	AccountMustBePresent(&supplyKeeper, types.AmcModuleName)
	AccountMustBePresent(&supplyKeeper, types.NvrpModuleName)
	AccountMustBePresent(&supplyKeeper, types.HRAHolderRewardModuleName)
	AccountMustBePresent(&supplyKeeper, types.DevelopmentFundModuleName)
	AccountMustBePresent(&supplyKeeper, types.SecurityTokenFundModuleName)

	return Keeper{
		cdc:          cdc,
		storeKey:     key,
		paramSpace:   paramSpace.WithKeyTable(types.ParamKeyTable()),
		supplyKeeper: supplyKeeper,
		stakingKeeper: stakingKeeper,
		HraKeeper:    hraKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func AccountMustBePresent(supplyKeeper *supply.Keeper, accountName string) {
	if addr := supplyKeeper.GetModuleAddress(accountName); addr == nil {
		panic(fmt.Sprintf("the %s module account has not been set", accountName))
	}
}

