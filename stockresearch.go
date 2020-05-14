// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// AdvancedStats provides everything in key stats plus additional advanced
// stats such as EBITDA, ratios, key financial data, and more.
type AdvancedStats struct {
	KeyStats
	Beta                     float64 `json:"beta"`
	TotalCash                float64 `json:"totalCash"`
	CurrentDebt              float64 `json:"currentDebt"`
	Revenue                  float64 `json:"revenue"`
	GrossProfit              float64 `json:"grossProfit"`
	TotalRevenue             float64 `json:"totalRevenue"`
	EBITDA                   float64 `json:"EBITDA"`
	RevenuePerShare          float64 `json:"revenuePerShare"`
	RevenuePerEmployee       float64 `json:"revenuePerEmployee"`
	DebtToEquity             float64 `json:"debtToEquity"`
	ProfitMargin             float64 `json:"profitMargin"`
	EnterpriseValue          float64 `json:"enterpriseValue"`
	EnterpriseValueToRevenue float64 `json:"enterpriseValueToRevenue"`
	PriceToSales             float64 `json:"priceToSales"`
	PriceToBook              float64 `json:"priceToBook"`
	ForwardPERatio           float64 `json:"forwardPERatio"`
	PEGRatio                 float64 `json:"pegRatio"`
	PEHigh                   float64 `json:"peHigh"`
	PELow                    float64 `json:"peLow"`
	Week52HighDate           float64 `json:"week52highDate"`
	Week52LowDate            float64 `json:"week52lowDate"`
	PutCallRatio             float64 `json:"putCallRatio"`
}

// Recommendation models the buy, hold, sell recommendations for a stock.
type Recommendation struct {
	ConsensusEndDate            EpochTime `json:"consensusEndDate"`
	ConsensusStartDate          EpochTime `json:"consensusStartDate"`
	CorporateActionsAppliedDate EpochTime `json:"corporateActionsAppliedDate"`
	BuyRatings                  int       `json:"ratingBuy"`
	OverweightRatings           int       `json:"ratingOverweight"`
	HoldRatings                 int       `json:"ratingHold"`
	UnderweightRatings          int       `json:"ratingUnderweight"`
	SellRatings                 int       `json:"ratingSell"`
	NoRatings                   int       `json:"ratingNone"`
	ConsensusRating             float64   `json:"ratingScaleMark"`
}

// Estimates models the latest consensus esimtate for the next fiscal period.
type Estimates struct {
	Symbol    string     `json:"symbol"`
	Estimates []Estimate `json:"estimates"`
}

// Estimate models one estimate.
type Estimate struct {
	ConsensusEPS      float64 `json:"consensusEPS"`
	NumberOfEstimates int     `json:"numberOfEstimates"`
	FiscalPeriod      string  `json:"fiscalPeriod"`
	FiscalEndDate     Date    `json:"fiscalEndDate"`
	ReportDate        Date    `json:"reportDate"`
	AnnounceTime      string  `json:"announceTime"`
	Currency          string  `json:"currency"`
}

// FundOwner models a fund owning a stock.
type FundOwner struct {
	AdjustedHolding     float64   `json:"adjHolding"`
	AdjustedMarketValue float64   `json:"adjMv"`
	Name                string    `json:"entityProperName"`
	ReportDate          EpochTime `json:"reportDate"`
	ReportedHolding     float64   `json:"reportedHolding"`
	ReportedMarketValue float64   `json:"reportedMv"`
}

// InstitutionalOwner models an institutional owner of a stock.
type InstitutionalOwner struct {
	AdjustedHolding     float64   `json:"adjHolding"`
	AdjustedMarketValue float64   `json:"adjMv"`
	EntityName          string    `json:"entityProperName"`
	ReportDate          EpochTime `json:"reportDate"`
	ReportedHolding     float64   `json:"reportedHolding"`
}

// KeyStats models the data returned from IEX Cloud's /stats endpoint.
type KeyStats struct {
	Name                string  `json:"companyName"`
	MarketCap           float64 `json:"marketCap"`
	Week52High          float64 `json:"week52High"`
	Week52Low           float64 `json:"week52Low"`
	Week52Change        float64 `json:"week52Change"`
	SharesOutstanding   float64 `json:"sharesOutstanding"`
	Float               float64 `json:"float"`
	Avg10Volume         float64 `json:"avg10Volume"`
	Avg30Volume         float64 `json:"avg30Volume"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
	Employees           int     `json:"employees"`
	TTMEPS              float64 `json:"ttmEPS"`
	TTMDividendRate     float64 `json:"ttmDividendRate"`
	DividendYield       float64 `json:"dividendYield"`
	NextDividendDate    Date    `json:"nextDividendDate"`
	ExDividendDate      Date    `json:"exDividendDate"`
	NextEarningsDate    Date    `json:"nextEarningsDate"`
	PERatio             float64 `json:"peRatio"`
	Beta                float64 `json:"beta"`
	MaxChangePercent    float64 `json:"maxChangePercent"`
	Year5ChangePercent  float64 `json:"year5ChangePercent"`
	Year2ChangePercent  float64 `json:"year2ChangePercent"`
	Year1ChangePercent  float64 `json:"year1ChangePercent"`
	YTDChangePercent    float64 `json:"ytdChangePercent"`
	Month6ChangePercent float64 `json:"month6ChangePercent"`
	Month3ChangePercent float64 `json:"month3ChangePercent"`
	Month1ChangePercent float64 `json:"month1ChangePercent"`
	Day30ChangePercent  float64 `json:"day30ChangePercent"`
	Day5ChangePercent   float64 `json:"day5ChangePercent"`
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
}
