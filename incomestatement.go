// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// IncomeStatements pulls income statement data. Available quarterly (4 quarters) and
// annually (4 years).
type IncomeStatements struct {
	Symbol     string            `json:"symbol"`
	Statements []IncomeStatement `json:"income"`
}

// IncomeStatement models one income statement.
type IncomeStatement struct {
	ReportDate             Date    `json:"reportDate"`
	TotalRevenue           float64 `json:"totalRevenue"`
	CostOfRevenue          float64 `json:"costOfRevenue"`
	GrossProfit            float64 `json:"grossProfit"`
	ResearchAndDevelopment float64 `json:"researchAndDevelopment"`
	SellingGeneralAndAdmin float64 `json:"sellingGeneralAndAdmin"`
	OperatingExpense       float64 `json:"operatingExpense"`
	OperatingIncome        float64 `json:"operatingIncome"`
	OtherIncomeExpenseNet  float64 `json:"otherIncomeExpenseNet"`
	EBIT                   float64 `json:"ebit"`
	InterestIncome         float64 `json:"interestIncome"`
	PretaxIncome           float64 `json:"pretaxIncome"`
	IncomeTax              float64 `json:"incomeTax"`
	MinorityInterest       float64 `json:"minorityInterest"`
	NetIncome              float64 `json:"netIncome"`
	NetIncomeBasic         float64 `json:"netIncomeBasic"`
}

// AnnualIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualIncomeStatements(stock string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d?period=annual",
		url.PathEscape(stock), num)
	return c.incomeStatements(endpoint)
}

// QuarterlyIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyIncomeStatements(stock string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.incomeStatements(endpoint)
}

func (c Client) incomeStatements(endpoint string) (IncomeStatements, error) {
	is := &IncomeStatements{}
	err := c.GetJSON(endpoint, is)
	return *is, err
}
