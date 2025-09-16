package keeper

import (
	"context"

	"andean/x/itzel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitPrice(goCtx context.Context, msg *types.MsgSubmitPrice) (*types.MsgSubmitPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	price := types.Price{
		Source:    msg.Source,
		Value:     msg.Price,
		Timestamp: int32(ctx.BlockTime().Unix()),
	}

	k.SetPrice(ctx, price)

	return &types.MsgSubmitPriceResponse{}, nil
}
