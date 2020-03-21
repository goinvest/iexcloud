// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// News models a news item either for the market or for an individual stock.
type News struct {
	Time       EpochTime `json:"datetime"`
	Headline   string    `json:"headline"`
	Source     string    `json:"source"`
	URL        string    `json:"url"`
	Summary    string    `json:"summary"`
	Related    string    `json:"related"`
	Image      string    `json:"image"`
	Language   string    `json:"lang"`
	HasPaywall bool      `json:"hasPaywall"`
}
