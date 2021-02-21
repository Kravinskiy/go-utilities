package types

type ExecuteTaskRequest struct {
	Type string `json:"type"`
	Currency  string `json:"currency"`
	Ticker_Symbol string `json:"ticker_symbol"`
	Ticker_Symbol_Id int `json:"ticker_symbol_id"`
}