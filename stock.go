// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// CoreEstimate modules a current or historical consensus analyst
// recommendation and price target.
type CoreEstimate struct {
	ID              string    `json:"CORE_ESTIMATES"`
	Key             string    `json:"key"`
	Subkey          string    `json:"subkey"`
	Symbol          string    `json:"symbol"`
	AnalystCount    int       `json:"analystCount"`
	ConsensusDate   Date      `json:"consensusDate"`
	MarketConsensus float64   `json:"marketConsensus"`
	TargetPrice     float64   `json:"marketConsensusTargetPrice"`
	Date            EpochTime `json:"date"`
	Updated         EpochTime `json:"updated"`
}
