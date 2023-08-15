package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/QubeLedger/core/x/stable/types"
)

type RegisterPairProposal struct {
	BaseReq           rest.BaseReq       `json:"base_req" yaml:"base_req"`
	Title             string             `json:"title" yaml:"title"`
	Description       string             `json:"description" yaml:"description"`
	Deposit           sdk.Coins          `json:"deposit" yaml:"deposit"`
	AmountInMetadata  banktypes.Metadata `json:"amountInMetadata" yaml:"amountInMetadata"`
	AmountOutMetadata banktypes.Metadata `json:"amountOutMetadata" yaml:"amountOutMetadata"`
	MinAmountIn       string             `json:"minAmountIn" yaml:"minAmountIn"`
	MinAmountOut      string             `json:"minAmountOut" yaml:"minAmountOut"`
}

type RegisterChangeBurningFundAddressProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Address     string       `json:"address" yaml:"address"`
}

func RegisterPairRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterPair(clientCtx),
	}
}

func newRegisterPair(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterPairProposal

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewRegisterPairProposal(req.Title, req.Description, req.AmountInMetadata, req.AmountOutMetadata, req.MinAmountIn, req.MinAmountOut)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func RegisterChangeBurningFundAddressProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeBurningFundAddressProposal(clientCtx),
	}
}

func newRegisterChangeBurningFundAddressProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeBurningFundAddressProposal

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewRegisterChangeBurningFundAddressProposal(req.Title, req.Description, req.Address)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
