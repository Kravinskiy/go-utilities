package models;

import (
	"gorm.io/gorm"
)

type StockPriceDaily struct {
	gorm.Model
	ID int
	TickerSymbolID int
	TickerSymbol TickerSymbol
	OpenPrice float32
	HighPrice float32
	CurrentPrice float32
}