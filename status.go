// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Status models the IEX Cloud API system status
type Status struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Time    int    `json:"time"`
}

// Status returns the IEX Cloud system status.
func (c Client) Status() (Status, error) {
	status := Status{}
	endpoint := "/status"
	err := c.GetJSONWithoutToken(endpoint, &status)
	return status, err
}
