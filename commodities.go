// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// CommodityPrice models the price for a single commodity when returned using
// the time series endpoint.
type CommodityPrice struct {
	Value   float64   `json:"value"`
	ID      string    `json:"id"`
	Source  string    `json:"source"`
	Key     string    `json:"key"`
	Subkey  string    `json:"subkey"`
	Date    EpochTime `json:"date"`
	Updated EpochTime `json:"updated"`
}
