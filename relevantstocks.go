// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// RelevantStocks models a list of relevant stocks that may or may not be
// peers.
type RelevantStocks struct {
	Peers   bool     `json:"peers"`
	Symbols []string `json:"symbols"`
}

// RelevantStocks is similar to the peers endpoint, except this will return
// most active market symbols when peers are not available. If the symbols
// returned are not peers, the peers key will be false. This is not intended to
// represent a definitive or accurate list of peers, and is subject to change
// at any time.
func (c Client) RelevantStocks(stock string) (RelevantStocks, error) {
	rs := &RelevantStocks{}
	endpoint := fmt.Sprintf("/stock/%s/relevant", url.PathEscape(stock))
	err := c.GetJSON(endpoint, rs)
	return *rs, err
}
