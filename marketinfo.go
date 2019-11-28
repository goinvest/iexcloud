// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// EarningsToday models the earning that will be reported today as two arrays:
// before the open and after market close. Each array contains an object with
// all keys from earnings, a quote object, and a headline key.
type EarningsToday struct {
	BeforeOpen []TodayEarning `json:"bto"`
	AfterClose []TodayEarning `json:"amc"`
}

// TodayEarning models a single earning being reported today containing all
// keys from earnings, a quote object, and a headline.
type TodayEarning struct {
	Earning
	EstimatedChangePercent float64 `json:"estimatedChangePercent"`
	SymbolID               int     `json:"symbolId"`
	Symbol                 string  `json:"symbol"`
	Quote                  Quote   `json:"quote"`
	Headline               string  `json:"headline"`
}

// Market models the traded volume on U.S. markets.
type Market struct {
	MIC         string    `json:"mic"`
	TapeID      string    `json:"tapeId"`
	Venue       string    `json:"venueName"`
	Volume      int       `json:"volume"`
	TapeA       int       `json:"tapeA"`
	TapeB       int       `json:"tapeB"`
	TapeC       int       `json:"tapeC"`
	Percent     float64   `json:"marketPercent"`
	LastUpdated EpochTime `json:"lastUpdated"`
}

// SectorPerformance models the performance based on each sector ETF.
type SectorPerformance struct {
	Type        string    `json:"sector"`
	Name        string    `json:"name"`
	Performance float64   `json:"performance"`
	LastUpdated EpochTime `json:"lastUpdated"`
}
