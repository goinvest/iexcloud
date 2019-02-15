// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Dividends models the dividens for a given range.
type Dividends []Dividend

// Dividend models one dividend.
type Dividend struct {
	ExDate       Date    `json:"exDate"`
	PaymentDate  Date    `json:"paymentDate"`
	RecordDate   Date    `json:"recordDate"`
	DeclaredDate Date    `json:"declaredDate"`
	Amount       float64 `json:"amount"`
	Flag         string  `json:"flag"`
}

// Dividends returns the dividends from the IEX Cloud endpoint for the given
// stock symbol and the given date range.
func (c Client) Dividends(stock string, r PathRange) (Dividends, error) {
	dividends := Dividends{}
	endpoint := fmt.Sprintf("/stock/%s/dividends/%s",
		url.PathEscape(stock), PathRangeJSON[r])
	err := c.GetJSON(endpoint, &dividends)
	return dividends, err
}
