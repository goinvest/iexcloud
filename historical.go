// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"time"
)

// HistoricalTimeFrame enum for selecting time frame of historical data
type HistoricalTimeFrame string

const (
	// FiveDayHistorical Five days historically adjusted market-wide data
	FiveDayHistorical HistoricalTimeFrame = "5d"
	// FiveDay10MinuteHistorical Five days historically adjusted market-wide data in 10 minute intervals
	FiveDay10MinuteHistorical HistoricalTimeFrame = "5dm"
	// OneMonthHistorical One month (default) historically adjusted market-wide data
	OneMonthHistorical HistoricalTimeFrame = "1m"
	// OneMonth30MinuteHistorical One month historically adjusted market-wide data in 30 minute intervals
	OneMonth30MinuteHistorical HistoricalTimeFrame = "1mm"
	// ThreeMonthHistorical Three months historically adjusted market-wide data
	ThreeMonthHistorical HistoricalTimeFrame = "3m"
	// SixMonthHistorical Six months historically adjusted market-wide data
	SixMonthHistorical HistoricalTimeFrame = "6m"
	// OneYearHistorical One year historically adjusted market-wide data
	OneYearHistorical HistoricalTimeFrame = "1y"
	// TwoYearHistorical Two year historically adjusted market-wide data
	TwoYearHistorical HistoricalTimeFrame = "2y"
	// FiveYearHistorical Five year historically adjusted market-wide data
	FiveYearHistorical HistoricalTimeFrame = "5y"
	// YearToDateHistorical Year to date historically adjusted market-wide data
	YearToDateHistorical HistoricalTimeFrame = "ytd"
	// MaxHistorical All available historically adjusted market-wide data up to 15 years
	MaxHistorical HistoricalTimeFrame = "max"
)

// Valid Determines if HistoricalTimeFrame is a defined constant
func (htf HistoricalTimeFrame) Valid() bool {
	switch htf {
	case
		FiveDayHistorical,
		FiveDay10MinuteHistorical,
		OneMonthHistorical,
		OneMonth30MinuteHistorical,
		ThreeMonthHistorical,
		SixMonthHistorical,
		OneYearHistorical,
		TwoYearHistorical,
		FiveYearHistorical,
		YearToDateHistorical,
		MaxHistorical:
		return true
	default:
		return false
	}
}

// IntradayHistoricalDataPoint Represents a single intraday data point for a stock
type IntradayHistoricalDataPoint struct {
	Date                 Date    `json:"date"`
	Minute               string  `json:"minute"`
	Label                string  `json:"label"`
	High                 float64 `json:"high"`
	Low                  float64 `json:"low"`
	Average              float64 `json:"average"`
	Volume               int     `json:"volume"`
	Notional             float64 `json:"notional"`
	NumberOfTrades       int     `json:"numberOfTrades"`
	MarketHigh           float64 `json:"marketHigh"`
	MarketLow            float64 `json:"marketLow"`
	MarketAverage        float64 `json:"marketAverage"`
	MarketVolume         int     `json:"marketVolume"`
	MarketNotional       float64 `json:"marketNotional"`
	MarketNumberOfTrades int     `json:"marketNumberOfTrades"`
	Open                 float64 `json:"open"`
	Close                float64 `json:"close"`
	MarketOpen           float64 `json:"marketOpen"`
	MarketClose          float64 `json:"marketClose"`
	ChangeOverTime       float64 `json:"changeOverTime"`
	MarketChangeOverTime float64 `json:"marketChangeOverTime"`
}

// HistoricalOptions optional query params to pass to historical endpoint
// If values are false or 0 they aren't passed.
type HistoricalOptions struct {
	ChartCloseOnly  bool   `url:"chartCloseOnly,omitempty"`
	ChartSimplify   bool   `url:"chartSimplify,omitempty"`
	ChartInterval   int    `url:"chartInterval,omitempty"`
	ChangeFromClose bool   `url:"changeFromClose,omitempty"`
	ChartLast       int    `url:"chartLast,omitempty"`
	DisplayPercent  bool   `url:"displayPercent,omitempty"`
	Range           string `url:"range,omitempty"`
	ExactDate       string `url:"exactDate,omitempty"`
	Sort            string `url:"sort,omitempty"`
	IncludeToday    bool   `url:"includeToday,omitempty"`
}

// IntradayHistoricalOptions optional query params to pass to intraday historical endpoint
// If values are false or 0 they aren't passed.
type IntradayHistoricalOptions struct {
	ChartIEXOnly    bool `url:"chartIEXOnly,omitempty"`
	ChartReset      bool `url:"chartReset,omitempty"`
	ChartSimplify   bool `url:"chartSimplify,omitempty"`
	ChartInterval   int  `url:"chartInterval,omitempty"`
	ChangeFromClose bool `url:"changeFromClose,omitempty"`
	ChartLast       int  `url:"chartLast,omitempty"`
}

// HistoricalDataPoint Represents a single historical data point for a stock
type HistoricalDataPoint struct {
	Close          float64    `json:"close"`
	High           float64    `json:"high"`
	Low            float64    `json:"low"`
	Open           float64    `json:"open"`
	Symbol         string     `json:"symbol"`
	Volume         float64    `json:"volume"`
	ID             string     `json:"id"`
	Key            string     `json:"key"`
	Subkey         string     `json:"subkey"`
	Date           Date       `json:"date"`
	Minute         HourMinute `json:"minute"`
	UOpen          float64    `json:"uOpen"`
	UClose         float64    `json:"uClose"`
	UHigh          float64    `json:"uHigh"`
	ULow           float64    `json:"uLow"`
	UVolume        int        `json:"uVolume"`
	Change         float64    `json:"change"`
	ChangePercent  float64    `json:"changePercent"`
	Label          string     `json:"label"`
	ChangeOverTime float64    `json:"changeOverTime"`
}

// IntradayOptions optional query params to pass to intraday endpoint
// If values are false or 0 they aren't passed.
type IntradayOptions struct {
	ChartIEXOnly     bool   `url:"chartIEXOnly,omitempty"`
	ChartReset       bool   `url:"chartReset,omitempty"`
	ChartSimplify    bool   `url:"chartSimplify,omitempty"`
	ChartInterval    int    `url:"chartInterval,omitempty"`
	ChangeFromClose  bool   `url:"changeFromClose,omitempty"`
	ChartLast        int    `url:"chartLast,omitempty"`
	ExactDate        string `url:"exactDate,omitempty"` // Formatted as YYYYMMDD
	ChartIEXWhenNull bool   `url:"chartIEXWhenNull,omitempty"`
}

// SetExactDate formats a given date as IEX expects
func (opt *IntradayOptions) SetExactDate(day time.Time) {
	opt.ExactDate = day.Format("20060102")
}
