// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
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

// BalanceSheet models one balance sheet statement. Normally the amounts
// returned are integers, although the currentCash for UBNT returned is a
// float; therefore, these are all floats.
type BalanceSheet struct {
	ReportDate              Date    `json:"reportDate"`
	FilingType              string  `json:"filingType"`
	FiscalDate              Date    `json:"fiscalDate"`
	FiscalQuarter           int     `json:"fiscalQuarter"`
	FiscalYear              int     `json:"fiscalYear"`
	Currency                string  `json:"currency"`
	CurrentCash             float64 `json:"currentCash"`
	ShortTermInvestments    float64 `json:"shortTermInvestments"`
	Receivables             float64 `json:"receivables"`
	Inventory               float64 `json:"inventory"`
	OtherCurrentAssets      float64 `json:"otherCurrentAssets"`
	CurrentAssets           float64 `json:"currentAssets"`
	LongTermInvestments     float64 `json:"longTermInvestments"`
	PropertyPlantEquipment  float64 `json:"propertyPlantEquipment"`
	Goodwill                float64 `json:"goodwill"`
	IntangibleAssets        float64 `json:"intangibleAssets"`
	OtherAssets             float64 `json:"otherAssets"`
	TotalAssets             float64 `json:"totalAssets"`
	AccountsPayable         float64 `json:"accountsPayable"`
	CurrentLongTermDebt     float64 `json:"currentLongTermDebt"`
	OtherCurrentLiabilities float64 `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities float64 `json:"totalCurrentLiabilities"`
	LongTermDebt            float64 `json:"longTermDebt"`
	OtherLiablities         float64 `json:"otherLiabilities"`
	MinorityInterest        float64 `json:"minorityInterest"`
	TotalLiabilities        float64 `json:"totalLiabilities"`
	CommonStock             float64 `json:"commonStock"`
	RetainedEarnings        float64 `json:"retainedEarnings"`
	TreasuryStock           float64 `json:"treasuryStock"`
	CapitalSurplus          float64 `json:"capitalSurplus"`
	ShareholderEquity       float64 `json:"shareholderEquity"`
	NetTangibleAssets       float64 `json:"netTangibleAssets"`
}

// CashFlows pulls cash flow data. Available quarterly (4 quarters) or annually
// (4 years).
type CashFlows struct {
	Symbol     string     `json:"symbol"`
	Statements []CashFlow `json:"cashflow"`
}

// CashFlow models one cash flow statement.
type CashFlow struct {
	ReportDate              Date    `json:"reportDate"`
	FiscalDate              Date    `json:"fiscalDate"`
	Currency                string  `json:"currency"`
	NetIncome               float64 `json:"netIncome"`
	Depreciation            float64 `json:"depreciation"`
	ChangesInReceivables    float64 `json:"changesInReceivables"`
	ChangesInInventories    float64 `json:"changesInInventories"`
	CashChange              float64 `json:"cashChange"`
	CashFlow                float64 `json:"cashFlow"`
	CapitalExpenditures     float64 `json:"capitalExpenditures"`
	Investments             float64 `json:"investments"`
	InvestingActivityOther  float64 `json:"investingActivityOther"`
	TotalInvestingCashFlows float64 `json:"totalInvestingCashFlows"`
	DividendsPaid           float64 `json:"dividendsPaid"`
	NetBorrowings           float64 `json:"netBorrowings"`
	OtherFinancingCashFlows float64 `json:"otherFinancingCashFlows"`
	CashFlowFinancing       float64 `json:"cashFlowFinancing"`
	ExchangeRateEffect      float64 `json:"exchangeRateEffect"`
}

// Dividend models one dividend (basic) for the stock fundamentals.
type Dividend struct {
	Symbol       string  `json:"symbol"`
	ExDate       Date    `json:"exDate"`
	PaymentDate  Date    `json:"paymentDate"`
	RecordDate   Date    `json:"recordDate"`
	DeclaredDate Date    `json:"declaredDate"`
	Amount       float64 `json:"amount"`
	Flag         string  `json:"flag"`
	Currency     string  `json:"currency"`
	Description  string  `json:"description"`
	Frequency    string  `json:"frequency"`
}

// Earnings provides earnings data for a given company including the actual
// EPS, consensus, and fiscal period. Earnings are available quarterly (last 4
// quarters) and annually (last 4 years).
type Earnings struct {
	Symbol   string    `json:"symbol"`
	Earnings []Earning `json:"earnings"`
}

// Earning models the earnings for one date.
type Earning struct {
	EPSReportDate            Date         `json:"EPSReportDate"`
	EPSSurpriseDollar        float64      `json:"EPSSurpriseDollar"`
	EPSSurpriseDollarPercent float64      `json:"EPSSurpriseDollarPercent"`
	ActualEPS                float64      `json:"actualEPS"`
	AnnounceTime             AnnounceTime `json:"announceTime"`
	ConsensusEPS             float64      `json:"consensusEPS"`
	Currency                 string       `json:"currency"`
	FiscalEndDate            Date         `json:"fiscalEndDate"`
	FiscalPeriod             string       `json:"fiscalPeriod"`
	NumberOfEstimates        int          `json:"numberOfEstimates"`
	PeriodType               string       `json:"periodType"`
	Symbol                   string       `json:"symbol"`
	YearAgo                  float64      `json:"yearAgo"`
	YearAgoChangePercent     float64      `json:"yearAgoChangePercent"`
	ID                       string       `json:"id"`
	Key                      string       `json:"key"`
	SubKey                   string       `json:"subkey"`
	Date                     uint64       `json:"date"`
	Updated                  uint64       `json:"updated"`
}

// Financials models income statement, balance sheet, and cash flow data from
// the most recent reported quarter.
type Financials struct {
	Symbol     string      `json:"symbol"`
	Financials []Financial `json:"financials"`
}

// Financial pulls income statement, balance sheet, and cash flow data from the
// most recent reported quarter. This endpoint is carried over from the IEX 1.0
// API. Use the new cash-flow, income statement, and balance-sheet endpoints
// for new data.
type Financial struct {
	ReportDate             Date    `json:"reportDate"`
	FiscalDate             Date    `json:"fiscalDate"`
	Currency               string  `json:"currency"`
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
	CurrentDebt            float64 `json:"currentDebt"`
	ShortTermDebt          float64 `json:"shortTermDebt"`
	LongTermDebt           float64 `json:"LongTermDebt"`
	TotalCash              float64 `json:"totalCash"`
	TotalDebt              float64 `json:"totalDebt"`
	ShareholderEquity      float64 `json:"shareholderEquity"`
	CashChange             float64 `json:"cashChange"`
	CashFlow               float64 `json:"cashFlow"`
}

// FinancialsAsReported models multiple SEC financial 10-K or 10-Q filings.
type FinancialsAsReported []FinancialAsReported

// FinancialAsReported models an SEC financial filing.
type FinancialAsReported struct {
	ID            string    `json:"id"`
	Source        string    `json:"source"`
	Key           string    `json:"key"`
	Subkey        string    `json:"subkey"`
	Date          EpochTime `json:"date"`
	Updated       EpochTime `json:"updated"`
	FiscalYear    int64     `json:"formFiscalYear"`
	Version       string    `json:"version"`
	PeriodStart   EpochTime `json:"periodStart"`
	PeriodEnd     EpochTime `json:"periodEnd"`
	DateFiled     EpochTime `json:"dateFiled"`
	FiscalQuarter int64     `json:"formFiscalQuarter"`
}

// IncomeStatements pulls income statement data. Available quarterly (4 quarters) and
// annually (4 years).
type IncomeStatements struct {
	Symbol     string            `json:"symbol"`
	Statements []IncomeStatement `json:"income"`
}

// IncomeStatement models one income statement.
type IncomeStatement struct {
	ReportDate             Date    `json:"reportDate"`
	FiscalDate             Date    `json:"fiscalDate"`
	Currency               string  `json:"currency"`
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

// Split models the a stock split.
type Split struct {
	Symbol       string  `json:"symbol"`
	ExDate       Date    `json:"exDate"`
	DeclaredDate Date    `json:"declaredDate"`
	Ratio        float64 `json:"ratio"`
	ToFactor     float64 `json:"toFactor"`
	FromFactor   float64 `json:"fromFactor"`
	Description  string  `json:"description"`
}
