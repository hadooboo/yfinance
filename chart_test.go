package yfinance_test

import (
	"testing"

	"github.com/hadooboo/yfinance"
)

func TestClient_GetChart(t *testing.T) {
	type args struct {
		param *yfinance.GetChartParam
	}
	tests := []struct {
		name    string
		args    args
		want    func(*yfinance.Chart) bool
		wantErr bool
	}{
		{
			name: "test-GetChart-success(etf)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestETFSymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestETFSymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeETF
			},
		},
		{
			name: "test-GetChart-success(equity)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestEquitySymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestEquitySymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeEquity
			},
		},
		{
			name: "test-GetChart-success(mutual fund)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestMutualFundSymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestMutualFundSymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeMutualFund
			},
		},
		{
			name: "test-GetChart-success(currency)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestCurrencySymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestCurrencySymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeCurrency
			},
		},
		{
			name: "test-GetChart-success(crypto currency)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestCryptoCurrencySymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestCryptoCurrencySymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeCryptoCurrency
			},
		},
		{
			name: "test-GetChart-success(future)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestFutureSymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestFutureSymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeFuture
			},
		},
		{
			name: "test-GetChart-success(index)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker: TestIndexSymbol,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestIndexSymbol &&
					c.Meta.InstrumentType == yfinance.InstrumentTypeIndex
			},
		},
		{
			name: "test-GetChart-success(auto adjust)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker:     TestEquitySymbol,
					Range:      yfinance.DataRange1mo,
					Interval:   yfinance.DataGranularity1d,
					AutoAdjust: true,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestEquitySymbol
			},
		},
		{
			name: "test-GetChart-success(back adjust)",
			args: args{
				param: &yfinance.GetChartParam{
					Ticker:     TestEquitySymbol,
					Range:      yfinance.DataRange1mo,
					Interval:   yfinance.DataGranularity1d,
					BackAdjust: true,
				},
			},
			want: func(c *yfinance.Chart) bool {
				return c.Meta.Symbol == TestEquitySymbol
			},
		},
		{
			name: "test-GetChart-fail(nil param)",
			args: args{
				param: nil,
			},
			wantErr: true,
		},
		{
			name: "test-GetChart-fail(empty ticker)",
			args: args{
				param: &yfinance.GetChartParam{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetChart(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetChart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !tt.want(got) {
				t.Errorf("Client.GetChart() = %v", got)
			}
		})
	}
}
