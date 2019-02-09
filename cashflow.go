// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// CashFlows pulls cash flow data. Available quarterly (4 quarters) or annually
// (4 years).
type CashFlows struct {
	Symbol             string              `json:"symbol"`
	CashFlowStatements []CashFlowStatement `json:"cashflow"`
}

// CashFlowStatement models one cash flow statement.
type CashFlowStatement struct {
	ReportDate              ReportDate `json:"reportDate"`
	NetIncome               int        `json:"netIncome"`
	Depreciation            int        `json:"depreciation"`
	ChangesInReceivables    int        `json:"changesInReceivables"`
	ChangesInInventories    int        `json:"changesInInventories"`
	CashChange              int        `json:"cashChange"`
	CashFlow                int        `json:"cashFlow"`
	CapitalExpenditures     int        `json:"capitalExpenditures"`
	Investment              int        `json:"investments"`
	InvestingActivityOther  int        `json:"investingActivityOther"`
	TotalInvestingCashFloes int        `json:"totalInvestingCashFlows"`
	DividensPaid            int        `json:"dividendsPaid"`
	NetBorrowings           int        `json:"netBorrowings"`
	OtherFinancingCashFlows int        `json:"otherFinancingCashFlows"`
	CashFlowFinancing       int        `json:"cashFlowFinancing"`
	ExchangeRateEffect      int        `json:"exchangeRateEffect"`
}
