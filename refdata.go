// Copyright (c) 2019 The iexcloud developers. All rights reserved.
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

// Symbol models the data for one stock symbol that IEX Cloud supports for API
// calls.
type Symbol struct {
	TradedSymbol
	Name  string `json:"name"`
	Type  string `json:"type"`
	IEXID string `json:"iexId"`
}

// USExchange provides information about one U.S. exchange including the name,
// the Market identifier code, the ID used to identify the exchange on the
// consolidated tape, the FINRA OATS exchange participant ID, and the type of
// securities traded by the exchange.
type USExchange struct {
	Name     string `json:"name"`
	MarketID int    `json:"mic"`
	TapeID   string `json:"tapeId"`
	OATSID   string `json:"oatsId"`
	Type     string `json:"type"`
}

// OTCSymbol models an Over-the-Counter (OTC) symbol that IEX Cloud supports
// for API calls.
type OTCSymbol struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Date   Date   `json:"date"`
	Type   string `json:"type"`
	IEXID  string `json:"iexId"`
}

// FXSymbols provides a list of the currencies and currency pairs available
// from IEX Cloud.
type FXSymbols struct {
	Currencies []Currency     `json:"currencies"`
	Pairs      []CurrencyPair `json:"pairs"`
}

// CryptoSymbol models cryptocurrency symbol that IEX Cloud supports for API
// calls.
type CryptoSymbol struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      Date   `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
	Type      string `json:"type"`
	IEXID     string `json:"iexId"`
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
