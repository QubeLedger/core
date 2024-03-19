package keeper

import (
	liquidstakeibctypes "github.com/QuadrateOrg/core/x/liquidstakeibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

func (k *Keeper) GenerateAndExecuteICATx(
	ctx sdk.Context,
	connectionID string,
	ownerID string,
	messages []sdk.Msg,
) (string, error) {
	msgData, err := icatypes.SerializeCosmosTx(k.cdc, messages)
	if err != nil {
		return "", err
	}

	icaPacketData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: msgData,
	}

	portID, err := icatypes.NewControllerPortID(ownerID)
	if err != nil {
		return "", err
	}

	activeChannelID, found := k.icaControllerKeeper.GetOpenActiveChannel(ctx, connectionID, portID)
	if !found {
		return "", sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel on connection %s for port %s", connectionID, k.GetPortID(ownerID))
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, activeChannelID))
	if !found {
		return "", sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	//handler := k.msgRouter.Handler(msgSendTx)
	seq, err := k.icaControllerKeeper.SendTx(ctx, chanCap, connectionID, portID, icaPacketData, ^uint64(0))
	if err != nil {
		return "", sdkerrors.Wrapf(liquidstakeibctypes.ErrICATxFailure, "failed to send ica msg with err: %v", err)
	}

	channelID, found := k.icaControllerKeeper.GetOpenActiveChannel(ctx, connectionID, portID)
	if !found {
		return "", sdkerrors.Wrapf(
			liquidstakeibctypes.ErrICATxFailure,
			"failed to get ica active channel: %v",
			err,
		)
	}

	k.Logger(ctx).Info(
		"Sent ICA transactions",
		liquidstakeibctypes.SequenceIDKeyVal,
		seq,
		liquidstakeibctypes.ConnectionKeyVal,
		connectionID,
		liquidstakeibctypes.PortKeyVal,
		ownerID,
		liquidstakeibctypes.MessagesKeyVal,
		messages,
	)

	return k.GetTransactionSequenceID(channelID, seq), nil
}
