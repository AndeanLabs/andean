package keeper

import (
	"context"

	"cosmossdk.io/errors"

	"andean/x/itzel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitPrice(goCtx context.Context, msg *types.MsgSubmitPrice) (*types.MsgSubmitPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is an authorized oracle
	params := k.GetParams(ctx)
	isAuthorized := false
	for _, addr := range params.AuthorizedOracles {
		if addr == msg.Creator {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "sender is not an authorized oracle")
	}

	report := types.PriceReport{
		Source:    msg.Source,
		Oracle:    msg.Creator,
		Price:     msg.Price,
		Timestamp: int32(ctx.BlockTime().Unix()),
	}

	k.SetPriceReport(ctx, report)

	return &types.MsgSubmitPriceResponse{}, nil
}
