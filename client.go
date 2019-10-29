// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const apiURL = "https://cloud.iexapis.com/beta"

// Client models a client to consume the IEX Cloud API.
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// NewClient creates a client with the given authorization toke.
func NewClient(token string, options ...func(*Client)) *Client {
	client := &Client{
		token:      token,
		httpClient: &http.Client{},
	}

	// apply options
	for _, option := range options {
		option(client)
	}

	// set default values
	if client.baseURL == "" {
		client.baseURL = apiURL
	}

	return client
}

// WithHTTPClient sets the http.Client for a new IEX Client
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

// WithBaseURL sets the baseURL for a new IEX Client
func WithBaseURL(baseURL string) func(*Client) {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(ctx context.Context, endpoint string, v interface{}) error {
	address, err := c.addToken(endpoint)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Even if GET didn't return an error, check the status code to make sure
	// everything was ok.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

// GetJSONWithoutToken gets the JSON data from the given endpoint without
// adding a token to the URL.
func (c *Client) GetJSONWithoutToken(ctx context.Context, endpoint string, v interface{}) error {
	address := c.baseURL + endpoint
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Even if GET didn't return an error, check the status code to make sure
	// everything was ok.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) addToken(endpoint string) (string, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return "", err
	}
	v := u.Query()
	v.Add("token", c.token)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

// GetFloat64 gets the number from the given endpoint.
func (c *Client) GetFloat64(ctx context.Context, endpoint string) (float64, error) {
	address := c.baseURL + endpoint + "?token=" + c.token
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return 0.0, err
	}
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	// Even if GET didn't return an error, check the status code to make sure
	// everything was ok.
	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(string(b), 64)

}

// # Account related endpoints. #

// AccountMetadata returns information about an IEX Cloud account, such as
// current tier, payment status, message quote usage, etc. An SK token is
// required to access.
func (c Client) AccountMetadata(ctx context.Context) (AccountMetadata, error) {
	// FIXME(mdr): Since this requires an SK token, should Client be modified to
	// have an SK token? Should we change the token get the JSON and then change
	// it back? Need to think about this.
	r := AccountMetadata{}
	endpoint := "/account/metadata"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// Usage retrieves the current month usage for your account.
func (c Client) Usage(ctx context.Context) (Usage, error) {
	r := Usage{}
	endpoint := "/account/usage"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// # Stocks related endpoints. #

// AdvancedStats returns the everything in key stats plus additional advanced
// stats such as EBITDA, ratios, key financial data, and more.
func (c Client) AdvancedStats(ctx context.Context, symbol string) (AdvancedStats, error) {
	stats := AdvancedStats{}
	endpoint := fmt.Sprintf("/stock/%s/advanced-stats", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &stats)
	return stats, err
}

// AnnualBalanceSheets returns the specified number of most recent annual
// balance sheets from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualBalanceSheets(ctx context.Context, symbol string, num int) (BalanceSheets, error) {
	endpoint := fmt.Sprintf("/stock/%s/balance-sheet/%d?period=annual",
		url.PathEscape(symbol), num)
	return c.balanceSheets(ctx, endpoint)
}

// QuarterlyBalanceSheets returns the specified number of most recent quarterly
// balance sheets from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyBalanceSheets(ctx context.Context, symbol string, num int) (BalanceSheets, error) {
	endpoint := fmt.Sprintf("/stock/%s/balance-sheet/%d?period=quarter",
		url.PathEscape(symbol), num)
	return c.balanceSheets(ctx, endpoint)
}

func (c Client) balanceSheets(ctx context.Context, endpoint string) (BalanceSheets, error) {
	bs := BalanceSheets{}
	err := c.GetJSON(ctx, endpoint, &bs)
	return bs, err
}

// Book returns the quote, bids, asks, and trades for a given stock symbol.
func (c Client) Book(ctx context.Context, symbol string) (Book, error) {
	book := Book{}
	endpoint := fmt.Sprintf("/stock/%s/book", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &book)
	return book, err
}

// AnnualCashFlows returns the specified number of most recent annual cash flow
// statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualCashFlows(ctx context.Context, symbol string, num int) (CashFlows, error) {
	endpoint := fmt.Sprintf("/stock/%s/cash-flow/%d?period=annual",
		url.PathEscape(symbol), num)
	return c.cashFlows(ctx, endpoint)
}

