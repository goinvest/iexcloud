// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// PriceTarget models the latest average, high, and low analyst price target for
// a symbol.
type PriceTarget struct {
	Symbol      string  `json:"symbol"`
	UpdatedDate Date    `json:"updatedDate"`
	Average     float64 `json:"priceTargetAverage"`
	High        float64 `json:"priceTargetHigh"`
	Low         float64 `json:"priceTargetLow"`
	NumAnalysts int     `json:"numberOfAnalysts"`
}

// PriceTarget returns the latest average, high, and low analyst price target
// for a given stock symbol.
func (c Client) PriceTarget(stock string) (PriceTarget, error) {
	pt := PriceTarget{}
	endpoint := fmt.Sprintf("/stock/%s/price-target", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &pt)
	return pt, err
}
