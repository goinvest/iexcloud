// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// SymbolDetails models the details of a symbol.
type SymbolDetails struct {
	Symbol   string `json:"symbol"`
	Region   string `json:"region"`
	Exchange string `json:"exchange"`
	IEXId    string `json:"iexid"`
}
