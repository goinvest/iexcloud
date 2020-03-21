// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
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

// TOPS contains IEX's aggregated best quoted bid and offer position in near
// real time for all securities on IEX's displayed limit order book.
type TOPS struct {
	Symbol        string    `json:"symbol"`
	MarketPercent float64   `json:"marketPercent"`
	BidSize       int       `json:"bidSize"`
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

// DEEP is used to receive real-time depth of book quotations direct from IEX.
// The depth of book quotations received via DEEP provide an aggregated size of
// resting displayed orders at a price and side, and do not indicate the size or
// number of individual orders at any price level.
type DEEP struct {
	Symbol        string        `json:"symbol"`
	MarketPercent float64       `json:"marketPercent"`
	Volume        int           `json:"volume"`
	LastSalePrice float64       `json:"lastSalePrice"`
	LastSaleSize  int           `json:"lastSaleSize"`
	LastSaleTime  EpochTime     `json:"lastSaleTime"`
	LastUpdated   EpochTime     `json:"lastUpdated"`
	Bids          []BidAsk      `json:"bids"`
	Asks          []BidAsk      `json:"asks"`
	SystemEvent   SystemEvent   `json:"systemEvent"`
	TradingStatus TradingStatus `json:"tradingStatus"`
	OpHaltStatus  OpHaltStatus  `json:"opHaltStatus"`
	SSRStatus     SSRStatus     `json:"ssrStatus"`
	SecurityEvent SecurityEvent `json:"securityEvent"`
	Trades        []Trade       `json:"trades"`
	TradeBreaks   []Trade       `json:"tradeBreaks"`
	Auction       Auction       `json:"auction"`
}

// Auction models auction data for a security
type Auction struct {
	AuctionType          string    `json:"auctionType"`
	PairedShares         int       `json:"pairedShares"`
	ImbalanceShares      int       `json:"imbalanceShares"`
	ReferencePrice       float64   `json:"referencePrice"`
	IndicativePrice      float64   `json:"indicativePrice"`
	AuctionBookPrice     float64   `json:"auctionBookPrice"`
	CollarReferencePrice float64   `json:"collarReferencePrice"`
	LowerCollarPrice     float64   `json:"lowerCollarPrice"`
	UpperCollarPrice     float64   `json:"upperCollarPrice"`
	ExtensionNumber      int       `json:"extensionNumber"`
	StartTime            EpochTime `json:"startTime"`
	LastUpdate           EpochTime `json:"lastUpdate"`
}

// DEEPBook contains just the bids and asks for a specified symbol
type DEEPBook struct {
	Bids []BidAsk `json:"bids"`
	Asks []BidAsk `json:"asks"`
}

// OpHaltStatus models the operational halt status of a security
type OpHaltStatus struct {
	IsHalted  bool      `json:"isHalted"`
	Timestamp EpochTime `json:"timestamp"`
}

// SecurityEvent models events which apply to a specific security
type SecurityEvent struct {
	SecurityEvent string    `json:"securityEvent"`
	Timestamp     EpochTime `json:"timestamp"`
}

// SSRStatus models the short sale price test status for a security
type SSRStatus struct {
	IsSSR     bool      `json:"isSSR"`
	Detail    string    `json:"detail"`
	Timestamp EpochTime `json:"timestamp"`
}

// SystemEvent models a system event for a quote.
type SystemEvent struct {
	SystemEvent string    `json:"systemEvent"`
	Timestamp   EpochTime `json:"timestamp"`
}

// Trade models a trade for a quote.
type Trade struct {
	Price                 float64   `json:"price"`
	Size                  int       `json:"size"`
	TradeID               int       `json:"tradeId"`
	IsISO                 bool      `json:"isISO"`
	IsOddLot              bool      `json:"isOddLot"`
	IsOutsideRegularHours bool      `json:"isOutsideRegularHours"`
	IsSinglePriceCross    bool      `json:"isSinglePriceCross"`
	IsTradeThroughExempt  bool      `json:"isTradeThroughExempt"`
	Timestamp             EpochTime `json:"timestamp"`
}

// TradingStatus models the current trading status of a security
type TradingStatus struct {
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	Timestamp EpochTime `json:"timestamp"`
}

// VolumeRecord models the record volume.
type VolumeRecord struct {
	Value            float64 `json:"recordValue"`
	Date             Date    `json:"recordDate"`
	PreviousDayValue float64 `json:"previousDayValue"`
	Avg30Value       float64 `json:"avg30Value"`
}

// IntradayStats models the intraday stats on IEX.
type IntradayStats struct {
	Volume        Stat `json:"volume"`
	SymbolsTraded Stat `json:"symbolsTraded"`
	RoutedVolume  Stat `json:"routedVolume"`
	Notional      Stat `json:"notional"`
	MarketShare   Stat `json:"marketShare"`
}

// Stat models a single stat.
type Stat struct {
	Value       float64   `json:"value"`
	LastUpdated EpochTime `json:"lastUpdated"`
}
