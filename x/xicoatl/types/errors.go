package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/xicoatl module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrPoolNotFound  = sdkerrors.Register(ModuleName, 1101, "pool not found")
	ErrInvalidTokens = sdkerrors.Register(ModuleName, 1102, "invalid token pair for pool")
	ErrZeroOutput    = sdkerrors.Register(ModuleName, 1103, "swap results in zero output")
	ErrSlippage      = sdkerrors.Register(ModuleName, 1104, "slippage protection: output amount is less than minimum expected")
)
