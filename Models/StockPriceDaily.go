package models;

type StockPriceDaily struct {
	Id int
	Ticker_symbol_id int
	Ppen_price float32
	High_price float32
	Current_price float32
	Created_at string
}