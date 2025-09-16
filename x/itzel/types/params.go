package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyAuthorizedOracles = []byte("AuthorizedOracles")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(authorizedOracles []string) Params {
	return Params{
		AuthorizedOracles: authorizedOracles,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(nil)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAuthorizedOracles, &p.AuthorizedOracles, validateAuthorizedOracles),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateAuthorizedOracles(p.AuthorizedOracles)
}

func validateAuthorizedOracles(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, addr := range v {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return fmt.Errorf("invalid authorized oracle address: %s", err)
		}
	}

	return nil
}
