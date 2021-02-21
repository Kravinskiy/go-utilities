package models;

import (
	"gorm.io/gorm"
)

const StockPriceDailyType = "StockPriceDaily"

type StockPriceDaily struct {
	gorm.Model
	ID int
	TickerSymbolID int
	TickerSymbol TickerSymbol
	OpenPrice float32
	HighPrice float32
	CurrentPrice float32
}