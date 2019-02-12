// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Logo models the /logo endpoint.
type Logo struct {
	URL string `json:"url"`
}

// Logo returns the logo data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Logo(stock string) (Logo, error) {
	logo := &Logo{}
	endpoint := fmt.Sprintf("/stock/%s/logo", url.PathEscape(stock))
	err := c.GetJSON(endpoint, logo)
	return *logo, err
}
