package yfinance

import (
	"net/http"
	"net/url"

	"github.com/shopspring/decimal"
)

const getChartPath = "/v8/finance/chart"

type GetChartParam struct {
	Ticker string

	Range    DataRange
	Interval DataGranularity

	AutoAdjust bool
	BackAdjust bool
}

func (r GetChartParam) createURL() string {
	queryParam := make(url.Values)
	if r.Range.IsValid() {
		queryParam.Add("range", string(r.Range))
	}
	if r.Interval.IsValid() {
		queryParam.Add("interval", string(r.Range))
	}

	return yahooFinanceAPIURL + getChartPath + "/" + r.Ticker + "?" + queryParam.Encode()
}

func (r *Client) GetChart(param *GetChartParam) (*Chart, error) {
	if param == nil || len(param.Ticker) == 0 {
		return nil, createInvalidParameterError()
	}

	req, err := http.NewRequest(http.MethodGet, param.createURL(), nil)
	if err != nil {
		return nil, err
	}

	var res chart
	if err := r.do(req, &res); err != nil {
		return nil, err
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return res.ToChart(param)
}

type chart struct {
	Chart struct {
		Result *[]struct {
			Meta       ChartMeta `json:"meta"`
			Timestamp  []int64   `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open   []*decimal.Decimal `json:"open"`
					High   []*decimal.Decimal `json:"high"`
					Low    []*decimal.Decimal `json:"low"`
					Close  []*decimal.Decimal `json:"close"`
					Volume []uint64           `json:"volume"`
				} `json:"quote"`
				Adjclose *[]struct {
					AdjClose []*decimal.Decimal `json:"adjclose"`
				} `json:"adjclose,omitempty"`
			} `json:"indicators"`
		} `json:"result,omitempty"`
		Error *Error `json:"error,omitempty"`
	} `json:"chart"`
}

func (r chart) Err() error {
	if r.Chart.Error != nil {
		return r.Chart.Error
	}

	return nil
}

func (r chart) ToChart(param *GetChartParam) (*Chart, error) {
	result := r.Chart.Result
	if result == nil || len(*result) == 0 {
		return nil, createInvalidResponseError()
	}

	quote := (*result)[0].Indicators.Quote
	if len(quote) == 0 {
		return nil, createInvalidResponseError()
	}

	timestamps := (*result)[0].Timestamp
	opens := quote[0].Open
	highs := quote[0].High
	lows := quote[0].Low
	closes := quote[0].Close
	volumes := quote[0].Volume

	if len(timestamps) != len(opens) ||
		len(opens) != len(highs) ||
		len(highs) != len(lows) ||
		len(lows) != len(closes) ||
		len(closes) != len(volumes) {
		return nil, createInvalidResponseError()
	}

	adjclose := (*result)[0].Indicators.Adjclose
	if adjclose != nil {
		if len(*adjclose) == 0 {
			return nil, createInvalidResponseError()
		}
		adjcloses := (*adjclose)[0].AdjClose
		if len(timestamps) != len(adjcloses) {
			return nil, createInvalidResponseError()
		}

		for i, ac := range adjcloses {
			o, h, l, c := opens[i], highs[i], lows[i], closes[i]
			if ac == nil || o == nil || h == nil || l == nil || c == nil {
				continue
			}
			switch {
			case param.AutoAdjust:
				ratio := c.Div(*ac)
				*opens[i] = o.Div(ratio)
				*highs[i] = h.Div(ratio)
				*lows[i] = l.Div(ratio)
			case param.BackAdjust:
				ratio := ac.Div(*c)
				*opens[i] = o.Mul(ratio)
				*highs[i] = h.Mul(ratio)
				*lows[i] = l.Mul(ratio)
			}
		}
	}

	candles := make([]ChartCandle, 0)
	for i := range timestamps {
		o, h, l, c := opens[i], highs[i], lows[i], closes[i]
		if o == nil || h == nil || l == nil || c == nil {
			continue
		}
		candles = append(candles, ChartCandle{
			Open:      *o,
			High:      *h,
			Low:       *l,
			Close:     *c,
			Volume:    volumes[i],
			Timestamp: timestamps[i],
		})
	}

	return &Chart{
		Meta:    (*result)[0].Meta,
		Candles: candles,
	}, nil
}
