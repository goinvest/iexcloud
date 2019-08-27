// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// AdvancedStats provides everything in key stats plus additional advanced
// stats such as EBITDA, ratios, key financial data, and more.
type AdvancedStats struct {
	KeyStats
	TotalCash                float64 `json:"totalCash"`
	CurrentDebt              float64 `json:"currentDebt"`
	Revenue                  float64 `json:"revenue"`
	GrossProfit              float64 `json:"grossProfit"`
	TotalRevenue             float64 `json:"totalRevenue"`
	EBITDA                   float64 `json:"EBITDA"`
	RevenuePerShare          float64 `json:"revenuePerShare"`
	DebtToEquity             float64 `json:"debtToEquity"`
	ProfitMargin             float64 `json:"profitMargin"`
	EnterpriseValue          float64 `json:"enterpriseValue"`
	EnterpriseValueToRevenue float64 `json:"enterpriseValueToRevenue"`
	PriceToSales             float64 `json:"priceToSales"`
	PriceToBook              float64 `json:"priceToBook"`
	ForwardPERatio           float64 `json:"forwardPERatio"`
	PEGRatio                 float64 `json:"pegRatio"`
	Beta                     float64 `json:"beta"`
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
	CurrentCash             float64 `json:"currentCash"`
	ShortTermInvestments    float64 `json:"shortTermInvestments"`
	Receivables             float64 `json:"receivables"`
	Inventory               float64 `json:"inventory"`
	OtherCurrentAssets      float64 `json:"otherCurrentAssets"`
	CurrentAssets           float64 `json:"currentAssets"`
	LongTermInvestments     float64 `json:"longTermInvestments"`
	PropertyPlanetEquipment float64 `json:"propertyPlantEquipment"`
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

// Book models the data returned from the /book endpoint.
type Book struct {
	Quote       Quote       `json:"quote"`
	Bids        []BidAsk    `json:"bids"`
	Asks        []BidAsk    `json:"asks"`
	Trades      []Trade     `json:"trades"`
	SystemEvent SystemEvent `json:"systemEvent"`
}

// BidAsk models a bid or an ask for a quote.
type BidAsk struct {
	Price     float64   `json:"price"`
	Size      int       `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

// Trade models a trade for a quote.
type Trade struct {
	Price                 float64   `json:"price"`
	Size                  int       `json:"size"`
	TradeID               int       `json:"tradeId"`
	IsISO                 bool      `json:"isISO"`
	IsOddLot              bool      `json:"isOddLot"`
	IsOutsideRegularHours bool      `json:"isOutsideRegularHours"`
	IsSinglePriceCross    bool      `json:"isSinglePriceCross"`
	IsTradeThroughExempt  bool      `json:"isTradeThroughExempt"`
	Timestamp             EpochTime `json:"timestamp"`
}

// Auction models auction data for a security
type Auction struct {
	AuctionType          string    `json:"auctionType"`
	PairedShares         int       `json:"pairedShares"`
	ImbalanceShares      int       `json:"imbalanceShares"`
	ReferencePrice       float64   `json:"referencePrice"`
	IndicativePrice      float64   `json:"indicativePrice"`
	AuctionBookPrice     float64   `json:"auctionBookPrice"`
	CollarReferencePrice float64   `json:"collarReferencePrice"`
	LowerCollarPrice     float64   `json:"lowerCollarPrice"`
	UpperCollarPrice     float64   `json:"upperCollarPrice"`
	ExtensionNumber      int       `json:"extensionNumber"`
	StartTime            EpochTime `json:"startTime"`
	LastUpdate           EpochTime `json:"lastUpdate"`
}

// SystemEvent models a system event for a quote.
type SystemEvent struct {
	SystemEvent string    `json:"systemEvent"`
	Timestamp   EpochTime `json:"timestamp"`
}

// SecurityEvent models events which apply to a specific security
type SecurityEvent struct {
	SecurityEvent string    `json:"securityEvent"`
	Timestamp     EpochTime `json:"timestamp"`
}

// TradingStatus models the current trading status of a security
type TradingStatus struct {
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	Timestamp EpochTime `json:"timestamp"`
}

// OpHaltStatus models the operational halt status of a security
type OpHaltStatus struct {
	IsHalted  bool      `json:"isHalted"`
	Timestamp EpochTime `json:"timestamp"`
}

// SSRStatus models the short sale price test status for a security
type SSRStatus struct {
	IsSSR     bool      `json:"isSSR"`
	Detail    string    `json:"detail"`
	Timestamp EpochTime `json:"timestamp"`
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

// Company models the company data from the /company endpoint.
type Company struct {
	Symbol      string    `json:"symbol"`
	Name        string    `json:"companyName"`
	Exchange    string    `json:"exchange"`
	Industry    string    `json:"industry"`
	Website     string    `json:"website"`
	Description string    `json:"description"`
	CEO         string    `json:"CEO"`
	IssueType   IssueType `json:"issueType"`
	Sector      string    `json:"sector"`
	Employees   int       `json:"employees"`
	Tags        []string  `json:"tags"`
}

// Dividend models one dividend.
type Dividend struct {
	ExDate       Date    `json:"exDate"`
	PaymentDate  Date    `json:"paymentDate"`
	RecordDate   Date    `json:"recordDate"`
	DeclaredDate Date    `json:"declaredDate"`
	Amount       float64 `json:"amount"`
	Flag         string  `json:"flag"`
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
	ActualEPS            float64      `json:"actualEPS"`
	ConsensusEPS         float64      `json:"consensusEPS"`
	AnnounceTime         AnnounceTime `json:"announcetime"`
	NumberOfEstimates    int          `json:"numberOfEstimates"`
	EPSSurpriseDollar    float64      `json:"EPSSurpriseDollar"`
	EPSReportDate        Date         `json:"EPSReportDate"`
	FiscalPeriod         string       `json:"fiscalPeriod"`
	FiscalEndDate        Date         `json:"fiscalEndDate"`
	YearAgo              float64      `json:"yearAgo"`
	YearAgoChangePercent float64      `json:"yearAgoChangePercent"`
}

// EarningsToday models the earning that will be reported today as two arrays:
// before the open and after market close. Each array contains an object with
// all keys from earnings, a quote object, and a headline key.
type EarningsToday struct {
	BeforeOpen []TodayEarning `json:"bto"`
	AfterClose []TodayEarning `json:"amc"`
}

// TodayEarning models a single earning being reported today containing all
// keys from earnings, a quote object, and a headline.
type TodayEarning struct {
	Earning
	EstimatedChangePercent float64 `json:"estimatedChangePercent"`
	SymbolID               int     `json:"symbolId"`
	Symbol                 string  `json:"symbol"`
	Quote                  Quote   `json:"quote"`
	Headline               string  `json:"headline"`
}

// DelayedQuote returns the 15 minute delayed market quote.
type DelayedQuote struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float64 `json:"delayedPrice"`
	DelayedSize      int     `json:"delayedSize"`
	DelayedPriceTime int     `json:"delayedPriceTime"`
	High             float64 `json:"High"`
	Low              float64 `json:"Low"`
	TotalVolume      int     `json:"totalVolume"`
	ProcessedTime    int     `json:"processedTime"`
}

// EffectiveSpread models the effective spread, eligible volume, and price
// improvement of a stock by market.
type EffectiveSpread struct {
	Volume           int     `json:"volume"`
	Venue            string  `json:"venue"`
	VenueName        string  `json:"venueName"`
	EffectiveSpread  float64 `json:"effectiveSpread"`
	EffectiveQuoted  float64 `json:"effectiveQuoted"`
	PriceImprovement float64 `json:"priceImprovement"`
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
}

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
	CurrentDebt            float64 `json:"currentDebt"`
	ShortTermDebt          float64 `json:"shortTermDebt"`
	LongTermDebt           float64 `json:"LongTermDebt"`
	TotalCash              float64 `json:"totalCash"`
	TotalDebt              float64 `json:"totalDebt"`
	ShareholderEquity      float64 `json:"shareholderEquity"`
	CashChange             float64 `json:"cashChange"`
	CashFlow               float64 `json:"cashFlow"`
	OperatingGainsLosses   string  `json:"operatingGainsLosses"`
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

// HistoricalPrice models the data for a historical stock price.
type HistoricalPrice struct {
	Date string `json:"date"`
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

// InsiderRoster models the top 10 insiders with the most recent information.
type InsiderRoster struct {
	EntityName string `json:"entityName"`
	Position   int    `json:"position"`
	ReportDate Date   `json:"reportDate"`
}

// InsiderSummary models a summary of insider information.
type InsiderSummary struct {
	Name           string `json:"fullName"`
	NetTransaction int    `json:"netTransaction"`
	ReportedTitle  string `json:"reportedTitle"`
	TotalBought    int    `json:"totalBought"`
	TotalSold      int    `json:"totalSold"`
}

// InsiderTransaction models a buy or sell transaction made by an insider of a
// company.
type InsiderTransaction struct {
	EffectiveDate EpochTime `json:"effectiveDate"`
	Name          string    `json:"fullName"`
	ReportedTitle string    `json:"reportedTitle"`
	Price         float64   `json:"tranPrice"`
	Shares        int       `json:"tranShares"`
	Value         float64   `json:"tranValue"`
}

// InstitutionalOwner models an institutional owner of a stock.
type InstitutionalOwner struct {
	EntityName          string    `json:"entityProperName"`
	AdjustedHolding     float64   `json:"adjHolding"`
	AdjustedMarketValue float64   `json:"adjMv"`
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
	Avg30Volume         float64 `json:"avg30Volume"`
	Avg10Volume         float64 `json:"avg10Volume"`
	Float               float64 `json:"float"`
	Employees           int     `json:"employees"`
	TTMEPS              float64 `json:"ttmEPS"`
	TTMDividendRate     float64 `json:"ttmDividendRate"`
	DividendYield       float64 `json:"dividendYield"`
	NextDividendDate    Date    `json:"nextDividendDate"`
	ExDividendDate      Date    `json:"exDividendDate"`
	NextEarningsDate    Date    `json:"nextEarningsDate"`
	PERatio             float64 `json:"peRatio"`
	Beta                float64 `json:"beta"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
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

// LargestTrade models the 15 minute delayed, last sale eligible trades.
type LargestTrade struct {
	Price     float64 `json:"price"`
	Size      int     `json:"size"`
	Time      int     `json:"time"`
	TimeLabel string  `json:"timeLabel"`
	Venue     string  `json:"venue"`
	VenueName string  `json:"venueName"`
}

// Logo models the /logo endpoint.
type Logo struct {
	URL string `json:"url"`
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

// News models a news item either for the market or for an individual stock.
type News struct {
	Time       EpochTime `json:"datetime"`
	Headline   string    `json:"headline"`
	Source     string    `json:"source"`
	URL        string    `json:"url"`
	Summary    string    `json:"summary"`
	Related    string    `json:"related"`
	Image      string    `json:"image"`
	Language   string    `json:"lang"`
	HasPaywall bool      `json:"hasPaywall"`
}

// OHLC models the open, high, low, close for a stock.
type OHLC struct {
	Open  OpenClose `json:"open"`
	Close OpenClose `json:"close"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
}

// OpenClose provides the price and time for either the open or close price of
// a stock.
type OpenClose struct {
	Price float64 `json:"price"`
	Time  int     `json:"Time"`
}

// PreviousDay models the previous day adjusted price data.
type PreviousDay struct {
	Symbol           string  `json:"symbol"`
	Date             Date    `json:"date"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"Low"`
	Close            float64 `json:"close"`
	Volume           int     `json:"volume"`
	UnadjustedVolume int     `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
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

// Quote models the data returned from the IEX Cloud /quote endpoint.
type Quote struct {
	Symbol                string    `json:"symbol"`
	CompanyName           string    `json:"companyName"`
	CalculationPrice      string    `json:"calculationPrice"`
	Open                  float64   `json:"open"`
	OpenTime              EpochTime `json:"openTime"`
	Close                 float64   `json:"close"`
	CloseTime             EpochTime `json:"closeTime"`
	High                  float64   `json:"high"`
	Low                   float64   `json:"low"`
	LatestPrice           float64   `json:"latestPrice"`
	LatestSource          string    `json:"latestSource"`
	LatestTime            string    `json:"latestTime"`
	LatestUpdate          EpochTime `json:"latestUpdate"`
	LatestVolume          int       `json:"latestVolume"`
	IEXRealtimePrice      float64   `json:"iexRealtimePrice"`
	IEXRealtimeSize       int       `json:"iexRealtimeSize"`
	IEXLastUpdated        EpochTime `json:"iexLastUpdated"`
	DelayedPrice          float64   `json:"delayedPrice"`
	DelayedPriceTime      EpochTime `json:"delayedPriceTime"`
	ExtendedPrice         float64   `json:"extendedPrice"`
	ExtendedChange        float64   `json:"extendedChange"`
	ExtendedChangePercent float64   `json:"extendedChangePercent"`
	ExtendedPriceTime     EpochTime `json:"extendedPriceTime"`
	PreviousClose         float64   `json:"previousClose"`
	Change                float64   `json:"change"`
	ChangePercent         float64   `json:"changePercent"`
	IEXMarketPercent      float64   `json:"iexMarketPercent"`
	IEXVolume             int       `json:"iexVolume"`
	AvgTotalVolume        int       `json:"avgTotalVolume"`
	IEXBidPrice           float64   `json:"iexBidPrice"`
	IEXBidSize            int       `json:"iexBidSize"`
	IEXAskPrice           float64   `json:"iexAskPrice"`
	IEXAskSize            int       `json:"iexAskSize"`
	MarketCap             int       `json:"marketCap"`
	Week52High            float64   `json:"week52High"`
	Week52Low             float64   `json:"week52Low"`
	YTDChange             float64   `json:"ytdChange"`
	PERatio               float64   `json:"peRatio"`
}

// Recommendation models the buy, hold, sell recommendations for a stock.
type Recommendation struct {
	ConsensusEndDate   EpochTime `json:"consensusEndDate"`
	ConsensusStartDate EpochTime `json:"consensusStartDate"`
	BuyRatings         int       `json:"ratingBuy"`
	HoldRatings        int       `json:"ratingHold"`
	NoRatings          int       `json:"ratingNone"`
	OverweightRatings  int       `json:"ratingOverweight"`
	SellRatings        int       `json:"ratingSell"`
	UnderweightRatings int       `json:"ratingUnderweight"`
	ConsensusRating    float64   `json:"ratingScaleMark"`
}

// RelevantStocks models a list of relevant stocks that may or may not be
// peers.
type RelevantStocks struct {
	Peers   bool     `json:"peers"`
	Symbols []string `json:"symbols"`
}

// SectorPerformance models the performance based on each sector ETF.
type SectorPerformance struct {
	Type        string    `json:"sector"`
	Name        string    `json:"name"`
	Performance float64   `json:"performance"`
	LastUpdated EpochTime `json:"lastUpdated"`
}

// Split models the a stock split.
type Split struct {
	ExDate       Date    `json:"exDate"`
	DeclaredDate Date    `json:"declaredDate"`
	Ratio        float64 `json:"ratio"`
	FromFactor   float64 `json:"fromFactor"`
	Description  string  `json:"description"`
}

// Volume models the 15 minute delayed and 30 day average consolidated volume
// percentage of a stock by market.
type Volume struct {
	Volume               int     `json:"volume"`
	Venue                string  `json:"venue"`
	VenueName            string  `json:"venueName"`
	Date                 Date    `json:"date"`
	MarketPercent        float64 `json:"marketPercent"`
	AverageMarketPercent float64 `json:"avgMarketPercent"`
}
