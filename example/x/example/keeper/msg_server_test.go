package keeper_test

import (
	"context"
	"testing"

	keepertest "example/testutil/keeper"
	"example/x/example/keeper"
	"example/x/example/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ExampleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
