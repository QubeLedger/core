package keeper

var (
	TestnetStatus bool = false
)

func (k Keeper) SetTestnetStatus(val bool) {
	TestnetStatus = val
}
