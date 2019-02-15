// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Estimates models the latest consensus esimtate for the next fiscal period.
type Estimates struct {
	Symbol    string     `json:"symbol"`
	Estimates []Estimate `json:"estimates"`
}

// Estimate models one estimate.
type Estimate struct {
	ConsensusEPS      float64 `json:"consensusEPS"`
	NumberOfEstimates int     `json:"numberOfEstimates"`
	FiscalPeriod      string  `json:"fiscalPeriod"`
	FiscalEndDate     Date    `json:"fiscalEndDate"`
	ReportDate        Date    `json:"reportDate"`
}

// Estimates returns the latest consensue estimates for the next fiscal period.
func (c Client) Estimates(stock string, num int) (Estimates, error) {
	estimates := Estimates{}
	endpoint := fmt.Sprintf("/stock/%s/estimates/%d", url.PathEscape(stock), num)
	err := c.GetJSON(endpoint, &estimates)
	return estimates, err
}
