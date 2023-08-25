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

	ExchangeRate *ExchangeRateQueryParams `json:"exchange_rate,omitempty"`
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

// ExchangeRateQueryParams query request params for exchange rates
type ExchangeRateQueryParams struct {
	Denom string `json:"denom"`
}

// ExchangeRateQueryResponse - exchange rates query response item
type ExchangeRateQueryResponse struct {
	Rate string `json:"rate"`
}
