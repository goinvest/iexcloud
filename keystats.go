// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// KeyStats models the data returned from IEX Cloud's /stats endpoint.
type KeyStats struct {
	Name                string  `json:"companyName"`
	MarketCap           int     `json:"marketCap"`
	Week52High          float64 `json:"week52High"`
	Week52Low           float64 `json:"week52Low"`
	Week52Change        float64 `json:"week52Change"`
	SharesOutstanding   int     `json:"sharesOutstanding"`
	Float               int     `json:"float"`
	Symbol              string  `json:"symbol"`
	Avg10Volume         int     `json:"avg10Volume"`
	Avg30Volume         int     `json:"avg30Volume"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
	Employees           int     `json:"employees"`
	MaxChangePercent    float64 `json:"maxChangePercent"`
	Year5ChangePercent  float64 `json:"year5ChangePercent"`
	Year2ChangePercent  float64 `json:"year2ChangePercent"`
	Year1ChangePercent  float64 `json:"year1ChangePercent"`
	YTDChangePercent    float64 `json:"ytdChangePercent"`
	Month6ChangePercent float64 `json:"month6ChangePercent"`
	Month3ChangePercent float64 `json:"month3ChangePercent"`
	Month1ChangePercent float64 `json:"month1ChangePercent"`
	Day30ChangePercent  float64 `json:"day30ChangePercent"`
	Day5ChangePercent   float64 `json:"day5ChangePercent"`
}

// KeyStats returns the key stats from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) KeyStats(stock string) (KeyStats, error) {
	stats := &KeyStats{}
	endpoint := fmt.Sprintf("/stock/%s/stats", url.PathEscape(stock))
	err := c.GetJSON(endpoint, stats)
	return *stats, err
}
