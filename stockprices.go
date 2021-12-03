// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Book models the data returned from the /book endpoint.
type Book struct {
	Quote       Quote       `json:"quote"`
	Bids        []BidAsk    `json:"bids"`
	Asks        []BidAsk    `json:"asks"`
	Trades      []Trade     `json:"trades"`
	SystemEvent SystemEvent `json:"systemEvent"`
}

// BidAsk models a bid or an ask for a quote.
type BidAsk struct {
	Price     float64   `json:"price"`
	Size      int       `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

// DelayedQuote returns the 15 minute delayed market quote.
type DelayedQuote struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float64 `json:"delayedPrice"`
	DelayedSize      int     `json:"delayedSize"`
	DelayedPriceTime int     `json:"delayedPriceTime"`
	High             float64 `json:"High"`
	Low              float64 `json:"Low"`
	TotalVolume      int     `json:"totalVolume"`
	ProcessedTime    int     `json:"processedTime"`
}

// HistoricalPrice models the data for a historical stock price.
type HistoricalPrice struct {
	Date string `json:"date"`
}

// IntradayPrice models the data for an aggregated intraday price in one minute
// buckets.
type IntradayPrice struct {
	Date                 Date       `json:"date"`
	Minute               HourMinute `json:"minute"`
	Label                string     `json:"label"`
	MarketOpen           float64    `json:"marketOpen"`
	MarketClose          float64    `json:"marketClose"`
	MarketHigh           float64    `json:"marketHigh"`
	MarketLow            float64    `json:"marketLow"`
	MarketAverage        float64    `json:"marketAverage"`
	MarketVolume         int        `json:"marketVolume"`
	MarketNotional       float64    `json:"marketNotional"`
	MarketNumTrades      int        `json:"marketNumberOfTrades"`
	MarketChangeOverTime float64    `json:"marketChangeOverTime"`
	High                 float64    `json:"High"`
	Low                  float64    `json:"Low"`
	Open                 float64    `json:"Open"`
	Close                float64    `json:"Close"`
	Average              float64    `json:"average"`
	Volume               int        `json:"volume"`
	Notional             float64    `json:"notional"`
	NumTrades            int        `json:"numberOfTrades"`
	ChangeOverTime       float64    `json:"changeOverTime"`
}

// LargestTrade models the 15 minute delayed, last sale eligible trades.
type LargestTrade struct {
	Price     float64 `json:"price"`
	Size      int     `json:"size"`
	Time      int     `json:"time"`
	TimeLabel string  `json:"timeLabel"`
	Venue     string  `json:"venue"`
	VenueName string  `json:"venueName"`
}

// OpenClose provides the price and time for either the open or close price of
// a stock.
type OpenClose struct {
	Price float64 `json:"price"`
	Time  int     `json:"Time"`
}

// OHLC models the open, high, low, close for a stock.
type OHLC struct {
	Open  OpenClose `json:"open"`
	Close OpenClose `json:"close"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
}

// PreviousDay models the previous day adjusted price data.
type PreviousDay struct {
	Symbol           string  `json:"symbol"`
	Date             Date    `json:"date"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"Low"`
	Close            float64 `json:"close"`
	Volume           int     `json:"volume"`
	UnadjustedVolume int     `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
}

// Quote models the data returned from the IEX Cloud /quote endpoint.
type Quote struct {
	Symbol                string    `json:"symbol,omitempty"`
	CompanyName           string    `json:"companyName,omitempty"`
	CalculationPrice      string    `json:"calculationPrice,omitempty"`
	Open                  float64   `json:"open,omitempty"`
	OpenTime              EpochTime `json:"openTime,omitempty"`
	Close                 float64   `json:"close,omitempty"`
	CloseTime             EpochTime `json:"closeTime,omitempty"`
	High                  float64   `json:"high,omitempty"`
	Low                   float64   `json:"low,omitempty"`
	LatestPrice           float64   `json:"latestPrice,omitempty"`
	LatestSource          string    `json:"latestSource,omitempty"`
	LatestTime            string    `json:"latestTime,omitempty"`
	LatestUpdate          EpochTime `json:"latestUpdate,omitempty"`
	LatestVolume          int       `json:"latestVolume,omitempty"`
	IEXRealtimePrice      float64   `json:"iexRealtimePrice,omitempty"`
	IEXRealtimeSize       int       `json:"iexRealtimeSize,omitempty"`
	IEXLastUpdated        EpochTime `json:"iexLastUpdated,omitempty"`
	DelayedPrice          float64   `json:"delayedPrice,omitempty"`
	DelayedPriceTime      EpochTime `json:"delayedPriceTime,omitempty"`
	ExtendedPrice         float64   `json:"extendedPrice,omitempty"`
	ExtendedChange        float64   `json:"extendedChange,omitempty"`
	ExtendedChangePercent float64   `json:"extendedChangePercent,omitempty"`
	ExtendedPriceTime     EpochTime `json:"extendedPriceTime,omitempty"`
	PreviousClose         float64   `json:"previousClose,omitempty"`
	Change                float64   `json:"change,omitempty"`
	ChangePercent         float64   `json:"changePercent,omitempty"`
	IEXMarketPercent      float64   `json:"iexMarketPercent,omitempty"`
	IEXVolume             int       `json:"iexVolume,omitempty"`
	AvgTotalVolume        int       `json:"avgTotalVolume,omitempty"`
	IEXBidPrice           float64   `json:"iexBidPrice,omitempty"`
	IEXBidSize            int       `json:"iexBidSize,omitempty"`
	IEXAskPrice           float64   `json:"iexAskPrice,omitempty"`
	IEXAskSize            int       `json:"iexAskSize,omitempty"`
	MarketCap             int       `json:"marketCap,omitempty"`
	Week52High            float64   `json:"week52High,omitempty"`
	Week52Low             float64   `json:"week52Low,omitempty"`
	YTDChange             float64   `json:"ytdChange,omitempty"`
	PERatio               float64   `json:"peRatio,omitempty"`
}

// VenueVolume models the 15 minute delayed and 30 day average consolidated
// volume percentage of a stock by market.
type VenueVolume struct {
	Volume               int     `json:"volume"`
	Venue                string  `json:"venue"`
	VenueName            string  `json:"venueName"`
	Date                 Date    `json:"date"`
	MarketPercent        float64 `json:"marketPercent"`
	AverageMarketPercent float64 `json:"avgMarketPercent"`
}
