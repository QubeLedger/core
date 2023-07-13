package types

import (
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// interquery message types.
const (
	TypeMsgMakeInterchainRequest = "makerequest"
	TypeMsgSubmitQueryResponse   = "submitqueryresponse"
)

var (
	_ sdk.Msg = &MsgSubmitQueryResponse{}
	_ sdk.Msg = &MsgMakeInterchainRequest{}
)

// Route Implements Msg.
func (msg MsgSubmitQueryResponse) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSubmitQueryResponse) Type() string { return TypeMsgSubmitQueryResponse }

// ValidateBasic Implements Msg.
func (msg MsgSubmitQueryResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return err
	}

	if msg.Height < 0 {
		return errors.New("height must be non-negative")
	}

	// TODO: is there a chain validation spec in ICS?
	chainParts := strings.Split(msg.ChainId, "-")
	if len(chainParts) < 2 {
		return errors.New("chainID must be of form XXXX-N")
	}

	if len(msg.QueryId) != 64 {
		return errors.New("invalid query id")
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSubmitQueryResponse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgSubmitQueryResponse) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{fromAddress}
}

func NewMsgMakeInterchainRequest(
	connectionID,
	chainID,
	queryType string,
	request []byte,
	period int64,
	module string,
	callbackID string,
	ttl int64,
	sender string,
) *MsgMakeInterchainRequest {
	return &MsgMakeInterchainRequest{
		ConnectionId: connectionID,
		ChainId:      chainID,
		QueryType:    queryType,
		Request:      request,
		Period:       period,
		Module:       module,
		CallbackId:   callbackID,
		Ttl:          ttl,
		Sender:       sender,
	}
}

// Route Implements Msg.
func (msg MsgMakeInterchainRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgMakeInterchainRequest) Type() string { return TypeMsgMakeInterchainRequest }

// ValidateBasic Implements Msg.
func (msg MsgMakeInterchainRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	if msg.Ttl < 0 {
		return errors.New("height must be non-negative")
	}

	if msg.Period < 0 {
		return errors.New("period must be non-negative")
	}

	// TODO: is there a chain validation spec in ICS?
	chainParts := strings.Split(msg.ChainId, "-")
	if len(chainParts) < 2 {
		return errors.New("chainID must be of form XXXX-N")
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMakeInterchainRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgMakeInterchainRequest) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}
