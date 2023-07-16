package bindings

type QubeQuery struct {
	FullDenom             *FullDenom             `json:"full_denom,omitempty"`
	DenomAdmin            *DenomAdmin            `json:"denom_admin,omitempty"`
	ActualPrice           *ActualPrice           `json:"actual_price,omitempty"`
	InterchainQuery       *InterchainQuery       `json:"interchain_query,omitempty"`
	InterchainQueryResult *InterchainQueryResult `json:"interchain_query_result,omitempty"`
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

type InterchainQuery struct {
	Id string `json:"id"`
}

type InterchainQueryResult struct {
	Module       string `json:"module"`
	ConnectionId string `json:"сonnection_id"`
	ChainId      string `json:"chain_id"`
	QueryType    string `json:"query_type"`
	Request      []byte `json:"request"`
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
