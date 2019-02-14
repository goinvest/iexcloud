// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// LargestTrades models multiple last sale eligible trades.
type LargestTrades []LargestTrade

// LargestTrade models the 15 minute delayed, last sale eligible trades.
type LargestTrade struct {
	Price     float64 `json:"price"`
	Size      int     `json:"size"`
	Time      int     `json:"time"`
	TimeLabel string  `json:"timeLabel"`
	Venue     string  `json:"venue"`
	VenueName string  `json:"venueName"`
}

// LargestTrades returns the 15 minute delayed, last sale eligible trade from
// the IEX Cloud endpoint for the given stock symbol.
func (c Client) LargestTrades(stock string) (LargestTrades, error) {
	lt := &LargestTrades{}
	endpoint := fmt.Sprintf("/stock/%s/largest-trades", url.PathEscape(stock))
	err := c.GetJSON(endpoint, lt)
	return *lt, err
}
