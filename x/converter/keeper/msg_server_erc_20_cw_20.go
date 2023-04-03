package keeper

import (
	"context"

	"github.com/0xknstntn/quadrate/x/converter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tharsis/evmos/contracts"
)

func (k msgServer) Erc20Cw20(goCtx context.Context, msg *types.MsgErc20Cw20) (*types.MsgErc20Cw20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	erc20 := contracts.ERC20MinterBurnerDecimalsContract.ABI
	// TODO: Handling the message
	receiver := common.HexToAddress(msg.Creator)
	contract := common.HexToAddress(msg.Token)
	moduleAddress := common.HexToAddress(types.ModuleName)
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	_, err = k.CallEVM(ctx, erc20, moduleAddress, contract, true, "mint", receiver, amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgErc20Cw20Response{}, nil
}
