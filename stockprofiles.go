// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Company models the company data from the /company endpoint.
type Company struct {
	Symbol         string   `json:"symbol"`
	Name           string   `json:"companyName"`
	Exchange       string   `json:"exchange"`
	Industry       string   `json:"industry"`
	Website        string   `json:"website"`
	Description    string   `json:"description"`
	CEO            string   `json:"CEO"`
	IssueType      string   `json:"issueType"`
	Sector         string   `json:"sector"`
	Employees      int      `json:"employees"`
	Tags           []string `json:"tags"`
	SecurityName   string   `json:"securityName"`
	PrimarySICCode int      `json:"primarySicCode"`
	Address        string   `json:"address"`
	Address2       string   `json:"address2"`
	State          string   `json:"state"`
	City           string   `json:"city"`
	Zip            string   `json:"zip"`
	Country        string   `json:"country"`
	Phone          string   `json:"phone"`
}

// InsiderRoster models the top 10 insiders with the most recent information.
type InsiderRoster struct {
	EntityName string `json:"entityName"`
	Position   int    `json:"position"`
	ReportDate Date   `json:"reportDate"`
}

// InsiderSummary models a summary of insider information.
type InsiderSummary struct {
	Name           string `json:"fullName"`
	NetTransaction int    `json:"netTransaction"`
	ReportedTitle  string `json:"reportedTitle"`
	TotalBought    int    `json:"totalBought"`
	TotalSold      int    `json:"totalSold"`
}

// InsiderTransaction models a buy or sell transaction made by an insider of a
// company.
type InsiderTransaction struct {
	EffectiveDate EpochTime `json:"effectiveDate"`
	Name          string    `json:"fullName"`
	ReportedTitle string    `json:"reportedTitle"`
	Price         float64   `json:"tranPrice"`
	Shares        int       `json:"tranShares"`
	Value         float64   `json:"tranValue"`
}

// Logo models the /logo endpoint.
type Logo struct {
	URL string `json:"url"`
}

// RelevantStocks models a list of relevant stocks that may or may not be
// peers.
type RelevantStocks struct {
	Peers   bool     `json:"peers"`
	Symbols []string `json:"symbols"`
}
