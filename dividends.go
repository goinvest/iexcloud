// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Dividends models the dividens for a given range.
type Dividends []Dividend

// Dividend models one dividend.
type Dividend struct {
	ExDate       Date    `json:"exDate"`
	PaymentDate  Date    `json:"paymentDate"`
	RecordDate   Date    `json:"recordDate"`
	DeclaredDeta Date    `json:"declaredDate"`
	Amount       float64 `json:"amount"`
	Flag         string  `json:"flag"`
}
