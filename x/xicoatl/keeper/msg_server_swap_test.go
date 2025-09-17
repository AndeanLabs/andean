package keeper_test

import (
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/testutil/sample"
	itzeltypes "andean/x/itzel/types"
	"andean/x/xicoatl/keeper"
	"andean/x/xicoatl/types"
)

func TestSwapDynamicFee(t *testing.T) {
	// Common setup
	creator := sample.AccAddress()
	poolAmount := uint64(1000)
	swapAmount := uint64(100)
	baseFee := math.LegacyNewDecWithPrec(1, 3) // 0.1%
	feeMultiplier := math.LegacyNewDecWithPrec(5, 1) // 0.5

	// Test cases
	tests := []struct {
		name         string
		mockOracle   bool
		oraclePrice  int32
		expectedLog  string
	}{
		{
			name:         "oracle price found - pool price is aligned",
			mockOracle:   true,
			oraclePrice:  1,
			expectedLog:  "Oracle price found, applying dynamic fee",
		},
		{
			name:         "oracle price found - pool price deviates",
			mockOracle:   true,
			oraclePrice:  2, // Pool price is 1, oracle is 2
			expectedLog:  "Oracle price found, applying dynamic fee",
		},
		{
			name:         "oracle price not found - base fee",
			mockOracle:   false,
			oraclePrice:  0,
			expectedLog:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup keeper and mocks for each test run
			k, ctx, _, _, mockItzelKeeper, buf := keepertest.XicoatlKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			// Set params
			params := types.NewParams(baseFee, feeMultiplier)
			k.SetParams(ctx, params)

			// Create a pool
			poolMsg := &types.MsgCreatePool{
				Creator: creator,
				Index:   "test-pool",
				DenomA:  "stake",
				AmountA: poolAmount,
				DenomB:  "token",
				AmountB: poolAmount,
			}
			_, err := srv.CreatePool(ctx, poolMsg)
			require.NoError(t, err)

			// Setup mock for ItzelKeeper based on the test case
			source := "stake/token"
			if tc.mockOracle {
				mockItzelKeeper.On("GetAggregatedPrice", mock.Anything, source).Return(itzeltypes.AggregatedPrice{Price: tc.oraclePrice}, true).Once()
			} else {
				mockItzelKeeper.On("GetAggregatedPrice", mock.Anything, source).Return(itzeltypes.AggregatedPrice{}, false).Once()
			}

			// Perform a swap
			swapMsg := &types.MsgSwap{
				Creator:           creator,
				PoolId:            "test-pool",
				TokenInDenom:      "stake",
				TokenInAmount:     swapAmount,
				TokenOutDenom:     "token",
				MinTokenOutAmount: 1,
			}
			res, err := srv.Swap(ctx, swapMsg)
			require.NoError(t, err)
			require.NotNil(t, res)

			// Calculate expected fee and output
			var expectedFeePercentage math.LegacyDec
			if tc.mockOracle {
				poolPrice := math.LegacyNewDecFromInt(math.NewIntFromUint64(poolAmount)).Quo(math.LegacyNewDecFromInt(math.NewIntFromUint64(poolAmount)))
				oraclePriceDec := math.LegacyNewDec(int64(tc.oraclePrice))
				var deviation math.LegacyDec
				if poolPrice.GT(oraclePriceDec) {
					deviation = poolPrice.Sub(oraclePriceDec).Quo(oraclePriceDec)
				} else {
					deviation = oraclePriceDec.Sub(poolPrice).Quo(oraclePriceDec)
				}
				dynamicFee := deviation.Mul(feeMultiplier)
				expectedFeePercentage = baseFee.Add(dynamicFee)
			} else {
				expectedFeePercentage = baseFee
			}

			tokenInAmountAfterFee := math.LegacyNewDecFromInt(math.NewIntFromUint64(swapAmount)).Mul(math.LegacyNewDec(1).Sub(expectedFeePercentage)).TruncateInt()
			numerator := math.NewIntFromUint64(poolAmount).Mul(tokenInAmountAfterFee)
			denominator := math.NewIntFromUint64(poolAmount).Add(tokenInAmountAfterFee)
			expectedOut := numerator.Quo(denominator)

			// Check the output amount
			require.Equal(t, expectedOut.Uint64(), res.Amount)

			// Check log message
			if tc.expectedLog != "" {
				require.True(t, strings.Contains(buf.String(), tc.expectedLog))
			} else {
				require.False(t, strings.Contains(buf.String(), "Oracle price found"))
			}
		})
	}
}
