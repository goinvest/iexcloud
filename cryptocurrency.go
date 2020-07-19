// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

type CryptoCurrencyBid struct {
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

type CryptoCurrencyAsk struct {
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

type Books struct {
	Bids []CryptoCurrencyBid `json:"bids"`
	Asks []CryptoCurrencyAsk `json:"asks"`
}

type Price struct {
	Price  string `json:"price"`
	Symbol string `json:"symbol"`
}

type Events struct {
	Symbol    string    `json:"symbol"`
	EventType string    `json:"eventType"`
	timestamp EpochTime `json:"timestamp"`
	Reason    string    `json:"reason"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Side      string    `json:"side"`
}

type CryptoCurrencyQuote struct {
	Symbol           string    `json:"symbol"`
	Sector           string    `json:"sector"`
	CalculationPrice string    `json:"calculationPrice"`
	LatestPrice      string    `json:"latestPrice"`
	LatestSource     string    `json:"latestSource"`
	LatestUpdate     EpochTime `json:"latestUpdate"`
	LatestVolume     string    `json:"latestVolume"`
	BidPrice         string    `json:"bidPrice"`
	BidSize          string    `json:"bidSize"`
	AskPrice         string    `json:"askPrice"`
	AskSize          string    `json:"askSize"`
	High             string    `json:"high"`
	Low              string    `json:"low"`
	PreviousClose    string    `json:"previousClose"`
}
