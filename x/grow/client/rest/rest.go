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

	"github.com/QuadrateOrg/core/x/grow/types"
)

type RegisterLendAssetProposal struct {
	BaseReq       rest.BaseReq       `json:"base_req" yaml:"base_req"`
	Title         string             `json:"title" yaml:"title"`
	Description   string             `json:"description" yaml:"description"`
	Deposit       sdk.Coins          `json:"deposit" yaml:"deposit"`
	AssetMetadata banktypes.Metadata `json:"assetMetadata" yaml:"assetMetadata"`
	OracleAssetId string             `json:"oracleAssetId" yaml:"oracleAssetId"`
}

type RegisterGTokenPairProposal struct {
	BaseReq        rest.BaseReq       `json:"base_req" yaml:"base_req"`
	Title          string             `json:"title" yaml:"title"`
	Description    string             `json:"description" yaml:"description"`
	Deposit        sdk.Coins          `json:"deposit" yaml:"deposit"`
	GTokenMetadata banktypes.Metadata `json:"gTokenMetadata" yaml:"gTokenMetadata"`
	QStablePairId  string             `json:"qStablePairId" yaml:"qStablePairId"`
	MinAmountIn    string             `json:"minAmountIn" yaml:"minAmountIn"`
	MinAmountOut   string             `json:"minAmountOut" yaml:"minAmountOut"`
}

type RegisterChangeGrowYieldReserveAddressProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Address     string       `json:"address" yaml:"address"`
}

type RegisterChangeUSQReserveAddressProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Address     string       `json:"address" yaml:"address"`
}

type RegisterChangeGrowStakingReserveAddressProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Address     string       `json:"address" yaml:"address"`
}

type RegisterChangeRealRateProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Rate        uint64       `json:"rate" yaml:"rate"`
}

type RegisterChangeBorrowRateProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Rate        uint64       `json:"rate" yaml:"rate"`
}

type RegisterChangeLendRateProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Rate        uint64       `json:"rate" yaml:"rate"`
	Id          string       `json:"id" yaml:"id"`
}

type RegisterActivateGrowModuleProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
}

type RegisterRemoveLendAssetProposal struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	LendAssetId string       `json:"lendAssetId" yaml:"lendAssetId"`
}

type RegisterRemoveGTokenPairProposal struct {
	BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title        string       `json:"title" yaml:"title"`
	Description  string       `json:"description" yaml:"description"`
	Deposit      sdk.Coins    `json:"deposit" yaml:"deposit"`
	GTokenPairID string       `json:"lendAssetId" yaml:"lendAssetId"`
}

func RegisterLendAssetProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterLendAssetProposal(clientCtx),
	}
}

func newRegisterLendAssetProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterLendAssetProposal

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

		content := types.NewRegisterLendAssetProposal(req.Title, req.Description, req.AssetMetadata, req.OracleAssetId)
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

func RegisterGTokenPairProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterGTokenPairProposal(clientCtx),
	}
}

func newRegisterGTokenPairProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterGTokenPairProposal

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

		content := types.NewRegisterGTokenPairProposal(req.Title, req.Description, req.GTokenMetadata, req.QStablePairId, req.MinAmountIn, req.MinAmountOut)
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

func RegisterChangeGrowYieldReserveAddressProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeGrowYieldReserveAddressProposal(clientCtx),
	}
}

func newRegisterChangeGrowYieldReserveAddressProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeGrowYieldReserveAddressProposal

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

		content := types.NewRegisterChangeGrowYieldReserveAddressProposal(req.Title, req.Description, req.Address)
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

func RegisterChangeUSQReserveAddressProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeUSQReserveAddressProposal(clientCtx),
	}
}

func newRegisterChangeUSQReserveAddressProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeUSQReserveAddressProposal

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

		content := types.NewRegisterChangeUSQReserveAddressProposal(req.Title, req.Description, req.Address)
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

func RegisterChangeGrowStakingReserveAddressProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeGrowStakingReserveAddressProposal(clientCtx),
	}
}

func newRegisterChangeGrowStakingReserveAddressProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeGrowStakingReserveAddressProposal

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

		content := types.NewRegisterChangeGrowStakingReserveAddressProposal(req.Title, req.Description, req.Address)
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

func RegisterChangeRealRateProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeRealRateProposal(clientCtx),
	}
}

func newRegisterChangeRealRateProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeRealRateProposal

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

		content := types.NewRegisterChangeRealRateProposal(req.Title, req.Description, req.Rate)
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

func RegisterChangeBorrowRateProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeBorrowRateProposal(clientCtx),
	}
}

func newRegisterChangeBorrowRateProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeBorrowRateProposal

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

		content := types.NewRegisterChangeBorrowRateProposal(req.Title, req.Description, req.Rate)
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

func RegisterChangeLendRateProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterChangeLendRateProposal(clientCtx),
	}
}

func newRegisterChangeLendRateProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeLendRateProposal

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

		content := types.NewRegisterChangeLendRateProposal(req.Title, req.Description, req.Rate, req.Id)
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

func RegisterActivateGrowModuleProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterActivateGrowModuleProposal(clientCtx),
	}
}

func newRegisterActivateGrowModuleProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterChangeRealRateProposal

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

		content := types.NewRegisterActivateGrowModuleProposal(req.Title, req.Description)
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

func RegisterRemoveLendAssetProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterRemoveLendAssetProposal(clientCtx),
	}
}

func newRegisterRemoveLendAssetProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRemoveLendAssetProposal

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

		content := types.NewRegisterRemoveLendAssetProposal(req.Title, req.Description, req.LendAssetId)
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

func RegisterRemoveGTokenPairProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterRemoveGTokenPairProposal(clientCtx),
	}
}

func newRegisterRemoveGTokenPairProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRemoveGTokenPairProposal

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

		content := types.NewRegisterRemoveGTokenPairProposal(req.Title, req.Description, req.GTokenPairID)
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
