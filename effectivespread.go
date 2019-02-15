// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// EffectiveSpreads models the effective spreads.
type EffectiveSpreads []EffectiveSpread

// EffectiveSpread models the effective spread, eligible volume, and price
// improvement of a stock by market.
type EffectiveSpread struct {
	Volume           int     `json:"volume"`
	Venue            string  `json:"venue"`
	VenueName        string  `json:"venueName"`
	EffectiveSpread  float64 `json:"effectiveSpread"`
	EffectiveQuoted  float64 `json:"effectiveQuoted"`
	PriceImprovement float64 `json:"priceImprovement"`
}

// EffectiveSpreads returns the effective spreads from the IEX Cloud endpoint
// for the given stock symbol.
func (c Client) EffectiveSpreads(stock string) (EffectiveSpreads, error) {
	es := EffectiveSpreads{}
	endpoint := fmt.Sprintf("/stock/%s/effective-spread", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &es)
	return es, err
}
