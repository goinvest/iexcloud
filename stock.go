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

// Company models the company data from the /company endpoint.
type Company struct {
	Symbol         string   `json:"symbol"`
	Name           string   `json:"companyName"`
	Exchange       string   `json:"exchange"`
	Industry       string   `json:"industry"`
	Website        string   `json:"website"`
	Description    string   `json:"description"`
	CEO            string   `json:"CEO"`
	IssueType      string   `json:"issueType"`
	Sector         string   `json:"sector"`
	Employees      int      `json:"employees"`
	Tags           []string `json:"tags"`
	SecurityName   string   `json:"securityName"`
	PrimarySICCode int      `json:"primarySicCode"`
	Address        string   `json:"address"`
	Address2       string   `json:"address2"`
	State          string   `json:"state"`
	City           string   `json:"city"`
	Zip            string   `json:"zip"`
	Country        string   `json:"country"`
	Phone          string   `json:"phone"`
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

// HistoricalPrice models the data for a historical stock price.
type HistoricalPrice struct {
	Date string `json:"date"`
}

// IntradayPrice models the data for an aggregated intraday price in one minute
// buckets.
type IntradayPrice struct {
	Date                 Date       `json:"date"`
	Minute               HourMinute `json:"minute"`
	Label                string     `json:"label"`
	MarketOpen           float64    `json:"marketOpen"`
	MarketClose          float64    `json:"marketClose"`
	MarketHigh           float64    `json:"marketHigh"`
	MarketLow            float64    `json:"marketLow"`
	MarketAverage        float64    `json:"marketAverage"`
	MarketVolume         int        `json:"marketVolume"`
	MarketNotional       float64    `json:"marketNotional"`
	MarketNumTrades      int        `json:"marketNumberOfTrades"`
	MarketChangeOverTime float64    `json:"marketChangeOverTime"`
	High                 float64    `json:"High"`
	Low                  float64    `json:"Low"`
	Open                 float64    `json:"Open"`
	Close                float64    `json:"Close"`
	Average              float64    `json:"average"`
	Volume               int        `json:"volume"`
	Notional             float64    `json:"notional"`
	NumTrades            int        `json:"numberOfTrades"`
	ChangeOverTime       float64    `json:"changeOverTime"`
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

// OpenClose provides the price and time for either the open or close price of
// a stock.
type OpenClose struct {
	Price float64 `json:"price"`
	Time  int     `json:"Time"`
}

// OHLC models the open, high, low, close for a stock.
type OHLC struct {
	Open  OpenClose `json:"open"`
	Close OpenClose `json:"close"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
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

// Quote models the data returned from the IEX Cloud /quote endpoint.
type Quote struct {
	Symbol                string    `json:"symbol,omitempty"`
	CompanyName           string    `json:"companyName,omitempty"`
	PrimaryExchange       string    `json:"primaryExchange,omitempty"`
	CalculationPrice      string    `json:"calculationPrice,omitempty"`
	Open                  float64   `json:"open,omitempty"`
	OpenTime              EpochTime `json:"openTime,omitempty"`
	OpenSource            string    `json:"openSource,omitempty"`
	Close                 float64   `json:"close,omitempty"`
	CloseTime             EpochTime `json:"closeTime,omitempty"`
	CloseSource           string    `json:"closeSource,omitempty"`
	High                  float64   `json:"high,omitempty"`
	HighTime              EpochTime `json:"highTime,omitempty"`
	HighSource            string    `json:"highSource,omitempty"`
	Low                   float64   `json:"low,omitempty"`
	LowTime               EpochTime `json:"lowTime,omitempty"`
	LowSource             string    `json:"lowSource,omitempty"`
	LatestPrice           float64   `json:"latestPrice,omitempty"`
	LatestSource          string    `json:"latestSource,omitempty"`
	LatestTime            string    `json:"latestTime,omitempty"`
	LatestUpdate          EpochTime `json:"latestUpdate,omitempty"`
	LatestVolume          int       `json:"latestVolume,omitempty"`
	IEXRealtimePrice      float64   `json:"iexRealtimePrice,omitempty"`
	IEXRealtimeSize       int       `json:"iexRealtimeSize,omitempty"`
	IEXLastUpdated        EpochTime `json:"iexLastUpdated,omitempty"`
	DelayedPrice          float64   `json:"delayedPrice,omitempty"`
	DelayedPriceTime      EpochTime `json:"delayedPriceTime,omitempty"`
	ExtendedPrice         float64   `json:"extendedPrice,omitempty"`
	ExtendedChange        float64   `json:"extendedChange,omitempty"`
	ExtendedChangePercent float64   `json:"extendedChangePercent,omitempty"`
	ExtendedPriceTime     EpochTime `json:"extendedPriceTime,omitempty"`
	PreviousClose         float64   `json:"previousClose,omitempty"`
	Change                float64   `json:"change,omitempty"`
	ChangePercent         float64   `json:"changePercent,omitempty"`
	IEXMarketPercent      float64   `json:"iexMarketPercent,omitempty"`
	IEXVolume             int       `json:"iexVolume,omitempty"`
	AvgTotalVolume        int       `json:"avgTotalVolume,omitempty"`
	IEXBidPrice           float64   `json:"iexBidPrice,omitempty"`
	IEXBidSize            int       `json:"iexBidSize,omitempty"`
	IEXAskPrice           float64   `json:"iexAskPrice,omitempty"`
	IEXAskSize            int       `json:"iexAskSize,omitempty"`
	MarketCap             int       `json:"marketCap,omitempty"`
	Week52High            float64   `json:"week52High,omitempty"`
	Week52Low             float64   `json:"week52Low,omitempty"`
	YTDChange             float64   `json:"ytdChange,omitempty"`
	PERatio               float64   `json:"peRatio,omitempty"`
}

// VenueVolume models the 15 minute delayed and 30 day average consolidated
// volume percentage of a stock by market.
type VenueVolume struct {
	Volume               int     `json:"volume"`
	Venue                string  `json:"venue"`
	VenueName            string  `json:"venueName"`
	Date                 Date    `json:"date"`
	MarketPercent        float64 `json:"marketPercent"`
	AverageMarketPercent float64 `json:"avgMarketPercent"`
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

// InsiderRoster models the top 10 insiders with the most recent information.
type InsiderRoster struct {
	EntityName string    `json:"entityName"`
	Position   float64   `json:"position"`
	ReportDate EpochTime `json:"reportDate"`
}

// InsiderSummary models a summary of insider information.
type InsiderSummary struct {
	Name           string    `json:"fullName"`
	NetTransaction int       `json:"netTransaction"`
	ReportedTitle  string    `json:"reportedTitle"`
	TotalBought    int       `json:"totalBought"`
	TotalSold      int       `json:"totalSold"`
	Updated        EpochTime `json:"updated"`
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

// Logo models the /logo endpoint.
type Logo struct {
	URL string `json:"url"`
}

// RelevantStocks models a list of relevant stocks that may or may not be
// peers.
type RelevantStocks struct {
	Peers   bool     `json:"peers"`
	Symbols []string `json:"symbols"`
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
	Week52HighDate           Date    `json:"week52highDate"`
	Week52LowDate            Date    `json:"week52lowDate"`
	PutCallRatio             float64 `json:"putCallRatio"`
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
	SymbolID string `json:"symbolId"`
}

// UpcomingEvents is all of the upcoming events.
type UpcomingEvents struct {
	IPOs      IPOCalendar       `json:"ipos"`
	Earnings  []UpcomingEarning `json:"earnings"`
	Dividends []Dividend        `json:"dividends"`
	Splits    []Split           `json:"splits"`
}
