package keeper

import (
	"bytes"
	"context"
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	itzeltypes "andean/x/itzel/types"
	"andean/x/xicoatl/keeper"
	"andean/x/xicoatl/types"
)

// MockAccountKeeper is a mock of AccountKeeper interface

type MockAccountKeeper struct {
	mock.Mock
}

func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI {
	args := m.Called(ctx, addr)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(sdk.AccountI)
}

func (m *MockAccountKeeper) GetModuleAddress(name string) sdk.AccAddress {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(sdk.AccAddress)
}

// MockBankKeeper is a mock of BankKeeper interface

type MockBankKeeper struct {
	mock.Mock
}

func (m *MockBankKeeper) SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins {
	args := m.Called(ctx, addr)
	return args.Get(0).(sdk.Coins)
}

func (m *MockBankKeeper) SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	args := m.Called(ctx, fromAddr, toAddr, amt)
	return args.Error(0)
}

// MockItzelKeeper is a mock of ItzelKeeper interface
type MockItzelKeeper struct {
	mock.Mock
}

func (m *MockItzelKeeper) GetAggregatedPrice(ctx context.Context, source string) (val itzeltypes.AggregatedPrice, found bool) {
	args := m.Called(ctx, source)
	return args.Get(0).(itzeltypes.AggregatedPrice), args.Bool(1)
}

func XicoatlKeeper(t testing.TB) (keeper.Keeper, sdk.Context, *MockBankKeeper, *MockAccountKeeper, *MockItzelKeeper, *bytes.Buffer) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	var buf bytes.Buffer
	logger := log.NewLogger(&buf)
	stateStore := store.NewCommitMultiStore(db, logger, metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	// Create mocks
	mockAccountKeeper := new(MockAccountKeeper)
	mockBankKeeper := new(MockBankKeeper)
	mockItzelKeeper := new(MockItzelKeeper)

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		logger,
		authority.String(),
		mockBankKeeper,
		mockAccountKeeper,
		mockItzelKeeper,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, logger)

	// Setup mock expectations
	moduleAddr := authtypes.NewModuleAddress(types.ModuleName)
	mockAccountKeeper.On("GetModuleAddress", types.ModuleName).Return(moduleAddr)
	mockBankKeeper.On("SendCoins", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx, mockBankKeeper, mockAccountKeeper, mockItzelKeeper, &buf
}
