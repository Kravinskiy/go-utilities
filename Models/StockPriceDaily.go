package models;

import (
	"gorm.io/gorm"
)

const StockPriceDailyType = "StockPriceDaily"

type StockPriceDaily struct {
	gorm.Model
	TickerSymbolID int `json:"tickerSymbolId"`
	TickerSymbol TickerSymbol `json:"tickerSymbol"`
	OpenPrice float32 `json:"openPrice"`
	HighPrice float32 `json:"highProce"`
	CurrentPrice float32 `json:"currentPrice"`
	PreviousClosePrice float32 `json:"previousClosePrice"`
	LowPrice float32 `json:"lowPrice"`
}
  
func (StockPriceDaily) TableName() string {
	// By default it is `stock_price_dailies` - which is wrong. So override that
	return "stock_price_daily"
}