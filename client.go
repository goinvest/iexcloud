// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const apiURL = "https://cloud.iexapis.com/beta"

// Client models a client to consume the IEX Cloud API.
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// NewClient creates a client with the given authorization toke.
func NewClient(token string, baseURL string) *Client {
	if baseURL == "" {
		baseURL = apiURL
	}
	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(endpoint string, v interface{}) error {
	address, err := c.addToken(endpoint)
	if err != nil {
		return err
	}
	resp, err := http.Get(address)
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
func (c *Client) GetJSONWithoutToken(endpoint string, v interface{}) error {
	address := c.baseURL + endpoint
	resp, err := http.Get(address)
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
func (c *Client) GetFloat64(endpoint string) (float64, error) {
	address := c.baseURL + endpoint + "?token=" + c.token
	resp, err := http.Get(address)
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

// AnnualBalanceSheets returns the specified number of most recent annual
// balance sheets from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualBalanceSheets(stock string, num int) (BalanceSheets, error) {
	endpoint := fmt.Sprintf("/stock/%s/balance-sheet/%d?period=annual",
		url.PathEscape(stock), num)
	return c.balanceSheets(endpoint)
}

// QuarterlyBalanceSheets returns the specified number of most recent quarterly
// balance sheets from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyBalanceSheets(stock string, num int) (BalanceSheets, error) {
	endpoint := fmt.Sprintf("/stock/%s/balance-sheet/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.balanceSheets(endpoint)
}

func (c Client) balanceSheets(endpoint string) (BalanceSheets, error) {
	bs := BalanceSheets{}
	err := c.GetJSON(endpoint, &bs)
	return bs, err
}

// Book returns the quote, bids, asks, and trades for a given stock symbol.
func (c Client) Book(stock string) (Book, error) {
	book := Book{}
	endpoint := fmt.Sprintf("/stock/%s/book", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &book)
	return book, err
}

// AnnualCashFlows returns the specified number of most recent annual cash flow
// statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualCashFlows(stock string, num int) (CashFlows, error) {
	endpoint := fmt.Sprintf("/stock/%s/cash-flow/%d?period=annual",
		url.PathEscape(stock), num)
	return c.cashFlows(endpoint)
}

// QuarterlyCashFlows returns the specified number of most recent annual
// cash flow statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyCashFlows(stock string, num int) (CashFlows, error) {
	endpoint := fmt.Sprintf("/stock/%s/cash-flow/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.cashFlows(endpoint)
}

func (c Client) cashFlows(endpoint string) (CashFlows, error) {
	cf := CashFlows{}
	err := c.GetJSON(endpoint, &cf)
	return cf, err
}

// Company returns the copmany data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Company(stock string) (Company, error) {
	company := Company{}
	endpoint := fmt.Sprintf("/stock/%s/company", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &company)
	return company, err
}

// Dividends returns the dividends from the IEX Cloud endpoint for the given
// stock symbol and the given date range.
func (c Client) Dividends(stock string, r PathRange) ([]Dividend, error) {
	dividends := []Dividend{}
	endpoint := fmt.Sprintf("/stock/%s/dividends/%s",
		url.PathEscape(stock), PathRangeJSON[r])
	err := c.GetJSON(endpoint, &dividends)
	return dividends, err
}

// Earnings returns the specified number of most recent earnings data from the
// IEX Cloud endpoint for the given stock symbol.
func (c Client) Earnings(stock string, num int) (Earnings, error) {
	earnings := Earnings{}
	endpoint := fmt.Sprintf("/stock/%s/earnings/%d", url.PathEscape(stock), num)
	err := c.GetJSON(endpoint, &earnings)
	return earnings, err
}

// EarningsToday returns the earnings that will be reported today before the
// open and after the market closes.
func (c Client) EarningsToday() (EarningsToday, error) {
	e := EarningsToday{}
	endpoint := "/stock/market/today-earnings"
	err := c.GetJSON(endpoint, &e)
	return e, err
}

// DelayedQuote returns the 15 minute delayed market quote from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) DelayedQuote(stock string) (DelayedQuote, error) {
	dq := DelayedQuote{}
	endpoint := fmt.Sprintf("/stock/%s/delayed-quote", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &dq)
	return dq, err
}

// EffectiveSpreads returns the effective spreads from the IEX Cloud endpoint
// for the given stock symbol.
func (c Client) EffectiveSpreads(stock string) ([]EffectiveSpread, error) {
	es := []EffectiveSpread{}
	endpoint := fmt.Sprintf("/stock/%s/effective-spread", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &es)
	return es, err
}

