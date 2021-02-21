package models;

import (
	"gorm.io/gorm"
)

type SelectedStocksForProcessing struct {
	gorm.Model
	ID int `json:"id"`
	TickerSymbolID int `json:"tickerSymbolId"`
	TickerSymbol TickerSymbol `json:"tickerSmybol"`
}

type Tabler interface {
	TableName() string
  }
  
func (SelectedStocksForProcessing) TableName() string {
	// By default it is `selected_stocks_for_processings` - which is wrong. So override that
	return "selected_stocks_for_processing"
}