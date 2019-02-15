// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// News models a news item either for the market or for an individual stock.
type News struct {
	Time       EpochTime `json:"datetime"`
	Headline   string    `json:"headline"`
	Source     string    `json:"source"`
	URL        string    `json:"url"`
	Summary    string    `json:"summary"`
	Related    string    `json:"related"`
	Image      string    `json:"image"`
	Language   string    `json:"lang"`
	HasPaywall bool      `json:"hasPaywall"`
}

// News retrieves the given number of news articles for the given stock symbol.
func (c Client) News(stock string, num int) ([]News, error) {
	n := &[]News{}
	endpoint := fmt.Sprintf("/stock/%s/news/last/%d",
		url.PathEscape(stock), num)
	err := c.GetJSON(endpoint, n)
	return *n, err
}

// MarketNews retrieves the given number of news articles for the market.
func (c Client) MarketNews(num int) ([]News, error) {
	n := &[]News{}
	endpoint := fmt.Sprintf("/stock/market/news/last/%d", num)
	err := c.GetJSON(endpoint, n)
	return *n, err
}
