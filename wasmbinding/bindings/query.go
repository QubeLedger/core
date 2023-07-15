package bindings

// OsmosisQuery contains osmosis custom queries.
// See https://github.com/osmosis-labs/osmosis-bindings/blob/main/packages/bindings/src/query.rs
type OsmosisQuery struct {
	/// Given a subdenom minted by a contract via `OsmosisMsg::MintTokens`,
	/// returns the full denom as used by `BankMsg::Send`.
	FullDenom *FullDenom `json:"full_denom,omitempty"`
	/// Returns the admin of a denom, if the denom is a Token Factory denom.
	DenomAdmin *DenomAdmin `json:"denom_admin,omitempty"`

	ActualPrice *ActualPrice `json:"actual_price,omitempty"`
}

type FullDenom struct {
	CreatorAddr string `json:"creator_addr"`
	Subdenom    string `json:"subdenom"`
}

type ActualPrice struct{}

type DenomAdmin struct {
	Subdenom string `json:"subdenom"`
}

type DenomAdminResponse struct {
	Admin string `json:"admin"`
}

type ActualPriceResponse struct {
	Atom   string `json:"atom"`
	StAtom string `json:"statom"`
}

type FullDenomResponse struct {
	Denom string `json:"denom"`
}

type InterchainQueryResponse struct {
	Id           string `json:"id"`
	ConnectionId string `json:"сonnection_id"`
	ChainId      string `json:"chain_id"`
	QueryType    string `json:"query_type"`
	Request      []byte `json:"request"`
	Period       int64  `json:"period"`
	LastHeight   int64  `json:"last_height"`
	CallbackId   string `json:"callback_id"`
	Ttl          uint64 `json:"ttl"`
}

type InterchainQueryResultResponse struct {
	Id           string `json:"id"`
	LocalHeight  int64  `json:"local_height"`
	RemoteHeight int64  `json:"remote_height"`
	Value        []byte `json:"value"`
}
