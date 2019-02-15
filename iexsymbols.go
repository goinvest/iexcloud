// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// IEXSymbol models the data for one stock symbol.
type IEXSymbol struct {
	Symbol    string `json:"symbol"`
	Date      Date   `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
}

// IEXSymbols returns an array of symbols that IEX Cloud supports for API calls.
func (c Client) IEXSymbols() ([]IEXSymbol, error) {
	symbols := []IEXSymbol{}
	endpoint := "/ref-data/iex/symbols"
	err := c.GetJSON(endpoint, &symbols)
	return symbols, err
}