// Estimates returns the latest consensue estimates for the next fiscal period.
func (c Client) Estimates(stock string, num int) (Estimates, error) {
	estimates := Estimates{}
	endpoint := fmt.Sprintf("/stock/%s/estimates/%d", url.PathEscape(stock), num)
	err := c.GetJSON(endpoint, &estimates)
	return estimates, err
}

// AnnualFinancials returns the specified number of most recent annual
// financials from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualFinancials(stock string, num int) (Financials, error) {
	endpoint := fmt.Sprintf("/stock/%s/financials/%d?period=annual",
		url.PathEscape(stock), num)
	return c.financials(endpoint)
}

// QuarterlyFinancials returns the specified number of most recent quarterly
// financials from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyFinancials(stock string, num int) (Financials, error) {
	endpoint := fmt.Sprintf("/stock/%s/financials/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.financials(endpoint)
}

func (c Client) financials(endpoint string) (Financials, error) {
	financials := Financials{}
	err := c.GetJSON(endpoint, &financials)
	return financials, err
}

// AnnualIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualIncomeStatements(stock string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d?period=annual",
		url.PathEscape(stock), num)
	return c.incomeStatements(endpoint)
}

// QuarterlyIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyIncomeStatements(stock string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d?period=quarter",
		url.PathEscape(stock), num)
	return c.incomeStatements(endpoint)
}

func (c Client) incomeStatements(endpoint string) (IncomeStatements, error) {
	is := IncomeStatements{}
	err := c.GetJSON(endpoint, &is)
	return is, err
}

// InsiderRoster returns the top 10 insiders with the most recent information
// for the given stock symbol.
func (c Client) InsiderRoster(stock string) ([]InsiderRoster, error) {
	r := []InsiderRoster{}
	endpoint := fmt.Sprintf("/stock/%s/insider-roster", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &r)
	return r, err
}

// InsiderSummary returns the insiders summary with the most recent information
// for the given stock symbol.
func (c Client) InsiderSummary(stock string) ([]InsiderSummary, error) {
	r := []InsiderSummary{}
	endpoint := fmt.Sprintf("/stock/%s/insider-summary", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &r)
	return r, err
}

// KeyStats returns the key stats from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) KeyStats(stock string) (KeyStats, error) {
	stats := KeyStats{}
	endpoint := fmt.Sprintf("/stock/%s/stats", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &stats)
	return stats, err
}

// LargestTrades returns the 15 minute delayed, last sale eligible trade from
// the IEX Cloud endpoint for the given stock symbol.
func (c Client) LargestTrades(stock string) ([]LargestTrade, error) {
	lt := []LargestTrade{}
	endpoint := fmt.Sprintf("/stock/%s/largest-trades", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &lt)
	return lt, err
}

// MostActive returns a list of quotes for the top 10 most active stocks from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) MostActive() ([]Quote, error) {
	return c.list("mostactive")
}

// Gainers returns a list of quotes for the top 10 stock gainers from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) Gainers() ([]Quote, error) {
	return c.list("gainers")
}

// Losers returns a list of quotes for the top 10 stock losers from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) Losers() ([]Quote, error) {
	return c.list("losers")
}

// IEXVolume returns a list of quotes for the top 10 IEX stocks by volume from
// the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) IEXVolume() ([]Quote, error) {
	return c.list("iexvolume")
}

// IEXPercent returns a list of quotes for the top 10 IEX stocks by percent
// from the IEX Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) IEXPercent() ([]Quote, error) {
	return c.list("iexpercent")
}

// InFocus returns a list of quotes for the top 10 in focus stocks from the IEX
// Cloud endpoint updated intrady, 15 minute delayed.
func (c Client) InFocus() ([]Quote, error) {
	return c.list("infocus")
}

func (c Client) list(list string) ([]Quote, error) {
	q := []Quote{}
	endpoint := "/stock/market/list/" + list
	err := c.GetJSON(endpoint, &q)
	return q, err
}

// Logo returns the logo data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Logo(stock string) (Logo, error) {
	logo := Logo{}
	endpoint := fmt.Sprintf("/stock/%s/logo", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &logo)
	return logo, err
}

// Markets returns real time traded volume on U.S. markets.
func (c Client) Markets() ([]Market, error) {
	m := []Market{}
	endpoint := "/market"
	err := c.GetJSON(endpoint, &m)
	return m, err
}

// News retrieves the given number of news articles for the given stock symbol.
func (c Client) News(stock string, num int) ([]News, error) {
	n := []News{}
	endpoint := fmt.Sprintf("/stock/%s/news/last/%d",
		url.PathEscape(stock), num)
	err := c.GetJSON(endpoint, &n)
	return n, err
}

