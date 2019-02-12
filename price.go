// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Price returns the current stock price from the IEX Cloud endpoint for the
// given stock symbol.
func (c Client) Price(stock string) (float64, error) {
	endpoint := fmt.Sprintf("/stock/%s/price", url.PathEscape(stock))
	return c.GetFloat64(endpoint)
}
