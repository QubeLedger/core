package types

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyReserveFundAddress = []byte("ReserveFundAddress")
	KeyBurningFundAddress = []byte("BurningFundAddress")
)

// Default parameter values
const (
	DefaultReserveAddress = "qube13zq340zzjgua9h98pltzwv0ga5r0kkn0ryjz4v"
	DefaultBurningAddress = "qube19r3sd4r4708v48kxr7mr3dlzr3uyvrqyle6f6k"
)

var _ paramstypes.ParamSet = &Params{}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		BurningFundAddress: DefaultBurningAddress,
		ReserveFundAddress: DefaultReserveAddress,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyBurningFundAddress, &p.BurningFundAddress, validateBurningFundAddress),
		paramstypes.NewParamSetPair(KeyReserveFundAddress, &p.ReserveFundAddress, validateReserveFundAddress),
	}
}

// TODO
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// TODO
func validateReserveFundAddress(i interface{}) error {
	return nil
}

func validateBurningFundAddress(i interface{}) error {
	return nil
}
