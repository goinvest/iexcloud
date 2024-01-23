// Copyright (c) 2019-2024 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

type TimeSeriesRange string

const (
	Today       TimeSeriesRange = "today"
	Yesterday   TimeSeriesRange = "yesterday"
	YearToDate  TimeSeriesRange = "ytd"
	LastWeek    TimeSeriesRange = "last-week"
	LastMonth   TimeSeriesRange = "last-month"
	LastQuarter TimeSeriesRange = "last-quarter"
)

type TimeSeriesQueryParameters struct {
}
