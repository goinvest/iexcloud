// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// BalanceSheets pulls balance sheet data. Available quarterly (4 quarters) and
// annually (4 years).
type BalanceSheets struct {
	Symbol     string         `json:"symbol"`
	Statements []BalanceSheet `json:"balancesheet"`
}

// BalanceSheet models one balance sheet statement.
type BalanceSheet struct {
	ReportDate              Date `json:"reportDate"`
	CurrentCash             int  `json:"currentCash"`
	ShortTermInvestments    int  `json:"shortTermInvestments"`
	Receivables             int  `json:"receivables"`
	Inventory               int  `json:"invetnory"`
	OtherCurrentAssets      int  `json:"otherCurrentAssets"`
	CurrentAssets           int  `json:"currentAssets"`
	LongTermInvestments     int  `json:"longTermInvestments"`
	PropertyPlanetEquipment int  `json:"propertyPlantEquipment"`
	Goodwill                int  `json:"goodwill"`
	IntangibleAssets        int  `json:"intangibleAssets"`
	OtherAssets             int  `json:"otherAssets"`
	TotalAssets             int  `json:"totalAssets"`
	AccountsPayable         int  `json:"accountsPayable"`
	CurrentLongTermDebt     int  `json:"currentLongTermDebt"`
	OtherCurrentLiabilities int  `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities int  `json:"totalCurrentLiabilities"`
	LongTermDebt            int  `json:"longTermDebt"`
	OtherLiablities         int  `json:"otherLiabilities"`
	MinorityInterest        int  `json:"minorityInterest"`
	TotalLiabilities        int  `json:"totalLiabilities"`
	CommonStock             int  `json:"commonStock"`
	RetainedEarnings        int  `json:"retainedEarnings"`
	TreasuryStock           int  `json:"treasuryStock"`
	CapitalSurplus          int  `json:"capitalSurplus"`
	ShareholderEquity       int  `json:"shareholderEquity"`
	NetTangibleAssets       int  `json:"netTangibleAssets"`
}

// AnnualBalanceSheets returns the annual balance sheets from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) AnnualBalanceSheets(stock string) (BalanceSheets, error) {
	endpoint := "/stock/" + stock + "/balance-sheet?period=annual"
	return c.balanceSheets(endpoint)
}

// QuarterlyBalanceSheets returns the quarterly balance sheets from the IEX
// Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyBalanceSheets(stock string) (BalanceSheets, error) {
	endpoint := "/stock/" + stock + "/balance-sheet"
	return c.balanceSheets(endpoint)
}

func (c Client) balanceSheets(endpoint string) (BalanceSheets, error) {
	bs := &BalanceSheets{}
	err := c.GetJSON(endpoint, bs)
	return *bs, err
}
