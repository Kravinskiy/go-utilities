package models;

type StockPriceDaily struct {
	id int
	ticker_symbol_id int
	open_price float32
	high_price float32
	current_price float32
	created_at string
}