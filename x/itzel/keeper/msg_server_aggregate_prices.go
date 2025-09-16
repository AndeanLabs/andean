package keeper

import (
	"context"

	"cosmossdk.io/errors"

	"andean/x/itzel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AggregatePrices(goCtx context.Context, msg *types.MsgAggregatePrices) (*types.MsgAggregatePricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get all price reports for the source
	reports := k.GetPriceReportsBySource(ctx, msg.Source)
	if len(reports) == 0 {
		return nil, errors.Wrap(sdkerrors.ErrNotFound, "no price reports found for this source")
	}

	// Calculate the average price
	var totalPrice int32
	for _, report := range reports {
		totalPrice += report.Price
	}
	averagePrice := totalPrice / int32(len(reports))

	// Create and save the aggregated price
	aggregatedPrice := types.AggregatedPrice{
		Source:    msg.Source,
		Price:     averagePrice,
		Timestamp: int32(ctx.BlockTime().Unix()),
	}

	k.SetAggregatedPrice(ctx, aggregatedPrice)

	return &types.MsgAggregatePricesResponse{}, nil
}
