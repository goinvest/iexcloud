// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// CoreEstimate modules a current or historical consensus analyst
// recommendation and price target.
type CoreEstimate struct {
	ID              string    `json:"CORE_ESTIMATES"`
	Key             string    `json:"key"`
	Subkey          string    `json:"subkey"`
	Symbol          string    `json:"symbol"`
	AnalystCount    int       `json:"analystCount"`
	ConsensusDate   Date      `json:"consensusDate"`
	MarketConsensus float64   `json:"marketConsensus"`
	TargetPrice     float64   `json:"marketConsensusTargetPrice"`
	Date            EpochTime `json:"date"`
	Updated         EpochTime `json:"updated"`
}

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
