package simulation

import (
	"math/rand"

	"github.com/0xknstntn/quadrate/x/converter/types"

	"github.com/0xknstntn/quadrate/x/converter/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgErc20Cw20(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgErc20Cw20{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Erc20Cw20 simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Erc20Cw20 simulation not implemented"), nil, nil
	}
}
