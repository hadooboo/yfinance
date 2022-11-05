package yfinance

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type DataGranularity string

const (
	// 1m,2m,5m,15m,30m,60m,90m,1h,1d,5d,1wk,1mo,3mo
	DataGranularity1m  DataGranularity = "1m"
	DataGranularity2m  DataGranularity = "2m"
	DataGranularity5m  DataGranularity = "5m"
	DataGranularity15m DataGranularity = "15m"
	DataGranularity30m DataGranularity = "30m"
	DataGranularity60m DataGranularity = "60m"
	DataGranularity90m DataGranularity = "90m"
	DataGranularity1h  DataGranularity = "1h"
	DataGranularity1d  DataGranularity = "1d"
	DataGranularity5d  DataGranularity = "5d"
	DataGranularity1wk DataGranularity = "1wk"
	DataGranularity1mo DataGranularity = "1mo"
	DataGranularity3mo DataGranularity = "3mo"
)

func (r DataGranularity) IsValid() bool {
	switch r {
	case DataGranularity1m, DataGranularity2m, DataGranularity5m, DataGranularity15m, DataGranularity30m,
		DataGranularity60m, DataGranularity90m, DataGranularity1h, DataGranularity1d, DataGranularity5d,
		DataGranularity1wk, DataGranularity1mo, DataGranularity3mo:
		return true
	default:
		return false
	}
}

type DataRange string

const (
	// 1d,5d,1mo,3mo,6mo,1y,2y,5y,10y,ytd,max
	DataRange1d  DataRange = "1d"
	DataRange5d  DataRange = "5d"
	DataRange1mo DataRange = "1mo"
	DataRange3mo DataRange = "3mo"
	DataRange6mo DataRange = "6mo"
	DataRange1y  DataRange = "1y"
	DataRange2y  DataRange = "2y"
	DataRange5y  DataRange = "5y"
	DataRange10y DataRange = "10y"
	DataRangeYtd DataRange = "ytd"
	DataRangeMax DataRange = "max"
)

func (r DataRange) IsValid() bool {
	switch r {
	case DataRange1d, DataRange5d, DataRange1mo, DataRange3mo, DataRange6mo,
		DataRange1y, DataRange2y, DataRange5y, DataRange10y, DataRangeYtd, DataRangeMax:
		return true
	default:
		return false
	}
}

type InstrumentType string

const (
	InstrumentTypeETF            InstrumentType = "ETF"
	InstrumentTypeEquity         InstrumentType = "EQUITY"
	InstrumentTypeMutualFund     InstrumentType = "MUTUALFUND"
	InstrumentTypeCurrency       InstrumentType = "CURRENCY"
	InstrumentTypeCryptoCurrency InstrumentType = "CRYPTOCURRENCY"
	InstrumentTypeFuture         InstrumentType = "FUTURE"
	InstrumentTypeIndex          InstrumentType = "INDEX"
	InstrumentTypeOption         InstrumentType = "OPTION"
)

func (r InstrumentType) IsValid() bool {
	switch r {
	case InstrumentTypeETF, InstrumentTypeEquity, InstrumentTypeMutualFund, InstrumentTypeCurrency,
		InstrumentTypeCryptoCurrency, InstrumentTypeFuture, InstrumentTypeIndex, InstrumentTypeOption:
		return true
	default:
		return false
	}
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (r Error) Error() string {
	return fmt.Sprintf("%s: %s", r.Code, r.Description)
}

type Chart struct {
	Meta    ChartMeta
	Candles []ChartCandle
}

type ChartCandle struct {
	Open      decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Close     decimal.Decimal
	Volume    uint64
	Timestamp int64
}

type ChartMeta struct {
	Currency             string               `json:"currency"`
	Symbol               string               `json:"symbol"`
	ExchangeName         string               `json:"exchangeName"`
	InstrumentType       InstrumentType       `json:"instrumentType"`
	FirstTradeDate       int64                `json:"firstTradeDate"`
	RegularMarketTime    int64                `json:"regularMarketTime"`
	Gmtoffset            int                  `json:"gmtoffset"`
	Timezone             string               `json:"timezone"`
	ExchangeTimezoneName string               `json:"exchangeTimezoneName"`
	RegularMarketPrice   decimal.Decimal      `json:"regularMarketPrice"`
	ChartPreviousClose   decimal.Decimal      `json:"chartPreviousClose"`
	PreviousClose        decimal.Decimal      `json:"previousClose"`
	Scale                uint                 `json:"scale"`
	PriceHint            uint                 `json:"priceHint"`
	CurrentTradingPeriod CurrentTradingPeriod `json:"currentTradingPeriod"`
	DataGranularity      DataGranularity      `json:"dataGranularity"`
	Range                DataRange            `json:"range"`
	ValidRanges          []DataRange          `json:"validRanges"`
}

type CurrentTradingPeriod struct {
	Pre     TradingPeriod `json:"pre"`
	Regular TradingPeriod `json:"regular"`
	Post    TradingPeriod `json:"post"`
}

type TradingPeriod struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Gmtoffset int    `json:"gmtoffset"`
}

type Autocomplete struct {
	Results []AutocompleteResult
}

type AutocompleteResult struct {
	Symbol             string         `json:"symbol"`
	Name               string         `json:"name"`
	Exch               string         `json:"exch"`
	InstrumentType     string         `json:"type"`
	ExchangeDisp       string         `json:"exchDisp"`
	InstrumentTypeDisp InstrumentType `json:"typeDisp"`
}
