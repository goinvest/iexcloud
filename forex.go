// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// ExchangeRate models the exchange rate of a given currency pair.
type ExchangeRate struct {
	Date         Date    `json:"date"`
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	Rate         float64 `json:"rate"`
}

// CurrencyRate returns real-time foreign currency exchange rates data.
type CurrencyRate struct {
	Symbol    string    `json:"symbol"`
	Rate      float64   `json:"rate"`
	Timestamp EpochTime `json:"timestamp"`
}
