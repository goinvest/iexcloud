// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// PreviousDay models the previous day adjusted price data.
type PreviousDay struct {
	Symbol           string  `json:"symbol"`
	Date             Date    `json:"date"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"Low"`
	Close            float64 `json:"close"`
	Volume           int     `json:"volume"`
	UnadjustedVolume int     `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
}

// PreviousDay returns the previous day adjusted price data from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) PreviousDay(stock string) (PreviousDay, error) {
	pd := &PreviousDay{}
	endpoint := fmt.Sprintf("/stock/%s/previous", url.PathEscape(stock))
	err := c.GetJSON(endpoint, pd)
	return *pd, err
}
