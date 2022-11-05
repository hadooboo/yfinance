package yfinance_test

import (
	"testing"

	"github.com/hadooboo/yfinance"
)

func TestClient_GetAutocomplete(t *testing.T) {
	type args struct {
		param *yfinance.GetAutocompleteParam
	}
	tests := []struct {
		name    string
		args    args
		want    func(*yfinance.Autocomplete) bool
		wantErr bool
	}{
		{
			name: "test-GetAutocomplete-success(query with apple)",
			args: args{
				param: &yfinance.GetAutocompleteParam{
					Query: "apple",
				},
			},
			want: func(a *yfinance.Autocomplete) bool {
				return len(a.Results) > 0
			},
		},
		{
			name: "test-GetAutocomplete-fail(nil param)",
			args: args{
				param: nil,
			},
			wantErr: true,
		},
		{
			name: "test-GetAutocomplete-fail(empty query)",
			args: args{
				param: &yfinance.GetAutocompleteParam{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetAutocomplete(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAutocomplete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !tt.want(got) {
				t.Errorf("Client.GetAutocomplete() = %v", got)
			}
		})
	}
}
