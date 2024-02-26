package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = &Params{}

const (
	DefaultLastTimeUpdateReserve     = uint64(1)
	DefaultGrowStakingReserveAddress = "qube13zq340zzjgua9h98pltzwv0ga5r0kkn0ryjz4v"
	DefaultUSQReserveAddress         = "qube1nx9scnpdnp5wsw88at9e35fng56788h7yz9srs"
	DefaultGrowYieldReserveAddress   = "qube1zzplgm7kqwe3vwqynzkvewrrhuffwhd7a77j7j"
	DefaultStatus                    = false
	DefaultInt                       = uint64(1)
)

var (
	KeyRealRate                  = []byte("RealRate")
	KeyBorrowRate                = []byte("BorrowRate")
	KeyLastTimeUpdateReserve     = []byte("LastTimeUpdateReserve")
	KeyGrowStakingReserveAddress = []byte("GrowStakingReserveAddress")
	KeyUSQReserveAddress         = []byte("USQReserveAddress")
	KeyGrowYieldReserveAddress   = []byte("GrowYieldReserveAddress")
	KeyDepositMethodStatus       = []byte("DepositMethodStatus")
	KeyCollateralMethodStatus    = []byte("CollateralMethodStatus")
	KeyBorrowMethodStatus        = []byte("BorrowMethodStatus")
	KeyUStaticVolatile           = []byte("UStaticVolatile")
	KeyUStaticStable             = []byte("UStaticStable")
	KeyMaxRateVolatile           = []byte("MaxRateVolatile")
	KeyMaxRateStable             = []byte("MaxRateStable")
	KeySlope                     = []byte("Slope")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		LastTimeUpdateReserve:     DefaultLastTimeUpdateReserve,
		GrowStakingReserveAddress: DefaultGrowStakingReserveAddress,
		USQReserveAddress:         DefaultUSQReserveAddress,
		GrowYieldReserveAddress:   DefaultGrowYieldReserveAddress,
		DepositMethodStatus:       DefaultStatus,
		CollateralMethodStatus:    DefaultStatus,
		BorrowMethodStatus:        DefaultStatus,
		UStaticVolatile:           DefaultInt,
		UStaticStable:             DefaultInt,
		MaxRateVolatile:           DefaultInt,
		MaxRateStable:             DefaultInt,
		Slope:                     DefaultInt,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLastTimeUpdateReserve, &p.LastTimeUpdateReserve, validate),
		paramtypes.NewParamSetPair(KeyGrowStakingReserveAddress, &p.GrowStakingReserveAddress, validate),
		paramtypes.NewParamSetPair(KeyUSQReserveAddress, &p.USQReserveAddress, validate),
		paramtypes.NewParamSetPair(KeyGrowYieldReserveAddress, &p.GrowYieldReserveAddress, validate),
		paramtypes.NewParamSetPair(KeyDepositMethodStatus, &p.DepositMethodStatus, validate),
		paramtypes.NewParamSetPair(KeyCollateralMethodStatus, &p.CollateralMethodStatus, validate),
		paramtypes.NewParamSetPair(KeyBorrowMethodStatus, &p.BorrowMethodStatus, validate),
		paramtypes.NewParamSetPair(KeyUStaticVolatile, &p.UStaticVolatile, validate),
		paramtypes.NewParamSetPair(KeyUStaticStable, &p.UStaticStable, validate),
		paramtypes.NewParamSetPair(KeyMaxRateVolatile, &p.MaxRateVolatile, validate),
		paramtypes.NewParamSetPair(KeyMaxRateStable, &p.MaxRateStable, validate),
		paramtypes.NewParamSetPair(KeySlope, &p.Slope, validate),
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
func validate(i interface{}) error {
	return nil
}
