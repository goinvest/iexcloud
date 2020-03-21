// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Status models the IEX Cloud API system status
type Status struct {
	Status  string    `json:"status"`
	Version string    `json:"version"`
	Time    EpochTime `json:"time"`
}
