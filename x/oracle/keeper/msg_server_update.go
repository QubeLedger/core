package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//TODO
	//Receive updates only from a specific validator

	var data = types.AcData{
		Validator: msg.Creator,
		Data:      msg.Data,
		Time:      msg.Time,
	}

	k.AppendAcData(ctx, data)

	return &types.MsgUpdateResponse{}, nil
}
