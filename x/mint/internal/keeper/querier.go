package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	sdkerrors "github.com/DFWallet/anatha/types/errors"
	"github.com/DFWallet/project-anatha/x/mint/internal/types"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, _ abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParameters:
			return queryParams(ctx, k)
		case types.QueryMinter:
			return queryMinter(ctx, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryMinter(ctx sdk.Context, k Keeper) ([]byte, error) {
	minter := k.GetMinter(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, minter)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
