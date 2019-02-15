// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Earnings provides earnings data for a given company including the actual
// EPS, consensus, and fiscal period. Earnings are available quarterly (last 4
// quarters) and annually (last 4 years).
type Earnings struct {
	Symbol   string    `json:"symbol"`
	Earnings []Earning `json:"earnings"`
}

// Earning models the earnings for one date.
type Earning struct {
	ActualEPS            float64      `json:"actualEPS"`
	ConsensusEPS         float64      `json:"consensusEPS"`
	AnnounceTime         AnnounceTime `json:"announcetime"`
	NumberOfEstimates    int          `json:"numberOfEstimates"`
	EPSSurpriseDollar    float64      `json:"EPSSurpriseDollar"`
	EPSReportDate        Date         `json:"EPSReportDate"`
	FiscalPeriod         string       `json:"fiscalPeriod"`
	FiscalEndDate        Date         `json:"fiscalEndDate"`
	YearAgo              float64      `json:"yearAgo"`
	YearAgoChangePercent float64      `json:"yearAgoChangePercent"`
}

// AnnounceTime refers to the time of earnings announcement.
type AnnounceTime int

const (
	bto AnnounceTime = iota
	dmt
	amc
)

var announceTimeDescription = map[AnnounceTime]string{
	bto: "Before open",
	dmt: "During trading",
	amc: "After close",
}

// AnnounceTimes maps the string keys from the JSON to the AnnounceType
// constant values.
var AnnounceTimes = map[string]AnnounceTime{
	"BTO": bto,
	"DMT": dmt,
	"AMC": amc,
}

// AnnounceTimeJSON maps an AnnounceTime to the string used in the JSON.
var AnnounceTimeJSON = map[AnnounceTime]string{
	bto: "BTO",
	dmt: "DMT",
	amc: "AMC",
}

// UnmarshalJSON implements the Unmarshaler interface for AnnounceTime.
func (a *AnnounceTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("announceTime should be a string, got %s", data)
	}
	return a.Set(s)
}

// Set sets the issue type using a string.
func (a *AnnounceTime) Set(s string) error {
	// Ensure the provided string matches on the keys in the map.
	got, ok := AnnounceTimes[s]
	if !ok {
		return fmt.Errorf("invalid issue type %q", s)
	}
	// Set the issue type to the value found in the map per the key.
	*a = got
	return nil
}

// MarshalJSON implements the Marshaler interface for AnnounceTime.
func (a *AnnounceTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(AnnounceTimeJSON[*a])
}

// String implements the Stringer interface for AnnounceTime.
func (a AnnounceTime) String() string {
	return announceTimeDescription[a]
}

// Earnings returns the specified number of most recent earnings data from the
// IEX Cloud endpoint for the given stock symbol.
func (c Client) Earnings(stock string, num int) (Earnings, error) {
	earnings := Earnings{}
	endpoint := fmt.Sprintf("/stock/%s/earnings/%d", url.PathEscape(stock), num)
	err := c.GetJSON(endpoint, &earnings)
	return earnings, err
}
