package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"andean/testutil/sample"
	keepertest "andean/testutil/keeper"
	"andean/x/inti/keeper"
	"andean/x/inti/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.IntiKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func TestCreateLazyTransfer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Define test constants
	creatorStr := sample.AccAddress()
	creator, err := sdk.AccAddressFromBech32(creatorStr)
	require.NoError(t, err)
	denom := sdk.DefaultBondDenom // "stake"
	initialBalance := math.NewInt(1000000)
	transferAmount := math.NewInt(100)

	// Fund the creator's account by minting coins to the module and then sending them to the creator
	require.NoError(t, k.BankKeeper.MintCoins(sdkCtx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, initialBalance))))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, creator, sdk.NewCoins(sdk.NewCoin(denom, initialBalance))))

	// Check initial balance
	require.Equal(t, initialBalance, k.BankKeeper.GetBalance(sdkCtx, creator, denom).Amount)

	// Create the message
	msg := &types.MsgCreateLazyTransfer{
		Creator:          creatorStr,
		Amount:           sdk.NewCoin(denom, transferAmount),
		Recipient:        "recipient_on_other_chain",
		DestinationChain: "other_chain",
	}

	// Execute the transaction
	res, err := ms.CreateLazyTransfer(ctx, msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	// --- Assertions ---

	// 1. Check final balance of the creator
	finalBalance := k.BankKeeper.GetBalance(sdkCtx, creator, denom).Amount
	expectedFinalBalance := initialBalance.Sub(transferAmount)
	require.True(t, expectedFinalBalance.Equal(finalBalance), "Creator's final balance is incorrect")

	// 2. Check the balance of the module account (funds should be locked there)
	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)
	moduleBalance := k.BankKeeper.GetBalance(sdkCtx, moduleAddr, denom).Amount
	require.True(t, transferAmount.Equal(moduleBalance), "Module's balance is incorrect")

	// 3. Check for the event
	var eventFound bool
	for _, event := range sdkCtx.EventManager().Events() {
		if event.Type == "LazyTransferCreated" {
			eventFound = true
		}
	}
	require.True(t, eventFound, "LazyTransferCreated event not found")
}
