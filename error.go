package yfinance

import "fmt"

const tag = "yfinance"

func createInvalidParameterError() error {
	return fmt.Errorf("%s: %s", tag, "invalid parameter")
}

func createInvalidResponseError() error {
	return fmt.Errorf("%s: %s", tag, "invalid response")
}
