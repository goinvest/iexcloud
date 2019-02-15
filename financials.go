// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// Financials models income statement, balance sheet, and cash flow data from
// the most recent reported quarter.
type Financials struct {
	Symbol     string      `json:"symbol"`
	Financials []Financial `json:"financials"`
}

// Financial pulls income statement, balance sheet, and cash flow data from
// the most recent reported quarter.
type Financial struct {
	ReportDate             Date    `json:"reportDate"`
	GrossProfit            float64 `json:"grossProfit"`
	CostOfRevenue          float64 `json:"costOfRevenue"`
	OperatingRevenue       float64 `json:"operatingRevenue"`
	TotalRevenue           float64 `json:"totalRevenue"`
	OperatingIncome        float64 `json:"operatingIncome"`
	NetIncome              float64 `json:"netIncome"`
	ResearchAndDevelopment float64 `json:"researchAndDevelopment"`
	OperatingExpense       float64 `json:"operatingExpense"`
	CurrentAssets          float64 `json:"currentAssets"`
	TotalAssets            float64 `json:"totalAssets"`
	TotalLiabilities       float64 `json:"totalLiabilities"`
	CurrentCash            float64 `json:"currentCash"`
	TotalCash              float64 `json:"totalCash"`
	TotalDebt              float64 `json:"totalDebt"`
	ShareholderEquity      float64 `json:"shareholderEquity"`
	CashChange             float64 `json:"cashChange"`
	CashFlow               float64 `json:"cashFlow"`
	OperatingGainsLosses   string  `json:"operatingGainsLosses"`
}

// AnnualFinancials returns the specified number of most recent annual
// financials from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualFinancials(stock string, num int) (Financials, error) {
	endpoint := fmt.Sprintf("/stock/%s/financials/%d?period=annual",
		url.PathEscape(stock), num)
	return c.financials(endpoint)
}

// QuarterlyFinancials returns the specified number of most recent quarterly
// financials from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyFinancials(stock string, num int) (Financials, error) {
	endpoint := fmt.Sprintf("/stock/%s/financials/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.financials(endpoint)
}

func (c Client) financials(endpoint string) (Financials, error) {
	financials := Financials{}
	err := c.GetJSON(endpoint, &financials)
	return financials, err
}
