// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Peers returns a slice of peer stock symbols from the IEX Cloud endpoint for
// the given stock symbol.
func (c Client) Peers(stock string) ([]string, error) {
	peers := &[]string{}
	endpoint := fmt.Sprintf("/stock/%s/peers", url.PathEscape(stock))
	err := c.GetJSON(endpoint, peers)
	return *peers, err
}
