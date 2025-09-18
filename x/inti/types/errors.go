package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/inti module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidStatus = sdkerrors.Register(ModuleName, 1102, "invalid status for transfer")
)
