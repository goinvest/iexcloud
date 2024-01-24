// Copyright (c) 2019-2024 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// TradedSymbol models a stock symbol the Investors Exchange supports for
// trading.
type TradedSymbol struct {
	Symbol    string `json:"symbol"`
	Date      Date   `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
}

// Symbol models the data for one stock, mutual fund, or OTC symbol that IEX
// Cloud supports for API calls.
type Symbol struct {
	Symbol    string `json:"symbol"`
	Exchange  string `json:"exchange"`
	Name      string `json:"name"`
	Date      Date   `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
	Type      string `json:"type"`
	Region    string `json:"region"`
	Currency  string `json:"currency"`
	IEXID     string `json:"iexId"`
	FIGI      string `json:"figi"`
	CIK       string `json:"cik"`
}

// SearchResult models the data for a single search result.
type SearchResult struct {
	Symbol       string `json:"symbol"`
	SecurityName string `json:"securityName"`
	SecurityType string `json:"securityType"`
	Region       string `json:"region"`
	Exchange     string `json:"exchange"`
	Sector       string `json:"sector"`
	CIK          string `json:"cik"`
}

// FIGIResult models the data for a single FIGI mapping search result.
type FIGIResult struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Region   string `json:"region"`
	IEXID    string `json:"iexId"`
}

// USExchange provides information about one U.S. exchange including the name,
// the Market identifier code, the ID used to identify the exchange on the
// consolidated tape, the FINRA OATS exchange participant ID, and the type of
// securities traded by the exchange.
type USExchange struct {
	Name     string `json:"name"`
	LongName string `json:"longName"`
	MarketID string `json:"mic"`
	TapeID   string `json:"tapeId"`
	OATSID   string `json:"oatsId"`
	RefID    string `json:"refId"`
	Type     string `json:"type"`
}

// TradeHolidayDate models either a trade date or a holiday.
type TradeHolidayDate struct {
	Date           Date `json:"date"`
	SettlementDate Date `json:"settlementDate"`
}

// FXSymbols provides a list of the currencies and currency pairs available
// from IEX Cloud.
type FXSymbols struct {
	Currencies []Currency     `json:"currencies"`
	Pairs      []CurrencyPair `json:"pairs"`
}

// Currency models the code and name for a currency.
type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// CurrencyPair models an available currency pair listing both the from
// currency and the to currency codes.
type CurrencyPair struct {
	From string `json:"fromCurrency"`
	To   string `json:"toCurrency"`
}

// Sector models an industry sector, as defined by IEX.
// i.e. "Technology", "Consumer Cyclical"
type Sector struct {
	Name string `json:"name"`
}

// Tag models the tag field specified for each symbol
// i.e. "Financial Services", "Industrials"
type Tag struct {
	Name string `json:"name"`
}
