// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// CryptoCurrencyBid models a cryptocurrency bid (price)
type CryptoCurrencyBid struct {
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

// CryptoCurrencyAsk models a cryptocurrency ask (price)
type CryptoCurrencyAsk struct {
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

// Books models cryptocurrency bids and asks (price)
type Books struct {
	Bids []CryptoCurrencyBid `json:"bids"`
	Asks []CryptoCurrencyAsk `json:"asks"`
}

// Price models a cryptocurrency price
type Price struct {
	Price  string `json:"price"`
	Symbol string `json:"symbol"`
}

// CryptoQuote models a quote for a cryptocurrency.
type CryptoQuote struct {
	Symbol           string    `json:"symbol"`
	Sector           string    `json:"sector"`
	CalculationPrice string    `json:"calculationPrice"`
	High             float64   `json:"high,string"`
	Low              float64   `json:"low,string"`
	LatestPrice      float64   `json:"latestPrice,string"`
	LatestSource     string    `json:"latestSource"`
	LatestUpdate     EpochTime `json:"latestUpdate"`
	LatestVolume     float64   `json:"latestVolume,string"`
	PreviousClose    float64   `json:"previousClose,string"`
	BidPrice         float64   `json:"bidPrice,string"`
	BidSize          float64   `json:"bidSize,string"`
	AskPrice         float64   `json:"askPrice,string"`
	AskSize          float64   `json:"askSize,string"`
}
