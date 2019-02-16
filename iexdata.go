// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Last provides trade data for executions on IEX. It is a near real time,
// intraday API that provides IEX last sale price, size and time. Last is ideal
// for developers that need a lightweight stock quote.
type Last struct {
	Symbol string    `json:"symbol"`
	Price  float64   `json:"Price"`
	Size   int       `json:"Size"`
	Time   EpochTime `json:"time"`
}

// Records models the stats records.
type Records struct {
	Volume VolumeRecord `json:"volume"`
}

// TOPS contain IEX's aggregated best quoted bid and offer position in near
// real time for all securities on IEX's displayed limit order book.
type TOPS struct {
	Symbol        string    `json:"symbol"`
	MarketPercent float64   `json:"marketPercent"`
	BidPrice      float64   `json:"bidPrice"`
	AskSize       int       `json:"AskSize"`
	AskPrice      float64   `json:"AskPrice"`
	Volume        int       `json:"volume"`
	LastSalePrice float64   `json:"lastSalePrice"`
	LastSaleTime  EpochTime `json:"lastSaleTime"`
	LastUpdated   EpochTime `json:"lastUpdated"`
	Sector        string    `json:"sector"`
	SecurityType  string    `json:"securityType"`
}

// VolumeRecord models the record volume.
type VolumeRecord struct {
	Value            float64 `json:"recordValue"`
	Date             Date    `json:"recordDate"`
	PreviousDayValue float64 `json:"previousDayValue"`
	Avg30Value       float64 `json:"avg30Value"`
}
