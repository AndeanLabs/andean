package types

import (
	"fmt"

	"cosmossdk.io/math"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyBaseFee       = []byte("BaseFee")
	KeyFeeMultiplier = []byte("FeeMultiplier")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(baseFee math.LegacyDec, feeMultiplier math.LegacyDec) Params {
	return Params{
		BaseFee:       baseFee,
		FeeMultiplier: feeMultiplier,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		math.LegacyNewDecWithPrec(3, 3), // 0.3%
		math.LegacyNewDecWithPrec(1, 1), // 0.1
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBaseFee, &p.BaseFee, validateBaseFee),
		paramtypes.NewParamSetPair(KeyFeeMultiplier, &p.FeeMultiplier, validateFeeMultiplier),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBaseFee(p.BaseFee); err != nil {
		return err
	}
	if err := validateFeeMultiplier(p.FeeMultiplier); err != nil {
		return err
	}
	return nil
}

func validateBaseFee(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("base fee cannot be nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("base fee cannot be negative")
	}

	return nil
}

func validateFeeMultiplier(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("fee multiplier cannot be nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("fee multiplier cannot be negative")
	}

	return nil
}
