package yfinance_test

import (
	"os"
	"testing"

	"github.com/hadooboo/yfinance"
	"github.com/stretchr/testify/assert"
)

var client *yfinance.Client

func TestMain(m *testing.M) {
	client = yfinance.New()
	code := m.Run()
	os.Exit(code)
}

func TestGetChart(t *testing.T) {
	res, err := client.GetChart(&yfinance.GetChartParams{
		Range:    yfinance.DataRange1y,
		Interval: yfinance.DataGranularity1d,
		Ticker:   "AAPL",
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)

	res, err = client.GetChart(&yfinance.GetChartParams{
		Range:    yfinance.DataRangeMax,
		Interval: yfinance.DataGranularity1mo,
		Ticker:   "MSFT",
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetAutocompleteResultSet(t *testing.T) {
	res, err := client.GetAutocompleteResultSet(&yfinance.GetAutocompleteResultSetParams{
		Query: "apple",
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
