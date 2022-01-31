package yfinance

import (
	"github.com/shopspring/decimal"
)

type Chart struct {
	Meta    *ChartMeta
	Candles []*ChartCandle
}

type ChartCandle struct {
	Open      *decimal.Decimal
	High      *decimal.Decimal
	Low       *decimal.Decimal
	Close     *decimal.Decimal
	Adjclose  *decimal.Decimal
	Volume    int64
	Timestamp int64
}

type ChartMeta struct {
	Currency             string           `json:"currency"`
	Symbol               string           `json:"symbol"`
	ExchangeName         string           `json:"exchangeName"`
	InstrumentType       InstrumentType   `json:"instrumentType"`
	FirstTradeDate       int64            `json:"firstTradeDate"`
	RegularMarketTime    int64            `json:"regularMarketTime"`
	Gmtoffset            int              `json:"gmtoffset"`
	Timezone             string           `json:"timezone"`
	ExchangeTimezoneName string           `json:"exchangeTimezoneName"`
	RegularMarketPrice   *decimal.Decimal `json:"regularMarketPrice"`
	ChartPreviousClose   *decimal.Decimal `json:"chartPreviousClose"`
	PreviousClose        *decimal.Decimal `json:"previousClose"`
	Scale                int              `json:"scale"`
	PriceHint            int              `json:"priceHint"`
	CurrentTradingPeriod *struct {
		Pre     *TradingPeriod `json:"pre"`
		Regular *TradingPeriod `json:"regular"`
		Post    *TradingPeriod `json:"post"`
	}
	TradingPeriods  [][]*TradingPeriod `json:"tradingPeriods"`
	DataGranularity DataGranularity    `json:"dataGranularity"`
	Range           DataRange          `json:"range"`
	ValidRanges     []DataRange        `json:"validRanges"`
}

type TradingPeriod struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Gmtoffset int    `json:"gmtoffset"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type DataGranularity string

const (
	DataGranularity1m  DataGranularity = "1m"
	DataGranularity5m  DataGranularity = "5m"
	DataGranularity15m DataGranularity = "15m"
	DataGranularity1d  DataGranularity = "1d"
	DataGranularity1wk DataGranularity = "1wk"
	DataGranularity1mo DataGranularity = "1mo"
)

type DataRange string

const (
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

type AutocompleteResultSet struct {
	Results []*AutocompleteResult
}

type AutocompleteResult struct {
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	ExchangeDisp       string `json:"exchDisp"`
	InstrumentTypeDisp string `json:"typeDisp"`
}
