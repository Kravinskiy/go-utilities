package models;

import (
	"gorm.io/gorm"
)

const StockPriceDailyType = "StockPriceDaily"

type StockPriceDaily struct {
	gorm.Model
	TickerSymbolID int `json:"tickerSymbolId"`
	TickerSymbol TickerSymbol `json:"tickerSymbol"`
	OpenPrice float64 `json:"openPrice"`
	HighPrice float64 `json:"highProce"`
	CurrentPrice float64 `json:"currentPrice"`
	PreviousClosePrice float64 `json:"previousClosePrice"`
	LowPrice float64 `json:"lowPrice"`
	TrailingPe float64 `json:"trailingPe"`
	ForwardPe float64 `json:"forwardPe"`
	BookValue float64 `json:"bookValue"`
	PriceToBook float64 `json:"priceToBook"`
	TrailingAnnualDividendYield float64 `json:"TrailingAnnualDividendYield"`
	TrailingAnnualDividendRate float64 `json:"TrailingAnnualDividendRate"`
	MarketCap int64 `json:"MarketCap"`
	SharesOutstanding int `json:"SharesOutstanding"`
}
  
func (StockPriceDaily) TableName() string {
	// By default it is `stock_price_dailies` - which is wrong. So override that
	return "stock_price_daily"
}