// QuarterlyCashFlows returns the specified number of most recent annual
// cash flow statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyCashFlows(ctx context.Context, symbol string, num int) (CashFlows, error) {
	endpoint := fmt.Sprintf("/stock/%s/cash-flow/%d?period=quarter",
		url.PathEscape(symbol), num)
	return c.cashFlows(ctx, endpoint)
}

func (c Client) cashFlows(ctx context.Context, endpoint string) (CashFlows, error) {
	cf := CashFlows{}
	err := c.GetJSON(ctx, endpoint, &cf)
	return cf, err
}

// CollectionBySector returns an array of quote objects for all
// symbols within the specified sector.
func (c Client) CollectionBySector(ctx context.Context, sector Sector) ([]Quote, error) {
	quotes := []Quote{}
	endpoint := fmt.Sprintf("/stock/market/collection/sector?collectionName=%s",
		url.QueryEscape(sector.Name))
	err := c.GetJSON(ctx, endpoint, &quotes)
	return quotes, err
}

// CollectionByTag returns an array of quote objects for all
// symbols within the specified tag collection.
func (c Client) CollectionByTag(ctx context.Context, tag Tag) ([]Quote, error) {
	quotes := []Quote{}
	endpoint := fmt.Sprintf("/stock/market/collection/tag?collectionName=%s",
		url.QueryEscape(tag.Name))
	err := c.GetJSON(ctx, endpoint, &quotes)
	return quotes, err
}

// Company returns the copmany data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Company(ctx context.Context, symbol string) (Company, error) {
	company := Company{}
	endpoint := fmt.Sprintf("/stock/%s/company", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &company)
	return company, err
}

// Dividends returns the dividends from the IEX Cloud endpoint for the given
// stock symbol and the given date range.
func (c Client) Dividends(ctx context.Context, symbol string, r PathRange) ([]Dividend, error) {
	dividends := []Dividend{}
	endpoint := fmt.Sprintf("/stock/%s/dividends/%s",
		url.PathEscape(symbol), PathRangeJSON[r])
	err := c.GetJSON(ctx, endpoint, &dividends)
	return dividends, err
}

// Earnings returns the specified number of most recent earnings data from the
// IEX Cloud endpoint for the given stock symbol.
func (c Client) Earnings(ctx context.Context, symbol string, num int) (Earnings, error) {
	earnings := Earnings{}
	endpoint := fmt.Sprintf("/stock/%s/earnings/%d", url.PathEscape(symbol), num)
	err := c.GetJSON(ctx, endpoint, &earnings)
	return earnings, err
}

// EarningsToday returns the earnings that will be reported today before the
// open and after the market closes.
func (c Client) EarningsToday(ctx context.Context) (EarningsToday, error) {
	e := EarningsToday{}
	endpoint := "/stock/market/today-earnings"
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// DelayedQuote returns the 15 minute delayed market quote from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) DelayedQuote(ctx context.Context, symbol string) (DelayedQuote, error) {
	dq := DelayedQuote{}
	endpoint := fmt.Sprintf("/stock/%s/delayed-quote", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &dq)
	return dq, err
}

// EffectiveSpreads returns the effective spreads from the IEX Cloud endpoint
// for the given stock symbol.
func (c Client) EffectiveSpreads(ctx context.Context, symbol string) ([]EffectiveSpread, error) {
	es := []EffectiveSpread{}
	endpoint := fmt.Sprintf("/stock/%s/effective-spread", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &es)
	return es, err
}

