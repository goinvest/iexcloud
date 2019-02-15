// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// OHLC models the open, high, low, close for a stock.
type OHLC struct {
	Open  OpenClose `json:"open"`
	Close OpenClose `json:"close"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
}

// OpenClose provides the price and time for either the open or close price of
// a stock.
type OpenClose struct {
	Price float64 `json:"price"`
	Time  int     `json:"Time"`
}

// OHLC returns the OHLC data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) OHLC(stock string) (OHLC, error) {
	ohlc := OHLC{}
	endpoint := fmt.Sprintf("/stock/%s/ohlc", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &ohlc)
	return ohlc, err
}
