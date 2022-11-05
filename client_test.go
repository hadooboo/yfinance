package yfinance_test

import (
	"os"
	"testing"

	"github.com/hadooboo/yfinance"
)

const (
	TestETFSymbol            = "VOO"
	TestEquitySymbol         = "AAPL"
	TestMutualFundSymbol     = "VFINX"
	TestCurrencySymbol       = "KRW=X"
	TestCryptoCurrencySymbol = "BTC-USD"
	TestFutureSymbol         = "GC=F"
	TestIndexSymbol          = "^GSPC"
)

var client *yfinance.Client

func TestMain(m *testing.M) {
	client = yfinance.New()
	code := m.Run()
	os.Exit(code)
}
