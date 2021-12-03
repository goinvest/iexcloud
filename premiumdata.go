// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// Recommendation models the buy, hold, sell recommendations for a stock.
type Recommendation struct {
	BuyRatings                  int       `json:"ratingBuy"`
	OverweightRatings           int       `json:"ratingOverweight"`
	HoldRatings                 int       `json:"ratingHold"`
	UnderweightRatings          int       `json:"ratingUnderweight"`
	SellRatings                 int       `json:"ratingSell"`
	NoRatings                   int       `json:"ratingNone"`
	ConsensusRating             float64   `json:"ratingScaleMark"`
	ConsensusStartDate          EpochTime `json:"consensusStartDate"`
	CorporateActionsAppliedDate EpochTime `json:"corporateActionsAppliedDate"`
	ConsensusEndDate            EpochTime `json:"consensusEndDate"`
	ConsensusRatingOneToFive    float64   `json:"ratingScaleMarkOneToFive"`
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

// Estimates models the latest consensus esimtate for the next fiscal period.
type Estimates struct {
	Symbol    string     `json:"symbol"`
	Estimates []Estimate `json:"estimates"`
}

// Estimate models one estimate.
type Estimate struct {
	AnnounceTime                string  `json:"announceTime"`
	ConsensusEPS                float64 `json:"consensusEPS"`
	BookValuePerShare           float64 `json:"consensusBPS"`
	CashFlowValuePerShare       float64 `json:"consensusCPS"`
	CapitalExpenditures         float64 `json:"consensusCPX"`
	DividendPerShare            float64 `json:"consensusDPS"`
	EBIT                        float64 `json:"consensusEBI"`
	EBITDA                      float64 `json:"consensusEBT"`
	FundsFromOperations         float64 `json:"consensusFFO"`
	EPSFullyReported            float64 `json:"consensusGPS"`
	GrossMargin                 float64 `json:"consensusGRM"`
	NetAssetValue               float64 `json:"consensusNAV"`
	NetIncome                   float64 `json:"consensusNET"`
	OperatingProfit             float64 `json:"consensusOPR"`
	PreTaxProfit                float64 `json:"consensusPRE"`
	ReturnOnAssets              float64 `json:"consensusROA"`
	ReturnOnEquity              float64 `json:"consensusROE"`
	Revenue                     float64 `json:"consensusSAL"`
	Currency                    string  `json:"currency"`
	FiscalEndDate               Date    `json:"fiscalEndDate"`
	FiscalPeriod                string  `json:"fiscalPeriod"`
	NumberOfEstimates           int     `json:"numberOfEstimates"`
	NumEstEPS                   int     `json:"numberOfEstimatesEPS"`
	NumEstBookValuePerShare     int     `json:"numberOfEstimatesBPS"`
	NumEstCashFlowValuePerShare int     `json:"numberOfEstimatesCPS"`
	NumEstCapitalExpenditures   int     `json:"numberOfEstimatesCPX"`
	NumEstDividendPerShare      int     `json:"numberOfEstimatesDPS"`
	NumEstEBIT                  int     `json:"numberOfEstimatesEBI"`
	NumEstEBITDA                int     `json:"numberOfEstimatesEBT"`
	NumEstFundsFronOperations   int     `json:"numberOfEstimatesFFO"`
	NumEstEPSFullyReported      int     `json:"numberOfEstimatesGPS"`
	NumEstGrossMargin           int     `json:"numberOfEstimatesGRM"`
	NumEstNetAssetValue         int     `json:"numberOfEstimatesNAV"`
	NumEstNetIncome             int     `json:"numberOfEstimatesNET"`
	NumEstOperatingProfit       int     `json:"numberOfEstimatesOPR"`
	NumEstPreTaxProfit          int     `json:"numberOfEstimatesPRE"`
	NumEstReturnOnAssets        int     `json:"numberOfEstimatesROA"`
	NumEstReturnOnEquity        int     `json:"numberOfEstimatesROE"`
	NumEstRevenue               int     `json:"numberOfEstimatesSAL"`
	PeriodType                  string  `json:"periodType"`
	ReportDate                  Date    `json:"reportDate"`
	Symbol                      string  `json:"symbol"`
}

// PriceTarget models the latest average, high, and low analyst price target for
// a symbol.
type PriceTarget struct {
	Symbol      string  `json:"symbol"`
	UpdatedDate Date    `json:"updatedDate"`
	Average     float64 `json:"priceTargetAverage"`
	High        float64 `json:"priceTargetHigh"`
	Low         float64 `json:"priceTargetLow"`
	NumAnalysts int     `json:"numberOfAnalysts"`
	Currency    string  `json:"currency"`
}
