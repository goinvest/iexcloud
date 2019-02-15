// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// USExchange provides information about one U.S. exchange including the name,
// the Market identifier code, the ID used to identify the exchange on the
// consolidated tape, the FINRA OATS exchange participant ID, and the type of
// securities traded by the exchange.
type USExchange struct {
	Name     string `json:"name"`
	MarketID int    `json:"mic"`
	TapeID   string `json:"tapeId"`
	OATSID   string `json:"oatsId"`
	Type     string `json:"type"`
}

// USExchanges returns an array of U.S. Exchanges.
func (c Client) USExchanges() ([]USExchange, error) {
	e := []USExchange{}
	endpoint := "/ref-data/market/us/exchanges"
	err := c.GetJSON(endpoint, &e)
	return e, err
}
