package models;

import (
	"gorm.io/gorm"
)

const StockPriceDailyType = "StockPriceDaily"

type StockPriceDaily struct {
	gorm.Model
	ID int `json:"id"`
	TickerSymbolID int `json:"tickerSymbolId"`
	TickerSymbol TickerSymbol `json:"tickerSymbol"`
	OpenPrice float32 `json:"openPrice"`
	HighPrice float32 `json:"highProce"`
	CurrentPrice float32 `json:"currentPrice"`
}