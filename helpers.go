// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date models a report date
type Date time.Time

// UnmarshalJSON implements the Unmarshaler interface for ReportDate.
func (d *Date) UnmarshalJSON(data []byte) error {
	var aux string
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return fmt.Errorf("error unmarshaling date to string: %s", err)
	}
	t, err := time.Parse("2006-01-02", aux)
	if err != nil {
		return fmt.Errorf("error converting string to date: %s", err)
	}
	*d = Date(t)
	return nil
}

// MarshalJSON implements the Marshaler interface for ReportDate.
func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	return json.Marshal(t.Format("2006-01-02"))
}

// PathRange refers to the date range used in the path of an endpoint.
type PathRange int

const (
	Mo1 PathRange = iota
	Mo3
	Mo6
	Yr1
	Yr2
	Yr5
	YTD
	Next
)

var pathRangeDescription = map[PathRange]string{
	Mo1:  "One month (default)",
	Mo3:  "Three months",
	Mo6:  "Six months",
	Yr1:  "One year",
	Yr2:  "Two years",
	Yr5:  "Five years",
	YTD:  "Year-to-data",
	Next: "Next upcoming",
}

// PathRanges maps the string keys from the JSON to the PathRange
// constant values.
var PathRanges = map[string]PathRange{
	"next": Next,
	"1m":   Mo1,
	"3m":   Mo3,
	"6m":   Mo6,
	"5y":   Yr5,
	"2y":   Yr2,
	"1y":   Yr1,
	"ytd":  YTD,
}

// PathRangeJSON maps a PathRange to the string used in the JSON.
var PathRangeJSON = map[PathRange]string{
	Mo1:  "1m",
	Mo3:  "3m",
	Mo6:  "6m",
	Yr1:  "1y",
	Yr2:  "2y",
	Yr5:  "5y",
	YTD:  "ytd",
	Next: "next",
}

// MarshalJSON implements the Marshaler interface for AnnounceTime.
func (p *PathRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(PathRangeJSON[*p])
}

// UnmarshalJSON implements the Unmarshaler interface for AnnounceTime.
func (p *PathRange) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling path range, got %s", data)
	}
	return p.Set(s)
}

// Set sets the issue type using a string.
func (p *PathRange) Set(s string) error {
	// Ensure the provided string matches on the keys in the map.
	got, ok := PathRanges[s]
	if !ok {
		return fmt.Errorf("invalid issue type %q", s)
	}
	// Set the issue type to the value found in the map per the key.
	*p = got
	return nil
}

// String implements the Stringer interface for PathRange.
func (p PathRange) String() string {
	return pathRangeDescription[p]
}