// Estimates returns the latest consensue estimates for the next fiscal period.
func (c Client) Estimates(ctx context.Context, symbol string, num int) (Estimates, error) {
	estimates := Estimates{}
	endpoint := fmt.Sprintf("/stock/%s/estimates/%d", url.PathEscape(symbol), num)
	err := c.GetJSON(ctx, endpoint, &estimates)
	return estimates, err
}

// AnnualFinancials returns the specified number of most recent annual
// financials from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualFinancials(ctx context.Context, symbol string, num int) (Financials, error) {
	endpoint := fmt.Sprintf("/stock/%s/financials/%d?period=annual",
		url.PathEscape(symbol), num)
	return c.financials(ctx, endpoint)
}

// QuarterlyFinancials returns the specified number of most recent quarterly
// financials from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyFinancials(ctx context.Context, symbol string, num int) (Financials, error) {
	endpoint := fmt.Sprintf("/stock/%s/financials/%d?period=quarter",
		url.PathEscape(symbol), num)
	return c.financials(ctx, endpoint)
}

func (c Client) financials(ctx context.Context, endpoint string) (Financials, error) {
	financials := Financials{}
	err := c.GetJSON(ctx, endpoint, &financials)
	return financials, err
}

// FundOwnership returns the ten top holders of the given stock.
func (c Client) FundOwnership(ctx context.Context, symbol string) ([]FundOwner, error) {
	r := []FundOwner{}
	endpoint := fmt.Sprintf("/stock/%s/fund-ownership", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// AnnualIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualIncomeStatements(ctx context.Context, symbol string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d?period=annual",
		url.PathEscape(symbol), num)
	return c.incomeStatements(ctx, endpoint)
}

// QuarterlyIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyIncomeStatements(ctx context.Context, symbol string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d?period=quarter",
		url.PathEscape(symbol), num)
	return c.incomeStatements(ctx, endpoint)
}

func (c Client) incomeStatements(ctx context.Context, endpoint string) (IncomeStatements, error) {
	is := IncomeStatements{}
	err := c.GetJSON(ctx, endpoint, &is)
	return is, err
}

// InsiderRoster returns the top 10 insiders with the most recent information
// for the given stock symbol.
func (c Client) InsiderRoster(ctx context.Context, symbol string) ([]InsiderRoster, error) {
	r := []InsiderRoster{}
	endpoint := fmt.Sprintf("/stock/%s/insider-roster", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// InsiderSummary returns the insiders summary with the most recent information
// for the given stock symbol.
func (c Client) InsiderSummary(ctx context.Context, symbol string) ([]InsiderSummary, error) {
	r := []InsiderSummary{}
	endpoint := fmt.Sprintf("/stock/%s/insider-summary", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// InsiderTransactions returns a list of insider transactions for the given stock symbol.
func (c Client) InsiderTransactions(ctx context.Context, symbol string) ([]InsiderTransaction, error) {
	r := []InsiderTransaction{}
	endpoint := fmt.Sprintf("/stock/%s/insider-transactions", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// InstitutionalOwnership returns the top 10 holders with the most recent
// information.
func (c Client) InstitutionalOwnership(ctx context.Context, symbol string) ([]InstitutionalOwner, error) {
	r := []InstitutionalOwner{}
	endpoint := fmt.Sprintf("/stock/%s/institutional-owernship", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// KeyStats returns the key stats from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) KeyStats(ctx context.Context, symbol string) (KeyStats, error) {
	stats := KeyStats{}
	endpoint := fmt.Sprintf("/stock/%s/stats", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &stats)
	return stats, err
}

// LargestTrades returns the 15 minute delayed, last sale eligible trade from
// the IEX Cloud endpoint for the given stock symbol.
func (c Client) LargestTrades(ctx context.Context, symbol string) ([]LargestTrade, error) {
	lt := []LargestTrade{}
	endpoint := fmt.Sprintf("/stock/%s/largest-trades", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &lt)
	return lt, err
}

// MostActive returns a list of quotes for the top 10 most active stocks from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) MostActive(ctx context.Context) ([]Quote, error) {
	return c.list(ctx, "mostactive")
}

// Gainers returns a list of quotes for the top 10 stock gainers from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) Gainers(ctx context.Context) ([]Quote, error) {
	return c.list(ctx, "gainers")
}

// Losers returns a list of quotes for the top 10 stock losers from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) Losers(ctx context.Context) ([]Quote, error) {
	return c.list(ctx, "losers")
}

// IEXVolume returns a list of quotes for the top 10 IEX stocks by volume from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) IEXVolume(ctx context.Context) ([]Quote, error) {
	return c.list(ctx, "iexvolume")
}

// IEXPercent returns a list of quotes for the top 10 IEX stocks by percent
// from the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) IEXPercent(ctx context.Context) ([]Quote, error) {
	return c.list(ctx, "iexpercent")
}

// InFocus returns a list of quotes for the top 10 in focus stocks from the IEX
// Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) InFocus(ctx context.Context) ([]Quote, error) {
	return c.list(ctx, "infocus")
}

func (c Client) list(ctx context.Context, list string) ([]Quote, error) {
	q := []Quote{}
	endpoint := "/stock/market/list/" + list
	err := c.GetJSON(ctx, endpoint, &q)
	return q, err
}

// Logo returns the logo data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Logo(ctx context.Context, symbol string) (Logo, error) {
	logo := Logo{}
	endpoint := fmt.Sprintf("/stock/%s/logo", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &logo)
	return logo, err
}