// MarketNews retrieves the given number of news articles for the market.
func (c Client) MarketNews(num int) ([]News, error) {
	n := []News{}
	endpoint := fmt.Sprintf("/stock/market/news/last/%d", num)
	err := c.GetJSON(endpoint, &n)
	return n, err
}

// OHLC returns the OHLC data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) OHLC(stock string) (OHLC, error) {
	ohlc := OHLC{}
	endpoint := fmt.Sprintf("/stock/%s/ohlc", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &ohlc)
	return ohlc, err
}

// Peers returns a slice of peer stock symbols from the IEX Cloud endpoint for
// the given stock symbol.
func (c Client) Peers(stock string) ([]string, error) {
	peers := []string{}
	endpoint := fmt.Sprintf("/stock/%s/peers", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &peers)
	return peers, err
}

// PreviousDay returns the previous day adjusted price data from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) PreviousDay(stock string) (PreviousDay, error) {
	pd := PreviousDay{}
	endpoint := fmt.Sprintf("/stock/%s/previous", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &pd)
	return pd, err
}

// Price returns the current stock price from the IEX Cloud endpoint for the
// given stock symbol.
func (c Client) Price(stock string) (float64, error) {
	endpoint := fmt.Sprintf("/stock/%s/price", url.PathEscape(stock))
	return c.GetFloat64(endpoint)
}

// PriceTarget returns the latest average, high, and low analyst price target
// for a given stock symbol.
func (c Client) PriceTarget(stock string) (PriceTarget, error) {
	pt := PriceTarget{}
	endpoint := fmt.Sprintf("/stock/%s/price-target", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &pt)
	return pt, err
}

// Quote returns the quote data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) Quote(stock string) (Quote, error) {
	quote := Quote{}
	endpoint := fmt.Sprintf("/stock/%s/quote", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &quote)
	return quote, err
}

// RelevantStocks is similar to the peers endpoint, except this will return
// most active market symbols when peers are not available. If the symbols
// returned are not peers, the peers key will be false. This is not intended to
// represent a definitive or accurate list of peers, and is subject to change
// at any time.
func (c Client) RelevantStocks(stock string) (RelevantStocks, error) {
	rs := RelevantStocks{}
	endpoint := fmt.Sprintf("/stock/%s/relevant", url.PathEscape(stock))
	err := c.GetJSON(endpoint, &rs)
	return rs, err
}

// USExchanges returns an array of U.S. Exchanges.
func (c Client) USExchanges() ([]USExchange, error) {
	e := []USExchange{}
	endpoint := "/ref-data/market/us/exchanges"
	err := c.GetJSON(endpoint, &e)
	return e, err
}

// Symbols returns an array of symbols that IEX Cloud supports for API calls.
func (c Client) Symbols() ([]Symbol, error) {
	symbols := []Symbol{}
	endpoint := "/ref-data/symbols"
	err := c.GetJSON(endpoint, &symbols)
	return symbols, err
}

// IEXSymbols returns an array of symbols the Investors Exchange supports for
// trading. This list is updated daily as of 7:45 a.m. ET. Symbols may be added
// or removed by the Investors Exchange after the list was produced.
func (c Client) IEXSymbols() ([]TradedSymbol, error) {
	symbols := []TradedSymbol{}
	endpoint := "/ref-data/iex/symbols"
	err := c.GetJSON(endpoint, &symbols)
	return symbols, err
}

// ExchangeRate returns an end of day exchange rate of a given currency pair.
func (c Client) ExchangeRate(from, to string) (ExchangeRate, error) {
	r := ExchangeRate{}
	endpoint := fmt.Sprintf("/fx/rate/%s/%s",
		url.PathEscape(from),
		url.PathEscape(to))
	err := c.GetJSON(endpoint, &r)
	return r, err
}

// Last provides trade data for executions on IEX. It is a near real time,
// intraday API that provides IEX last sale price, size and time. Last is ideal
// for developers that need a lightweight stock quote.
func (c Client) Last(stock string) ([]Last, error) {
	// FIXME: Change so that multiple stock symbols can be handled.
	x := []Last{}
	endpoint := "/tops/last?symbols=" + url.PathEscape(stock)
	err := c.GetJSON(endpoint, &x)
	return x, err
}

// Status returns the IEX Cloud system status.
func (c Client) Status() (Status, error) {
	status := Status{}
	endpoint := "/status"
	err := c.GetJSONWithoutToken(endpoint, &status)
	return status, err
}
