// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Company models the company data from the /company endpoint.
type Company struct {
	Symbol      string    `json:"symbol"`
	Name        string    `json:"companyName"`
	Exchange    string    `json:"exchange"`
	Industry    string    `json:"industry"`
	Website     string    `json:"website"`
	Description string    `json:"description"`
	CEO         string    `json:"CEO"`
	IssueType   IssueType `json:"issueType"`
	Sector      string    `json:"sector"`
	Tags        []string  `json:"tags"`
}

// IssueType refers to the common issue type of the stock.
type IssueType int

const (
	blank IssueType = iota
	ad
	re
	ce
	si
	lp
	cs
	et
)

var issueTypeDescription = map[IssueType]string{
	ad:    "American Depository Receipt (ADR)",
	re:    "Real Estate Investment Trust (REIT)",
	ce:    "Closed end fund (Stock and Bond Fund)",
	si:    "Secondary Issue",
	lp:    "Limited Partnership",
	cs:    "Common Stock",
	et:    "Exchange Taded Fund (ETF)",
	blank: "Not available",
}

// IssueTypes maps the string keys from the JSON to the IssueType constant
// values.
var IssueTypes = map[string]IssueType{
	"ad":  ad,
	"re":  re,
	"ce":  ce,
	"cef": ce,
	"si":  si,
	"lp":  lp,
	"cs":  cs,
	"et":  et,
	"":    blank,
}

// IssueTypeJSON maps an IssueType to the string used in the JSON.
var IssueTypeJSON = map[IssueType]string{
	ad:    "ad",
	re:    "re",
	ce:    "ce",
	si:    "si",
	lp:    "lp",
	cs:    "cs",
	et:    "et",
	blank: "",
}

// String implements the Stringer interface for IssueType.
func (i IssueType) String() string {
	return issueTypeDescription[i]
}

// MarshalJSON implements the Marshaler interface for IssueType.
func (i *IssueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(IssueTypeJSON[*i])
}

// UnmarshalJSON implements the Unmarshaler interface for IssueType.
func (i *IssueType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("issueType should be a string, got %s", data)
	}
	return i.Set(s)
}

// Set sets the issue type using a string.
func (i *IssueType) Set(s string) error {
	// Ensure the provided string matches on the keys in the map.
	got, ok := IssueTypes[s]
	if !ok {
		return fmt.Errorf("invalid issue type %q", s)
	}
	// Set the issue type to the value found in the map per the key.
	*i = got
	return nil
}

// Company returns the copmany data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Company(stock string) (Company, error) {
	company := &Company{}
	endpoint := fmt.Sprintf("/stock/%s/company", url.PathEscape(stock))
	err := c.GetJSON(endpoint, company)
	return *company, err
}
