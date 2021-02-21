package models;

import (
	"gorm.io/gorm"
)

type TickerSymbol struct {
	gorm.Model
	ID int
	Currency string
	Description string
	Symbol string
	SecurityType string
	StockExchange string
	Active bool
}