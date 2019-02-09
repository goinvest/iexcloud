// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// IncomeStatements pulls income statement data. Available quarterly (4 quarters) and
// annually (4 years).
type IncomeStatements struct {
	Symbol     string            `json:"symbol"`
	Statements []IncomeStatement `json:"income"`
}

// IncomeStatement models one income statement.
type IncomeStatement struct {
	ReportDate             Date `json:"reportDate"`
	TotalRevenue           int  `json:"totalRevenue"`
	CostOfRevenue          int  `json:"costOfRevenue"`
	GrossProfit            int  `json:"grossProfit"`
	ResearchAndDevelopment int  `json:"researchAndDevelopment"`
	SellingGeneralAndAdmin int  `json:"sellingGeneralAndAdmin"`
	OperatingExpense       int  `json:"operatingExpense"`
	OperatingIncome        int  `json:"operatingIncome"`
	OtherIncomeExpenseNet  int  `json:"otherIncomeExpenseNet"`
	EBIT                   int  `json:"ebit"`
	InterestIncome         int  `json:"interestIncome"`
	PretaxIncome           int  `json:"pretaxIncome"`
	IncomeTax              int  `json:"incomeTax"`
	MinorityInterest       int  `json:"minorityInterest"`
	NetIncome              int  `json:"netIncome"`
	NetIncomeBasic         int  `json:"netIncomeBasic"`
}
