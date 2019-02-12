// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"fmt"
	"net/url"
)

// CashFlows pulls cash flow data. Available quarterly (4 quarters) or annually
// (4 years).
type CashFlows struct {
	Symbol     string     `json:"symbol"`
	Statements []CashFlow `json:"cashflow"`
}

// CashFlow models one cash flow statement.
type CashFlow struct {
	ReportDate              Date    `json:"reportDate"`
	NetIncome               float64 `json:"netIncome"`
	Depreciation            float64 `json:"depreciation"`
	ChangesInReceivables    float64 `json:"changesInReceivables"`
	ChangesInInventories    float64 `json:"changesInInventories"`
	CashChange              float64 `json:"cashChange"`
	CashFlow                float64 `json:"cashFlow"`
	CapitalExpenditures     float64 `json:"capitalExpenditures"`
	Investment              float64 `json:"investments"`
	InvestingActivityOther  float64 `json:"investingActivityOther"`
	TotalInvestingCashFloes float64 `json:"totalInvestingCashFlows"`
	DividensPaid            float64 `json:"dividendsPaid"`
	NetBorrowings           float64 `json:"netBorrowings"`
	OtherFinancingCashFlows float64 `json:"otherFinancingCashFlows"`
	CashFlowFinancing       float64 `json:"cashFlowFinancing"`
	ExchangeRateEffect      float64 `json:"exchangeRateEffect"`
}

// AnnualCashFlows returns the specified number of most recent annual cash flow
// statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualCashFlows(stock string, num int) (CashFlows, error) {
	endpoint := fmt.Sprintf("/stock/%s/cash-flow/%d?period=annual",
		url.PathEscape(stock), num)
	return c.cashFlows(endpoint)
}

// QuarterlyCashFlows returns the specified number of most recent annual
// cash flow statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyCashFlows(stock string, num int) (CashFlows, error) {
	endpoint := fmt.Sprintf("/stock/%s/cash-flow/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.cashFlows(endpoint)
}

func (c Client) cashFlows(endpoint string) (CashFlows, error) {
	cf := &CashFlows{}
	err := c.GetJSON(endpoint, cf)
	return *cf, err
}
