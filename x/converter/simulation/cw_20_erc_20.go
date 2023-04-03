package simulation

import (
	"math/rand"

	"github.com/0xknstntn/quadrate/x/converter/keeper"
	"github.com/0xknstntn/quadrate/x/converter/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCw20Erc20(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCw20Erc20{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Cw20Erc20 simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Cw20Erc20 simulation not implemented"), nil, nil
	}
}
