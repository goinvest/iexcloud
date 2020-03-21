// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import "time"

// DataPoint models a single data point.
type DataPoint struct {
	Key         string    `json:"key"`
	Weight      int       `json:"weight"`
	Description string    `json:"description"`
	LastUpdated time.Time `json:"lastUpdated"`
}
