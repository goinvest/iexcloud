// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Quote models the data returned from the IEX Cloud /quote endpoint.
type Quote struct {
	Symbol                string     `json:"symbol"`
	CompanyName           string     `json:"companyName"`
	CalculationPrice      string     `json:"calculationPrice"`
	Open                  float64    `json:"open"`
	OpenTime              EpochTime  `json:"openTime"`
	Close                 float64    `json:"close"`
	CloseTime             EpochTime  `json:"closeTime"`
	High                  float64    `json:"high"`
	Low                   float64    `json:"low"`
	LatestPrice           float64    `json:"latestPrice"`
	LatestSource          string     `json:"latestSource"`
	LatestTime            string     `json:"latestTime"`
	LatestUpdate          EpochTime  `json:"latestUpdate"`
	LatestVolume          int        `json:"latestVolume"`
	IEXRealtimePrice      float64    `json:"iexRealtimePrice"`
	IEXRealtimeSize       int        `json:"iexRealtimeSize"`
	IEXLastUpdated        *EpochTime `json:"iexLastUpdated"`
	DelayedPrice          float64    `json:"delayedPrice"`
	DelayedTime           string     `json:"delayedTime"`
	ExtendedPrice         float64    `json:"extendedPrice"`
	ExtendedChange        float64    `json:"extendedChange"`
	ExtendedChangePercent float64    `json:"extendedChangePercent"`
	ExtendedPriceTime     int        `json:"extendedPriceTime"`
	PreviousClose         float64    `json:"previousClose"`
	Change                float64    `json:"change"`
	ChangePercent         float64    `json:"changePercent"`
	IEXMarketPercent      float64    `json:"iexMarketPercent"`
	IEXVolume             int        `json:"iexVolume"`
	AvgTotalVolume        int        `json:"avgTotalVolume"`
	IEXBidPrice           float64    `json:"iexBidPrice"`
	IEXBidSize            int        `json:"iexBidSize"`
	IEXAskPrice           float64    `json:"iexAskPrice"`
	IEXAskSize            int        `json:"iexAskSize"`
	MarketCap             int        `json:"marketCap"`
	Week52High            float64    `json:"week52High"`
	Week52Low             float64    `json:"week52Low"`
	YTDChange             float64    `json:"ytdChange"`
}

// Quote returns the quote data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) Quote(stock string) (Quote, error) {
	quote := Quote{}
	endpoint := fmt.Sprintf("/stock/%s/quote", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &quote)
	return quote, err
}
