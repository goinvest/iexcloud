// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
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
	Symbol                string    `json:"symbol"`
	CompanyName           string    `json:"companyName"`
	PrimaryExchange       string    `json:"primaryExchange"`
	Sector                string    `json:"sector"`
	CalculationPrice      string    `json:"calculationPrice"`
	Open                  float64   `json:"open"`
	OpenTime              EpochTime `json:"openTime"`
	Close                 float64   `json:"close"`
	CloseTime             EpochTime `json:"closeTime"`
	High                  float64   `json:"high"`
	Low                   float64   `json:"low"`
	LatestPrice           float64   `json:"latestPrice"`
	LatestSource          string    `json:"latestSource"`
	LatestTime            string    `json:"latestTime"`
	LatestUpdate          EpochTime `json:"latestUpdate"`
	LatestVolume          float64   `json:"latestVolume"`
	IEXRealtimePrice      float64   `json:"iexRealtimePrice"`
	IEXRealtimeSize       int       `json:"iexRealtimeSize"`
	IEXLastUpdated        EpochTime `json:"iexLastUpdated"`
	DelayedPrice          float64   `json:"delayedPrice"`
	DelayedPriceTime      EpochTime `json:"delayedPriceTime"`
	ExtendedPrice         float64   `json:"extendedPrice"`
	ExtendedChange        float64   `json:"extendedChange"`
	ExtendedChangePercent float64   `json:"extendedChangePercent"`
	ExtendedPriceTime     EpochTime `json:"extendedPriceTime"`
	PreviousClose         float64   `json:"previousClose"`
	Change                float64   `json:"change"`
	ChangePercent         float64   `json:"changePercent"`
	IEXMarketPercent      float64   `json:"iexMarketPercent"`
	IEXVolume             int       `json:"iexVolume"`
	AvgTotalVolume        int       `json:"avgTotalVolume"`
	IEXBidPrice           float64   `json:"iexBidPrice"`
	IEXBidSize            int       `json:"iexBidSize"`
	IEXAskPrice           float64   `json:"iexAskPrice"`
	IEXAskSize            int       `json:"iexAskSize"`
	MarketCap             int       `json:"marketCap"`
	PERatio               float64   `json:"peRatio"`
	Week52High            float64   `json:"week52High"`
	Week52Low             float64   `json:"week52Low"`
	YTDChange             float64   `json:"ytdChange"`
	BidPrice              float64   `json:"bidPrice"`
	BidSize               float64   `json:"bidSize"`
	AskPrice              float64   `json:"askPrice"`
	AskSize               float64   `json:"askSize"`
}
