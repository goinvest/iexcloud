// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Symbol models the data for one stock symbol.
type Symbol struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      Date   `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
	Type      string `json:"type"`
	IEXID     string `json:"iexId"`
}

// Symbols returns an array of symbols that IEX Cloud supports for API calls.
func (c Client) Symbols() ([]Symbol, error) {
	symbols := &[]Symbol{}
	endpoint := "/ref-data/symbols"
	err := c.GetJSON(endpoint, symbols)
	return *symbols, err
}
