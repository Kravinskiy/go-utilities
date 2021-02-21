package models;

import (
	"gorm.io/gorm"
)

type SelectedStocksForProcessing struct {
	gorm.Model
	ID int
	TickerSymbolID int
	TickerSymbol TickerSymbol
}