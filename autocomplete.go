package yfinance

import (
	"net/http"
	"net/url"
)

const getAutocompletePath = "/v7/finance/autocomplete"

type GetAutocompleteParam struct {
	Query string
}

func (r GetAutocompleteParam) createURL() string {
	queryParam := make(url.Values)
	queryParam.Add("query", r.Query)
	queryParam.Add("lang", "en-US")

	return yahooFinanceAPIURL + getAutocompletePath + "?" + queryParam.Encode()
}

func (r *Client) GetAutocomplete(param *GetAutocompleteParam) (*Autocomplete, error) {
	if param == nil || len(param.Query) == 0 {
		return nil, createInvalidParameterError()
	}

	req, err := http.NewRequest(http.MethodGet, param.createURL(), nil)
	if err != nil {
		return nil, err
	}

	var res autocomplete
	if err := r.do(req, &res); err != nil {
		return nil, err
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return res.ToAutocomplete()
}

type autocomplete struct {
	ResultSet *struct {
		Query  string               `json:"Query"`
		Result []AutocompleteResult `json:"Result"`
	} `json:"ResultSet,omitempty"`
	Error *struct {
		Result any    `json:"result"`
		Error  *Error `json:"error,omitempty"`
	} `json:"error,omitempty"`
}

func (r autocomplete) Err() error {
	if r.Error != nil && r.Error.Error != nil {
		return r.Error.Error
	}

	return nil
}

func (r autocomplete) ToAutocomplete() (*Autocomplete, error) {
	return &Autocomplete{
		Results: r.ResultSet.Result,
	}, nil
}