// Markets returns real time traded volume on U.S. markets.
func (c Client) Markets(ctx context.Context) ([]Market, error) {
	m := []Market{}
	endpoint := "/market"
	err := c.GetJSON(ctx, endpoint, &m)
	return m, err
}

// News retrieves the given number of news articles for the given stock symbol.
func (c Client) News(ctx context.Context, symbol string, num int) ([]News, error) {
	n := []News{}
	endpoint := fmt.Sprintf("/stock/%s/news/last/%d",
		url.PathEscape(symbol), num)
	err := c.GetJSON(ctx, endpoint, &n)
	return n, err
}

// MarketNews retrieves the given number of news articles for the market.
func (c Client) MarketNews(ctx context.Context, num int) ([]News, error) {
	n := []News{}
	endpoint := fmt.Sprintf("/stock/market/news/last/%d", num)
	err := c.GetJSON(ctx, endpoint, &n)
	return n, err
}

// OHLC returns the OHLC data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) OHLC(ctx context.Context, symbol string) (OHLC, error) {
	ohlc := OHLC{}
	endpoint := fmt.Sprintf("/stock/%s/ohlc", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &ohlc)
	return ohlc, err
}

// Peers returns a slice of peer stock symbols from the IEX Cloud endpoint for
// the given stock symbol.
func (c Client) Peers(ctx context.Context, symbol string) ([]string, error) {
	peers := []string{}
	endpoint := fmt.Sprintf("/stock/%s/peers", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &peers)
	return peers, err
}

// PreviousDay returns the previous day adjusted price data from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) PreviousDay(ctx context.Context, symbol string) (PreviousDay, error) {
	pd := PreviousDay{}
	endpoint := fmt.Sprintf("/stock/%s/previous", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &pd)
	return pd, err
}

// Price returns the current stock price from the IEX Cloud endpoint for the
// given stock symbol.
func (c Client) Price(ctx context.Context, symbol string) (float64, error) {
	endpoint := fmt.Sprintf("/stock/%s/price", url.PathEscape(symbol))
	return c.GetFloat64(ctx, endpoint)
}

// PriceTarget returns the latest average, high, and low analyst price target
// for a given stock symbol.
func (c Client) PriceTarget(ctx context.Context, symbol string) (PriceTarget, error) {
	pt := PriceTarget{}
	endpoint := fmt.Sprintf("/stock/%s/price-target", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &pt)
	return pt, err
}

