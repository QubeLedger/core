package simulation

import (
	"math/rand"

	"github.com/QuadrateOrg/core/x/oracle/keeper"
	"github.com/QuadrateOrg/core/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdate(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdate{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Update simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Update simulation not implemented"), nil, nil
	}
}
