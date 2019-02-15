// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Date models a report date
type Date time.Time

// UnmarshalJSON implements the Unmarshaler interface for Date.
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

// MarshalJSON implements the Marshaler interface for Date.
func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	return json.Marshal(t.Format("2006-01-02"))
}

// PathRange refers to the date range used in the path of an endpoint.
type PathRange int

// Enum values for PathRange.
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

// MarshalJSON implements the Marshaler interface for PathRange.
func (p *PathRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(PathRangeJSON[*p])
}

// UnmarshalJSON implements the Unmarshaler interface for PathRange.
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

// EpochTime refers to unix timestamps used for some fields in the API
type EpochTime time.Time

// MarshalJSON implements the Marshaler interface for EpochTime.
func (e *EpochTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(time.Time(*e).Unix())), nil
}

// UnmarshalJSON implements the Unmarshaler interface for EpochTime.
func (e *EpochTime) UnmarshalJSON(data []byte) (err error) {
	ts, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	// Per docs: If the value is -1, IEX has not quoted the symbol in the trading day.
	if ts == -1 {
		return
	}

	*e = EpochTime(time.Unix(int64(ts)/1000, 0))
	return nil
}

// String implements the Stringer interface for EpochTime.
func (e EpochTime) String() string { return time.Time(e).String() }
