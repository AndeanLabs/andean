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

	// Test cases
	tests := []struct {
		name              string
		mockOracle        bool
		expectedFee       math.LegacyDec
		expectedLog       string
	}{
		{
			name:              "oracle price found - 0.5% fee",
			mockOracle:        true,
			expectedFee:       math.LegacyNewDecWithPrec(5, 3), // 0.5%
			expectedLog:       "Oracle price found for source",
		},
		{
			name:              "oracle price not found - 0.3% fee",
			mockOracle:        false,
			expectedFee:       math.LegacyNewDecWithPrec(3, 3), // 0.3%
			expectedLog:       "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup keeper and mocks for each test run
			k, ctx, _, _, mockItzelKeeper, buf := keepertest.XicoatlKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

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
				oraclePrice := int32(1) // 1:1 price
				mockItzelKeeper.On("GetAggregatedPrice", mock.Anything, source).Return(itzeltypes.AggregatedPrice{Price: oraclePrice}, true).Once()
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
				MinTokenOutAmount: 1, // Low min amount to not interfere with fee check
			}
			res, err := srv.Swap(ctx, swapMsg)
			require.NoError(t, err)
			require.NotNil(t, res)

			// Calculate expected output
			tokenInAmountAfterFee := math.LegacyNewDecFromInt(math.NewIntFromUint64(swapAmount)).Mul(math.LegacyNewDec(1).Sub(tc.expectedFee)).TruncateInt()
			numerator := math.NewIntFromUint64(poolAmount).Mul(tokenInAmountAfterFee)
			denominator := math.NewIntFromUint64(poolAmount).Add(tokenInAmountAfterFee)
			expectedOut := numerator.Quo(denominator)

			// Check the output amount
			require.Equal(t, expectedOut.Uint64(), res.Amount)

			// Check log message
			if tc.expectedLog != "" {
				require.True(t, strings.Contains(buf.String(), tc.expectedLog))
			} else {
				require.False(t, strings.Contains(buf.String(), "Oracle price found for source"))
			}
		})
	}
}
