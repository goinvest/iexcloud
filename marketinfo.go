// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// EarningsToday models the earning that will be reported today as two arrays:
// before the open and after market close. Each array contains an object with
// all keys from earnings, a quote object, and a headline key.
type EarningsToday struct {
	BeforeOpen    []TodayEarning `json:"bto"`
	AfterClose    []TodayEarning `json:"amc"`
	DuringTrading []TodayEarning `json:"other"`
}

// TodayEarning models a single earning being reported today containing all
// keys from earnings, a quote object, and a headline.
type TodayEarning struct {
	ConsensusEPS      float64      `json:"consensusEPS"`
	AnnounceTime      AnnounceTime `json:"announcetime"`
	NumberOfEstimates int          `json:"numberOfEstimates"`
	FiscalPeriod      string       `json:"fiscalPeriod"`
	FiscalEndDate     Date         `json:"fiscalEndDate"`
	Symbol            string       `json:"symbol"`
	Quote             Quote        `json:"quote"`
}

// Market models the traded volume on U.S. markets.
type Market struct {
	MIC         string    `json:"mic"`
	TapeID      string    `json:"tapeId"`
	Venue       string    `json:"venueName"`
	Volume      int       `json:"volume"`
	TapeA       int       `json:"tapeA"`
	TapeB       int       `json:"tapeB"`
	TapeC       int       `json:"tapeC"`
	Percent     float64   `json:"marketPercent"`
	LastUpdated EpochTime `json:"lastUpdated"`
}

// SectorPerformance models the performance based on each sector ETF.
type SectorPerformance struct {
	Type        string    `json:"sector"`
	Name        string    `json:"name"`
	Performance float64   `json:"performance"`
	LastUpdated EpochTime `json:"lastUpdated"`
}

// IPO is all available data for an IPO.
type IPO struct {
	Symbol                 string   `json:"symbol"`
	CompanyName            string   `json:"companyName"`
	ExpectedDate           Date     `json:"expectedDate"`
	LeadUnderwriters       []string `json:"leadUnderwriters"`
	Underwriters           []string `json:"underwriters"`
	CompanyCounsel         []string `json:"companyCounsel"`
	UnderwriterCounsel     []string `json:"underwriterCounsel"`
	Auditor                string   `json:"auditor"`
	Market                 string   `json:"market"`
	CIK                    string   `json:"cik"`
	Address                string   `json:"address"`
	City                   string   `json:"city"`
	State                  string   `json:"state"`
	Zip                    string   `json:"zip"`
	Phone                  string   `json:"phone"`
	CEO                    string   `json:"ceo"`
	Employees              int      `json:"employees"`
	URL                    string   `json:"url"`
	Status                 string   `json:"status"`
	SharesOffered          int      `json:"sharesOffered"`
	PriceLow               float64  `json:"priceLow"`
	PriceHigh              float64  `json:"priceHigh"`
	OfferAmount            int      `json:"offerAmount"`
	TotalExpenses          int      `json:"totalExpenses"`
	SharesOverAlloted      int      `json:"sharesOverAlloted"`
	ShareholderShares      int      `json:"shareholderShares"`
	SharesOutstanding      int      `json:"sharesOutstanding"`
	LockupPeriodExpiration string   `json:"lockupPeriodExpiration"`
	QuietPeriodExpiration  string   `json:"quietPeriodExpiration"`
	Revenue                int      `json:"revenue"`
	NetIncome              int      `json:"netIncome"`
	TotalAssets            int      `json:"totalAssets"`
	TotalLiabilities       int      `json:"totalLiabilities"`
	StockholderEquity      int      `json:"stockholderEquity"`
	CompanyDescription     string   `json:"companyDescription"`
	BusinessDescription    string   `json:"businessDescription"`
	UseOfProceeds          string   `json:"useOfProceeds"`
	Competition            string   `json:"competition"`
	Amount                 int      `json:"amount"`
	PercentOffered         string   `json:"percentOffered"`
}

// IPOView is IPO data structured for display to a user.
type IPOView struct {
	Company  string `json:"Company"`
	Symbol   string `json:"Symbol"`
	Price    string `json:"Price"`
	Shares   string `json:"Shares"`
	Amount   string `json:"Amount"`
	Float    string `json:"Float"`
	Percent  string `json:"Percent"`
	Market   string `json:"Market"`
	Expected Date   `json:"Expected"`
}

// IPOCalendar is a list of IPOs.
type IPOCalendar struct {
	RawData  []IPO     `json:"rawData"`
	ViewData []IPOView `json:"viewData"`
}

// UpcomingEarning is an upcoming earnings event.
type UpcomingEarning struct {
	Estimate
	Symbol   string `json:"symbol"`
	SymbolId string `json:"symbolId"`
}

// UpcomingEvents is all of the upcoming events.
type UpcomingEvents struct {
	IPOs      IPOCalendar       `json:"ipos"`
	Earnings  []UpcomingEarning `json:"earnings"`
	Dividends []Dividend        `json:"dividends"`
	Splits    []Split           `json:"splits"`
}
