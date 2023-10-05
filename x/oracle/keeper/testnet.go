package keeper

var (
	TestnetStatus bool
)

func (k Keeper) SetTestnetStatus(val bool) {
	TestnetStatus = val
}
