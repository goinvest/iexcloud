// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// MostActive returns a list of quotes for the top 10 most active stocks from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) MostActive() ([]Quote, error) {
	return c.list("mostactive")
}

// Gainers returns a list of quotes for the top 10 stock gainers from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) Gainers() ([]Quote, error) {
	return c.list("gainers")
}

// Losers returns a list of quotes for the top 10 stock losers from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) Losers() ([]Quote, error) {
	return c.list("losers")
}

// IEXVolume returns a list of quotes for the top 10 IEX stocks by volume from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) IEXVolume() ([]Quote, error) {
	return c.list("iexvolume")
}

// IEXPercent returns a list of quotes for the top 10 IEX stocks by percent
// from the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) IEXPercent() ([]Quote, error) {
	return c.list("iexpercent")
}

// InFocus returns a list of quotes for the top 10 in focus stocks from the IEX
// Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) InFocus() ([]Quote, error) {
	return c.list("infocus")
}

func (c Client) list(list string) ([]Quote, error) {
	q := &[]Quote{}
	endpoint := "/stock/market/list/" + list
	err := c.GetJSON(endpoint, q)
	return *q, err
}
