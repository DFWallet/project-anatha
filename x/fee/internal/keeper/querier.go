package keeper

import (
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	sdkerrors "github.com/DFWallet/anatha/types/errors"
	"github.com/DFWallet/project-anatha/x/fee/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryParameters          = "parameters"
	QueryFeeExcludedMessages = "excluded-messages"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryParameters:
			return queryParams(ctx, k)

		case QueryFeeExcludedMessages:
			return queryFeeExcludedMessages(ctx, k)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown distribution query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryFeeExcludedMessages(ctx sdk.Context, k Keeper) ([]byte, error) {
	feeExcludedMessages := k.GetFeeExcludedMessages(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, feeExcludedMessages)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