// Quote returns the quote data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) Quote(ctx context.Context, symbol string) (Quote, error) {
	r := Quote{}
	endpoint := fmt.Sprintf("/stock/%s/quote", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// RecommendationTrends provides a list of recommendations with the start and
// end date for each rating.
func (c Client) RecommendationTrends(ctx context.Context, symbol string) ([]Recommendation, error) {
	r := []Recommendation{}
	endpoint := fmt.Sprintf("/stock/%s/recommendation-trends", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// RelevantStocks is similar to the peers endpoint, except this will return
// most active market symbols when peers are not available. If the symbols
// returned are not peers, the peers key will be false. This is not intended to
// represent a definitive or accurate list of peers, and is subject to change
// at any time.
func (c Client) RelevantStocks(ctx context.Context, symbol string) (RelevantStocks, error) {
	r := RelevantStocks{}
	endpoint := fmt.Sprintf("/stock/%s/relevant", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// SectorPerformance returns the performance of each sector for the current
// trading day. Performance is based on each sector ETF.
func (c Client) SectorPerformance(ctx context.Context) ([]Sector, error) {
	r := []Sector{}
	endpoint := "/stock/market/sector-performance"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// # Alternative Data related endpoints. #

// Crypto provides a quote for a given cryptocurrency symbol.
func (c Client) Crypto(ctx context.Context, symbol string) (CryptoQuote, error) {
	r := CryptoQuote{}
	endpoint := fmt.Sprintf("/crypto/%s/quote", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// CEOCompensation provides CEO compensation for the given stock symbol.
func (c Client) CEOCompensation(ctx context.Context, symbol string) (CEOCompensation, error) {
	r := CEOCompensation{}
	endpoint := fmt.Sprintf("/stock/%s/ceo-compensation", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// # Reference Data related endpoints. #

// Symbols returns an array of symbols that IEX Cloud supports for API calls.
func (c Client) Symbols(ctx context.Context) ([]Symbol, error) {
	symbols := []Symbol{}
	endpoint := "/ref-data/symbols"
	err := c.GetJSON(ctx, endpoint, &symbols)
	return symbols, err
}

// IEXSymbols returns an array of symbols the Investors Exchange supports for
// trading. This list is updated daily as of 7:45 a.m. ET. Symbols may be added
// or removed by the Investors Exchange after the list was produced.
func (c Client) IEXSymbols(ctx context.Context) ([]TradedSymbol, error) {
	symbols := []TradedSymbol{}
	endpoint := "/ref-data/iex/symbols"
	err := c.GetJSON(ctx, endpoint, &symbols)
	return symbols, err
}

// USExchanges returns an array of U.S. Exchanges.
func (c Client) USExchanges(ctx context.Context) ([]USExchange, error) {
	e := []USExchange{}
	endpoint := "/ref-data/market/us/exchanges"
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// NextTradingDay returns the date of the next trading day.
func (c Client) NextTradingDay(ctx context.Context) (TradeHolidayDate, error) {
	r := TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/trade/next/1"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// NextTradingDays returns the dates of the next trading days for the given
// number of days.
func (c Client) NextTradingDays(ctx context.Context, numDays int) (TradeHolidayDate, error) {
	r := TradeHolidayDate{}
	endpoint := fmt.Sprintf("/ref-data/us/dates/trade/next/%d", numDays)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// NextHoliday returns the date of the next holiday.
func (c Client) NextHoliday(ctx context.Context) (TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/holiday/next/1"
	err := c.GetJSON(ctx, endpoint, &r)
	return r[0], err
}

// NextHolidays returns the dates of the next holidays for the given
// number of days.
func (c Client) NextHolidays(ctx context.Context, numDays int) ([]TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := fmt.Sprintf("/ref-data/us/dates/holiday/next/%d", numDays)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// PreviousTradingDay returns the date of the previous trading day.
func (c Client) PreviousTradingDay(ctx context.Context) (TradeHolidayDate, error) {
	r := TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/trade/last/1"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// PreviousHoliday returns the date of the previous holiday.
func (c Client) PreviousHoliday(ctx context.Context) (TradeHolidayDate, error) {
	r := TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/holiday/last/1"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// Sectors returns an array of all sectors
func (c Client) Sectors(ctx context.Context) ([]Sector, error) {
	r := []Sector{}
	endpoint := "/ref-data/sectors"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// Tags returns an array of tags.  Tags can
// be found for each on each company.
func (c Client) Tags(ctx context.Context) ([]Tag, error) {
	r := []Tag{}
	endpoint := "/ref-data/tags"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// MutualFundSymbols returns an array of mutual funds that IEX Cloud supports
// for API calls.
func (c Client) MutualFundSymbols(ctx context.Context) ([]Symbol, error) {
	r := []Symbol{}
	endpoint := "/ref-data/mutual-funds/symbols"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// OTCSymbols returns an array of Over-the-Counter (OTC) stocks that IEX Cloud
// supports for API calls.
func (c Client) OTCSymbols(ctx context.Context) ([]Symbol, error) {
	r := []Symbol{}
	endpoint := "/ref-data/otc/symbols"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// FXSymbols returns a list of currencies and a list of foreign exchange
// currency pairs that are available supported by IEX Cloud.
func (c Client) FXSymbols(ctx context.Context) (FXSymbols, error) {
	r := FXSymbols{}
	endpoint := "/ref-data/fx/symbols"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// CryptoSymbols returns a list of cryptocurrencies that are supported by IEX
// Cloud.
func (c Client) CryptoSymbols(ctx context.Context) ([]CryptoSymbol, error) {
	r := []CryptoSymbol{}
	endpoint := "/ref-data/crypto/symbols"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// # Forex / Currencies related endpoints. #

// ExchangeRate returns an end of day exchange rate of a given currency pair.
func (c Client) ExchangeRate(ctx context.Context, from, to string) (ExchangeRate, error) {
	r := ExchangeRate{}
	endpoint := fmt.Sprintf("/fx/rate/%s/%s",
		url.PathEscape(from),
		url.PathEscape(to))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// # Investors Exchange Data related endpoints. #

// TOPS is used to receive real-time top of book quotations direct from IEX.
// The quotations received via TOPS provide an aggregated size and do not
// indicate the size or number of individual orders at the best bid or ask.
// Non-displayed orders and non-displayed portions of reserve orders are not
// represented in TOPS. TOPS also provides last trade price and size
// information. Trades resulting from either displayed or non-displayed orders
// matching on IEX will be reported.  Routed executions will not be reported.
func (c Client) TOPS(ctx context.Context, symbols []string) ([]TOPS, error) {
	r := []TOPS{}
	s := strings.Join(symbols, ",")
	endpoint := "/tops?symbols=" + url.PathEscape(s)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// OneTOPS returns TOPS for one stock symbol.
func (c Client) OneTOPS(ctx context.Context, symbol string) ([]TOPS, error) {
	r := []TOPS{}
	endpoint := "/tops?symbols=" + url.PathEscape(symbol)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// Last provides trade data for executions on IEX. It is a near real time,
// intraday API that provides IEX last sale price, size and time. Last is ideal
// for developers that need a lightweight stock quote.
func (c Client) Last(ctx context.Context, symbols []string) ([]Last, error) {
	r := []Last{}
	s := strings.Join(symbols, ",")
	endpoint := "/tops/last?symbols=" + url.PathEscape(s)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// OneLast provides the last trade data executions for one stock symbol.
func (c Client) OneLast(ctx context.Context, symbol string) ([]Last, error) {
	r := []Last{}
	endpoint := "/tops/last?symbols=" + url.PathEscape(symbol)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// DEEP provides all DEEP data for one stock symbol.
func (c Client) DEEP(ctx context.Context, symbol string) (DEEP, error) {
	r := DEEP{}
	endpoint := "/deep?symbols=" + url.PathEscape(symbol)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// DEEPBook provides DEEP book data for multiple symbols
func (c Client) DEEPBook(ctx context.Context, symbols []string) (map[string]DEEPBook, error) {
	r := make(map[string]DEEPBook)
	s := strings.Join(symbols, ",")
	endpoint := "/deep/book?symbols=" + url.PathEscape(s)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// DEEPTrades provides DEEP trades data for multiple symbols.
func (c Client) DEEPTrades(ctx context.Context, symbols []string) (map[string][]Trade, error) {
	r := make(map[string][]Trade)
	s := strings.Join(symbols, ",")
	endpoint := "/deep/trades?symbols=" + url.PathEscape(s)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// IntradayStats retrieves the intraday stats on IEX.
func (c Client) IntradayStats(ctx context.Context, symbol string) (IntradayStats, error) {
	r := IntradayStats{}
	endpoint := "/stats/intraday"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// # Historical Data related endpoints. #

// HistoricalPrices retrieves historically adjusted market-wide data
func (c Client) HistoricalPrices(ctx context.Context, symbol string, timeframe HistoricalTimeFrame, options *HistoricalOptions) ([]HistoricalDataPoint, error) {
	h := make([]HistoricalDataPoint, 0)
	if !timeframe.Valid() {
		return h, errors.New("invalid timeframe passed to method")
	}
	endpoint := fmt.Sprintf("/stock/%s/chart/%s",
		url.PathEscape(symbol), timeframe)

	endpoint, err := c.historicalEndpointWithOpts(endpoint, options)
	if err != nil {
		return h, err
	}
	err = c.GetJSON(ctx, endpoint, &h)
	return h, err
}

// HistoricalPricesByDay retrieves historically adjusted market-wide data for a given day
func (c Client) HistoricalPricesByDay(ctx context.Context, symbol string, day time.Time, options *HistoricalOptions) ([]HistoricalDataPoint, error) {
	h := make([]HistoricalDataPoint, 0)
	endpoint := fmt.Sprintf("/stock/%s/chart/date/%s",
		url.PathEscape(symbol), day.Format("20060102"))
	endpoint, err := c.historicalEndpointWithOpts(endpoint, options)
	if err != nil {
		return h, err
	}

	err = c.GetJSON(ctx, endpoint, &h)
	return h, err
}

func (c Client) historicalEndpointWithOpts(endpoint string, opts *HistoricalOptions) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}
	return endpoint, nil
}

// IntradayHistoricalPrices retrieves intraday historical market-wide data
func (c Client) IntradayHistoricalPrices(ctx context.Context, symbol string, options *IntradayHistoricalOptions) ([]IntradayHistoricalDataPoint, error) {
	h := make([]IntradayHistoricalDataPoint, 0)
	endpoint := fmt.Sprintf("/stock/%s/chart/1d",
		url.PathEscape(symbol))
	endpoint, err := c.intradayHistoricalEndpointWithOpts(endpoint, options, false)
	if err != nil {
		return h, err
	}

	err = c.GetJSON(ctx, endpoint, &h)
	return h, err
}

// IntradayHistoricalPricesByDay retrieves intraday historical market-wide data for a given day
func (c Client) IntradayHistoricalPricesByDay(ctx context.Context, symbol string, day time.Time, options *IntradayHistoricalOptions) ([]IntradayHistoricalDataPoint, error) {
	h := make([]IntradayHistoricalDataPoint, 0)
	endpoint := fmt.Sprintf("/stock/%s/chart/date/%s?chartByDay=true",
		url.PathEscape(symbol), day.Format("20060102"))
	endpoint, err := c.intradayHistoricalEndpointWithOpts(endpoint, options, true)
	if err != nil {
		return h, err
	}
	err = c.GetJSON(ctx, endpoint, &h)
	return h, err
}

func (c Client) intradayHistoricalEndpointWithOpts(endpoint string, opts *IntradayHistoricalOptions, existingParams bool) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	sep := "?"
	if existingParams {
		sep = "&"
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s%s%s", endpoint, sep, optParams)
	}
	return endpoint, nil
}

// # API System Metadata related endpoints. #

// Status returns the IEX Cloud system status.
func (c Client) Status(ctx context.Context) (Status, error) {
	status := Status{}
	endpoint := "/status"
	err := c.GetJSONWithoutToken(ctx, endpoint, &status)
	return status, err
}
