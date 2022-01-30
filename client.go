package yfinance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/shopspring/decimal"
)

const (
	yahooFinanceAPIURL = "https://query2.finance.yahoo.com"
	defaultTimeout     = 30 * time.Second

	getChartPath = "/v8/finance/chart"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{httpClient: &http.Client{
		Timeout: defaultTimeout,
	}}
}

func (r *Client) do(req *http.Request, v interface{}) error {
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("received an error response from api: " + resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (r *Client) get(path, pathParam string, queryParams *url.Values, v interface{}) error {
	fullURL := yahooFinanceAPIURL + path
	if pathParam != "" {
		fullURL += fmt.Sprintf("/%s", pathParam)
	}
	fullURL += fmt.Sprintf("?%s", queryParams.Encode())

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}

	return r.do(req, v)
}

func (r *Client) GetChart(params *GetChartParams) (*Chart, error) {
	pathParam := params.Ticker
	queryParams := &url.Values{}
	queryParams.Set("range", string(params.Range))
	queryParams.Set("interval", string(params.Interval))

	resp := new(getChartResp)
	if err := r.get(getChartPath, pathParam, queryParams, resp); err != nil {
		return nil, err
	}

	return &Chart{
		Meta:    resp.Meta,
		Candles: resp.Candles,
	}, nil
}

type GetChartParams struct {
	Range    DataRange
	Interval DataGranularity
	Ticker   string
}

type getChartResp struct {
	Meta    *ChartMeta
	Candles []*ChartCandle
}

func (r *getChartResp) UnmarshalJSON(data []byte) error {
	chart := struct {
		Chart *struct {
			Result []*struct {
				Meta       *ChartMeta `json:"meta"`
				Timestamp  []int64    `json:"timestamp"`
				Indicators *struct {
					Quote []*struct {
						Open   []*decimal.Decimal `json:"open"`
						High   []*decimal.Decimal `json:"high"`
						Low    []*decimal.Decimal `json:"low"`
						Close  []*decimal.Decimal `json:"close"`
						Volume []int64            `json:"volume"`
					} `json:"quote"`
					Adjclose []*struct {
						Adjclose []*decimal.Decimal `json:"adjclose"`
					} `json:"adjclose,omitempty"`
				} `json:"indicators"`
			} `json:"result"`
			Error *Error `json:"error"`
		} `json:"chart"`
	}{}
	if err := json.Unmarshal(data, &chart); err != nil {
		return err
	}
	if chart.Chart.Error != nil {
		return errors.New(chart.Chart.Error.Description)
	}

	v := chart.Chart.Result[0]
	if len(v.Timestamp) != len(v.Indicators.Quote[0].Open) {
		return errors.New("length of timestamp and open should be equal")
	}
	if len(v.Timestamp) != len(v.Indicators.Quote[0].High) {
		return errors.New("length of timestamp and high should be equal")
	}
	if len(v.Timestamp) != len(v.Indicators.Quote[0].Low) {
		return errors.New("length of timestamp and low should be equal")
	}
	if len(v.Timestamp) != len(v.Indicators.Quote[0].Close) {
		return errors.New("length of timestamp and close should be equal")
	}
	if len(v.Timestamp) != len(v.Indicators.Quote[0].Volume) {
		return errors.New("length of timestamp and volume should be equal")
	}
	if v.Indicators.Adjclose != nil {
		if len(v.Timestamp) != len(v.Indicators.Adjclose[0].Adjclose) {
			return errors.New("length of timestamp and adjclose should be equal")
		}
	}

	candles := make([]*ChartCandle, 0, len(v.Timestamp))
	for i := 0; i < len(v.Timestamp); i++ {
		candle := &ChartCandle{
			Open:      v.Indicators.Quote[0].Open[i],
			High:      v.Indicators.Quote[0].High[i],
			Low:       v.Indicators.Quote[0].Low[i],
			Close:     v.Indicators.Quote[0].Close[i],
			Volume:    v.Indicators.Quote[0].Volume[i],
			Timestamp: v.Timestamp[i],
		}
		if v.Indicators.Adjclose != nil {
			candle.Adjclose = v.Indicators.Adjclose[0].Adjclose[i]
		}
		candles = append(candles, candle)
	}

	r.Meta = v.Meta
	r.Candles = candles
	return nil
}
