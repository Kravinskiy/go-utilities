package models;

type StockPriceDaily struct {
	ID int
	TickerSymbolID int
	TickerSymbol TickerSymbol
	OpenPrice float32
	HighPrice float32
	CurrentPrice float32
}