// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
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
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/google/go-querystring/query"
)

const apiURL = "https://cloud.iexapis.com/v1"

// Client models a client to consume the IEX Cloud API.
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Error represents an IEX API error
type Error struct {
	StatusCode int
	Message    string
}

// ClientOption applies an option to the client.
type ClientOption func(*Client)

// Error implements the error interface
func (e Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), e.Message)
}

// NewClient creates a client with the given authorization token.
func NewClient(token string, options ...ClientOption) *Client {
	client := &Client{
		token:      token,
		httpClient: &http.Client{Timeout: time.Second * 60},
	}

	// apply options
	for _, applyOption := range options {
		applyOption(client)
	}

	// set default values
	if client.baseURL == "" {
		client.baseURL = apiURL
	}

	return client
}

// WithHTTPClient sets the http.Client for a new IEX Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

// WithSecureHTTPClient sets a secure http.Client for a new IEX Client
func WithSecureHTTPClient() ClientOption {
	return func(client *Client) {
		client.httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			}}
	}
}

// WithBaseURL sets the baseURL for a new IEX Client
func WithBaseURL(baseURL string) ClientOption {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(ctx context.Context, endpoint string, v interface{}) error {
	u, err := c.url(endpoint, map[string]string{"token": c.token})
	if err != nil {
		return err
	}
	return c.FetchURLToJSON(ctx, u, v)
}

// GetJSONWithQueryParams gets the JSON data from the given endpoint with the query parameters attached.
func (c *Client) GetJSONWithQueryParams(ctx context.Context, endpoint string, queryParams map[string]string, v interface{}) error {
	queryParams["token"] = c.token
	u, err := c.url(endpoint, queryParams)
	if err != nil {
		return err
	}
	return c.FetchURLToJSON(ctx, u, v)
}

// Fetches JSON content from the given URL and unmarshals it into `v`.
func (c *Client) FetchURLToJSON(ctx context.Context, u *url.URL, v interface{}) error {
	data, err := c.getBytes(ctx, u.String())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// GetJSONWithoutToken gets the JSON data from the given endpoint without
// adding a token to the URL.
func (c *Client) GetJSONWithoutToken(ctx context.Context, endpoint string, v interface{}) error {
	u, err := c.url(endpoint, nil)
	if err != nil {
		return err
	}
	return c.FetchURLToJSON(ctx, u, v)
}

// GetBytes gets the data from the given endpoint.
func (c *Client) GetBytes(ctx context.Context, endpoint string) ([]byte, error) {
	u, err := c.url(endpoint, map[string]string{"token": c.token})
	if err != nil {
		return nil, err
	}
	return c.getBytes(ctx, u.String())
}

// GetFloat64 gets the number from the given endpoint.
func (c *Client) GetFloat64(ctx context.Context, endpoint string) (float64, error) {
	b, err := c.GetBytes(ctx, endpoint)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(string(b), 64)
}

func (c *Client) getBytes(ctx context.Context, address string) ([]byte, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return []byte{}, err
	}
	glog.V(1).Infof("Sending request to IEX Cloud: %v", req.URL.String())
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	// Even if GET didn't return an error, check the status code to make sure
	// everything was ok.
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		msg := ""

		if err == nil {
			msg = string(b)
		}

		return []byte{}, Error{StatusCode: resp.StatusCode, Message: msg}
	}
	return ioutil.ReadAll(resp.Body)
}

// Returns an URL object that points to the endpoint with optional query parameters.
func (c *Client) url(endpoint string, queryParams map[string]string) (*url.URL, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, err
	}

	if queryParams != nil {
		q := u.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}
	return u, nil
}

//////////////////////////////////////////////////////////////////////////////
//
// Data Points Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// AvailableDataPoints returns a list of the available data points for a given
// symbol and the weight of each data point.
func (c Client) AvailableDataPoints(ctx context.Context, symbol string) ([]DataPoint, error) {
	var dataPoints []DataPoint
	endpoint := fmt.Sprintf("/data-points/%s", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &dataPoints)
	return dataPoints, err
}

// DataPoint returns the plain text value for the requested data point key for
// the given symbol.
func (c Client) DataPoint(ctx context.Context, symbol, key string) ([]byte, error) {
	endpoint := fmt.Sprintf("/data-points/%s/%s", url.PathEscape(symbol), url.PathEscape(key))
	return c.GetBytes(ctx, endpoint)
}

// DataPointNumber returns the float64 for the requested data point key and the
// given symbol.
func (c Client) DataPointNumber(ctx context.Context, symbol, key string) (float64, error) {
	endpoint := fmt.Sprintf("/data-points/%s/%s", url.PathEscape(symbol), url.PathEscape(key))
	return c.GetFloat64(ctx, endpoint)
}

//////////////////////////////////////////////////////////////////////////////
//
// Account Endpoints
//
//////////////////////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////////////////////////////
//
// API System Metadata
//
//////////////////////////////////////////////////////////////////////////////

// Status returns the IEX Cloud system status.
func (c Client) Status(ctx context.Context) (Status, error) {
	status := Status{}
	endpoint := "/status"
	err := c.GetJSONWithoutToken(ctx, endpoint, &status)
	return status, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Stock / Equities Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// AnalystRecommendationsAndTargets provides current and historical consensus
// analyst recommendations and price targets.
func (c Client) AnalystRecommendationsAndTargets(ctx context.Context, symbol string) (CoreEstimate, error) {
	estimate := CoreEstimate{}
	endpoint := fmt.Sprintf("/time-series/CORE_ESTIMATES/%s", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &estimate)
	return estimate, err
}

// AnnualBalanceSheets returns the specified number of most recent annual
// balance sheets from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualBalanceSheets(ctx context.Context, symbol string, num int) (BalanceSheets, error) {
	return c.BalanceSheets(ctx, symbol, "annual", num)
}

// QuarterlyBalanceSheets returns the specified number of most recent quarterly
// balance sheets from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyBalanceSheets(ctx context.Context, symbol string, num int) (BalanceSheets, error) {
	return c.BalanceSheets(ctx, symbol, "quarter", num)
}

// BalanceSheets returns the specified number of most recent balance sheets
// with the given period (either "annual" or "quarter").
func (c Client) BalanceSheets(ctx context.Context, symbol, period string, num int) (BalanceSheets, error) {
	endpoint := fmt.Sprintf("/stock/%s/balance-sheet/%d", url.PathEscape(symbol), num)
	bs := BalanceSheets{}
	err := c.GetJSONWithQueryParams(ctx, endpoint, map[string]string{"period": period}, &bs)
	return bs, err
}

// Book returns the quote, bids, asks, and trades for a given stock symbol.
func (c Client) Book(ctx context.Context, symbol string) (Book, error) {
	book := Book{}
	endpoint := fmt.Sprintf("/stock/%s/book", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &book)
	return book, err
}

// DelayedQuote returns the 15 minute delayed market quote from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) DelayedQuote(ctx context.Context, symbol string) (DelayedQuote, error) {
	dq := DelayedQuote{}
	endpoint := fmt.Sprintf("/stock/%s/delayed-quote", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &dq)
	return dq, err
}

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
	endpoint := fmt.Sprintf("/stock/%s/chart/date/%s?chartByDay=true",
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
	endpoint := fmt.Sprintf("/stock/%s/chart/date/%s",
		url.PathEscape(symbol), day.Format("20060102"))
	endpoint, err := c.intradayHistoricalEndpointWithOpts(endpoint, options, false)
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

// IntradayPrices returns the aggregated intraday prices in one minute buckets.
func (c Client) IntradayPrices(ctx context.Context, symbol string) ([]IntradayPrice, error) {
	ip := []IntradayPrice{}
	endpoint := fmt.Sprintf("/stock/%s/intraday-prices", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &ip)
	return ip, err
}

// IntradayPricesWithOpts returns the aggregated intraday prices in one minute buckets for the given options.
func (c Client) IntradayPricesWithOpts(ctx context.Context, symbol string, options *IntradayOptions) ([]IntradayPrice, error) {
	ip := []IntradayPrice{}
	endpoint := fmt.Sprintf("/stock/%s/intraday-prices", url.PathEscape(symbol))
	endpoint, err := c.intradayEndpointWithOpts(endpoint, options, false)
	if err != nil {
		return ip, err
	}
	err = c.GetJSON(ctx, endpoint, &ip)
	return ip, err
}

func (c Client) intradayEndpointWithOpts(endpoint string, opts *IntradayOptions, existingParams bool) (string, error) {
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

// LargestTrades returns the 15 minute delayed, last sale eligible trade from
// the IEX Cloud endpoint for the given stock symbol.
func (c Client) LargestTrades(ctx context.Context, symbol string) ([]LargestTrade, error) {
	lt := []LargestTrade{}
	endpoint := fmt.Sprintf("/stock/%s/largest-trades", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &lt)
	return lt, err
}

// OHLC returns the OHLC data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) OHLC(ctx context.Context, symbol string) (OHLC, error) {
	ohlc := OHLC{}
	endpoint := fmt.Sprintf("/stock/%s/ohlc", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &ohlc)
	return ohlc, err
}

// PreviousDay returns the previous day adjusted price data from the IEX Cloud
// endpoint for the given stock symbol.
func (c Client) PreviousDay(ctx context.Context, symbol string) (PreviousDay, error) {
	pd := PreviousDay{}
	endpoint := fmt.Sprintf("/stock/%s/previous", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &pd)
	return pd, err
}

// Price returns the current stock price for the given stock symbol.
func (c Client) Price(ctx context.Context, symbol string) (float64, error) {
	endpoint := fmt.Sprintf("/stock/%s/price", url.PathEscape(symbol))
	return c.GetFloat64(ctx, endpoint)
}

// Quote returns the quote data from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) Quote(ctx context.Context, symbol string) (Quote, error) {
	r := Quote{}
	endpoint := fmt.Sprintf("/stock/%s/quote", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// BatchQuote returns the quote data for up to 100 stock symbols.
func (c Client) BatchQuote(ctx context.Context, symbols []string) (map[string]Quote, error) {
	r := map[string]struct {
		Quote Quote
	}{}
	endpoint := fmt.Sprintf("/stock/market/batch?symbols=%s&types=quote", url.PathEscape(strings.Join(symbols, ",")))
	err := c.GetJSON(ctx, endpoint, &r)
	quotes := make(map[string]Quote, len(r))
	for symbol, quote := range r {
		quotes[symbol] = quote.Quote
	}
	return quotes, err
}

// BatchPrevious returns the previous day price for up to 100 stock symbols.
func (c Client) BatchPrevious(ctx context.Context, symbols []string) (map[string]PreviousDay, error) {
	r := map[string]struct {
		Previous PreviousDay
	}{}
	endpoint := fmt.Sprintf("/stock/market/batch?symbols=%s&types=previous", url.PathEscape(strings.Join(symbols, ",")))
	err := c.GetJSON(ctx, endpoint, &r)
	previousday := make(map[string]PreviousDay, len(r))
	for symbol, quote := range r {
		previousday[symbol] = quote.Previous
	}
	return previousday, err
}

// VolumeByVenue returns the 15 minute delayed and 30 day average consolidated
// volume percentage of a stock by market. This will return 13 values sorted in
// ascending order by current day trading volume percentage.
func (c Client) VolumeByVenue(ctx context.Context, symbol string) ([]VenueVolume, error) {
	r := []VenueVolume{}
	endpoint := fmt.Sprintf("/stock/%s/volume-by-venue", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Stock Profiles Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// Company returns the copmany data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Company(ctx context.Context, symbol string) (Company, error) {
	company := Company{}
	endpoint := fmt.Sprintf("/stock/%s/company", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &company)
	return company, err
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

// Logo returns the logo data from the IEX Cloud endpoint for the given
// stock symbol.
func (c Client) Logo(ctx context.Context, symbol string) (Logo, error) {
	logo := Logo{}
	endpoint := fmt.Sprintf("/stock/%s/logo", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &logo)
	return logo, err
}

// Peers returns a slice of peer stock symbols for the given stock symbol.
func (c Client) Peers(ctx context.Context, symbol string) ([]string, error) {
	peers := []string{}
	endpoint := fmt.Sprintf("/stock/%s/peers", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &peers)
	return peers, err
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

//////////////////////////////////////////////////////////////////////////////
//
// Stock Fundamentals Endpoints
//
//////////////////////////////////////////////////////////////////////////////

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

// Dividends returns the dividends from the IEX Cloud endpoint for the given
// stock symbol and the given date range.
func (c Client) Dividends(ctx context.Context, symbol string, r PathRange) ([]Dividend, error) {
	dividends := []Dividend{}
	endpoint := fmt.Sprintf("/stock/%s/dividends/%s",
		url.PathEscape(symbol), PathRangeJSON[r])
	err := c.GetJSON(ctx, endpoint, &dividends)
	return dividends, err
}

// AnnualFinancials returns the specified number of most recent annual
// financials for the given stock symbol. This endpoint is carried over from
// the IEX 1.0 API and may be deprecated in the future.
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

// QuarterlyFinancialsAsReported returns the specified number of most recent
// quarterly 10-Q filings from the IEX Cloud endpoint for the given stock
// symbol.
func (c Client) QuarterlyFinancialsAsReported(ctx context.Context, symbol string, num int) (FinancialsAsReported, error) {
	endpoint := fmt.Sprintf("/time-series/reported_financials/%s/10-Q?limit=%d",
		url.PathEscape(symbol), num)
	return c.financialsAsReported(ctx, endpoint)
}

func (c Client) financialsAsReported(ctx context.Context, endpoint string) (FinancialsAsReported, error) {
	f := FinancialsAsReported{}
	err := c.GetJSON(ctx, endpoint, &f)
	return f, err
}

// AnnualIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) AnnualIncomeStatements(ctx context.Context, symbol string, num int) (IncomeStatements, error) {
	return c.incomeStatements(ctx, symbol, "annual", num)
}

// QuarterlyIncomeStatements returns the specified number of most recent annual
// income statements from the IEX Cloud endpoint for the given stock symbol.
func (c Client) QuarterlyIncomeStatements(ctx context.Context, symbol string, num int) (IncomeStatements, error) {
	return c.incomeStatements(ctx, symbol, "quarter", num)
}

// incomeStatements returns the specified number of most recent
// income statements from the IEX Cloud endpoint for the given stock symbol and period.
func (c Client) incomeStatements(ctx context.Context, symbol string, period string, num int) (IncomeStatements, error) {
	endpoint := fmt.Sprintf("/stock/%s/income/%d",
		url.PathEscape(symbol), num)
	is := IncomeStatements{}
	err := c.GetJSONWithQueryParams(ctx, endpoint, map[string]string{"period": period}, &is)
	return is, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Stock Research Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// AdvancedStats returns the everything in key stats plus additional advanced
// stats such as EBITDA, ratios, key financial data, and more.
func (c Client) AdvancedStats(ctx context.Context, symbol string) (AdvancedStats, error) {
	stats := AdvancedStats{}
	endpoint := fmt.Sprintf("/stock/%s/advanced-stats", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &stats)
	return stats, err
}

// RecommendationTrends provides a list of recommendations with the start and
// end date for each rating. Keep to not break API.
func (c Client) RecommendationTrends(ctx context.Context, symbol string) ([]Recommendation, error) {
	return c.AnalystRecommendations(ctx, symbol)
}

// FundOwnership returns the ten top holders of the given stock.
func (c Client) FundOwnership(ctx context.Context, symbol string) ([]FundOwner, error) {
	r := []FundOwner{}
	endpoint := fmt.Sprintf("/stock/%s/fund-ownership", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// InstitutionalOwnership returns the top 10 holders with the most recent
// information.
func (c Client) InstitutionalOwnership(ctx context.Context, symbol string) ([]InstitutionalOwner, error) {
	r := []InstitutionalOwner{}
	endpoint := fmt.Sprintf("/stock/%s/institutional-ownership", url.PathEscape(symbol))
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

//////////////////////////////////////////////////////////////////////////////
//
// Corporate Actions Endpoints
//
//////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////
//
// Market Info Endpoints
//
//////////////////////////////////////////////////////////////////////////////

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

// EarningsToday returns the earnings that will be reported today before the
// open and after the market closes.
func (c Client) EarningsToday(ctx context.Context) (EarningsToday, error) {
	e := EarningsToday{}
	endpoint := "/stock/market/today-earnings"
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// IPOsToday returns the IPOs that are scheduled to occur today.
func (c Client) IPOsToday(ctx context.Context) (IPOCalendar, error) {
	ic := IPOCalendar{}
	endpoint := "/stock/market/today-ipos"
	err := c.GetJSON(ctx, endpoint, &ic)
	return ic, err
}

// MostActive returns a list of quotes for the top 10 most active stocks from
// the IEX Cloud endpoint updated intraday, 15 minute delayed.
func (c Client) MostActive(ctx context.Context, limit int) ([]Quote, error) {
	return c.list(ctx, "mostactive", limit)
}

// Gainers returns a list of quotes for the top 10 stock gainers from
// the IEX Cloud endpoint updated intraday, 15 minute delayed.
func (c Client) Gainers(ctx context.Context, limit int) ([]Quote, error) {
	return c.list(ctx, "gainers", limit)
}

// Losers returns a list of quotes for the top 10 stock losers from
// the IEX Cloud endpoint updated intraday, 15 minute delayed.
func (c Client) Losers(ctx context.Context, limit int) ([]Quote, error) {
	return c.list(ctx, "losers", limit)
}

// IEXVolume returns a list of quotes for the top 10 IEX stocks by volume from
// the IEX Cloud endpoint updated intraday, 15 minute delayed.
func (c Client) IEXVolume(ctx context.Context, limit int) ([]Quote, error) {
	return c.list(ctx, "iexvolume", limit)
}

// IEXPercent returns a list of quotes for the top 10 IEX stocks by percent
// from the IEX Cloud endpoint updated intraday, 15 minute delayed.
func (c Client) IEXPercent(ctx context.Context, limit int) ([]Quote, error) {
	return c.list(ctx, "iexpercent", limit)
}

// InFocus returns a list of quotes for the top 10 in focus stocks from the IEX
// Cloud endpoint updated intraday, 15 minute delayed.
func (c Client) InFocus(ctx context.Context, limit int) ([]Quote, error) {
	return c.list(ctx, "infocus", limit)
}

func (c Client) list(ctx context.Context, list string, limit int) ([]Quote, error) {
	q := []Quote{}
	endpoint := "/stock/market/list/" + list
	params := make(map[string]string)
	if limit > 0 {
		params["listLimit"] = fmt.Sprintf("%d", limit)
	}
	err := c.GetJSONWithQueryParams(ctx, endpoint, params, &q)
	return q, err
}

// Markets returns real time traded volume on U.S. markets. This may be
// deprecated in the future. Use MarketVolume instead.
func (c Client) Markets(ctx context.Context) ([]Market, error) {
	return c.MarketVolume(ctx)
}

// MarketVolume returns the real time traded volume on U.S. markets.
func (c Client) MarketVolume(ctx context.Context) ([]Market, error) {
	m := []Market{}
	endpoint := "/stock/market/volume"
	err := c.GetJSON(ctx, endpoint, &m)
	return m, err
}

// SectorPerformance returns the performance of each sector for the current
// trading day. Performance is based on each sector ETF.
func (c Client) SectorPerformance(ctx context.Context) ([]SectorPerformance, error) {
	r := []SectorPerformance{}
	endpoint := "/stock/market/sector-performance"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// UpcomingEvents returns all upcoming events for a given symbol.  If an empty string is passed in for the symbol,
// data for the entire market, including IPOs, is returned.  If fullUpcomingEarnings is set to true, full estimates
// objects are returned; otherwise, earnings will only return Symbol and ReportDate.
func (c Client) UpcomingEvents(ctx context.Context, symbol string, fullUpcomingEarnings bool) (UpcomingEvents, error) {
	if symbol == "" {
		symbol = "market"
	}

	fue := ""

	if fullUpcomingEarnings {
		fue = "?fullUpcomingEarnings=true"
	}

	e := UpcomingEvents{}
	endpoint := fmt.Sprintf("/stock/%v/upcoming-events%v", symbol, fue)
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// UpcomingEarnings returns all upcoming earnings for a given symbol.  If an empty string is passed in for the symbol,
// data for the entire market is returned.  If fullUpcomingEarnings is set to true, full estimates
// objects are returned; otherwise, earnings will only return Symbol and ReportDate.
func (c Client) UpcomingEarnings(ctx context.Context, symbol string, fullUpcomingEarnings bool) ([]UpcomingEarning, error) {
	if symbol == "" {
		symbol = "market"
	}

	fue := ""

	if fullUpcomingEarnings {
		fue = "?fullUpcomingEarnings=true"
	}

	e := []UpcomingEarning{}
	endpoint := fmt.Sprintf("/stock/%v/upcoming-earnings%v", symbol, fue)
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// UpcomingDividends returns all upcoming dividends for a given symbol.  If an empty string is passed in for the symbol,
// data for the entire market is returned.
func (c Client) UpcomingDividends(ctx context.Context, symbol string) ([]Dividend, error) {
	if symbol == "" {
		symbol = "market"
	}

	e := []Dividend{}
	endpoint := fmt.Sprintf("/stock/%v/upcoming-dividends", symbol)
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// UpcomingSplits returns all upcoming splits for a given symbol.  If an empty string is passed in for the symbol,
// data for the entire market is returned.
func (c Client) UpcomingSplits(ctx context.Context, symbol string) ([]Split, error) {
	if symbol == "" {
		symbol = "market"
	}

	e := []Split{}
	endpoint := fmt.Sprintf("/stock/%v/upcoming-splits", symbol)
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

// UpcomingIPOs returns all upcoming IPOs for the entire market.
func (c Client) UpcomingIPOs(ctx context.Context) (IPOCalendar, error) {
	e := IPOCalendar{}
	endpoint := "/stock/market/upcoming-ipos"
	err := c.GetJSON(ctx, endpoint, &e)
	return e, err
}

//////////////////////////////////////////////////////////////////////////////
//
// News Endpoints
//
//////////////////////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////////////////////////////
//
// Cryptocurrency Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// Crypto provides a quote for a given cryptocurrency symbol.
func (c Client) Crypto(ctx context.Context, symbol string) (CryptoQuote, error) {
	return c.CryptoQuote(ctx, symbol)
}

// CryptoQuote provides a quote for a given cryptocurrency symbol.
func (c Client) CryptoQuote(ctx context.Context, symbol string) (CryptoQuote, error) {
	r := CryptoQuote{}
	endpoint := fmt.Sprintf("/crypto/%s/quote", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// CryptoPrice returns the price for a given cryptocurrency symbol.
func (c Client) CryptoPrice(ctx context.Context, symbol string) (Price, error) {
	r := Price{}
	endpoint := fmt.Sprintf("/crypto/%s/price", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// CryptoBooks returns a current snapshot of the book for a specified cryptocurrency
func (c Client) CryptoBooks(ctx context.Context, symbol string) (Books, error) {
	r := Books{}
	endpoint := fmt.Sprintf("/crypto/%s/book", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Forex / Currencies Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// CurrencyRates returns real-time foreign currency exchange rates data
// updated every 250 milliseconds.
func (c Client) CurrencyRates(ctx context.Context, symbols []string) ([]CurrencyRate, error) {
	r := []CurrencyRate{}
	endpoint := fmt.Sprintf("/fx/latest?symbols=%s", url.PathEscape(strings.Join(symbols, ",")))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// ExchangeRate returns an end of day exchange rate of a given currency pair.
//
// Deprecated: This endpoint does no longer exist.
// See https://www.iexcloud.io/docs/api/#forex-currencies
func (c Client) ExchangeRate(ctx context.Context, from, to string) (ExchangeRate, error) {
	r := ExchangeRate{}
	endpoint := fmt.Sprintf("/fx/rate/%s/%s",
		url.PathEscape(from),
		url.PathEscape(to))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Options Endpoints
//
//////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////
//
// Social Sentiment Endpoints
//
//////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////
//
// CEO Compensation Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// CEOCompensation provides CEO compensation for the given stock symbol.
func (c Client) CEOCompensation(ctx context.Context, symbol string) (CEOCompensation, error) {
	r := CEOCompensation{}
	endpoint := fmt.Sprintf("/stock/%s/ceo-compensation", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Treasuries Endpoints
//
//////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////
//
// Commodities Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// CommodityType indicates the type of commodity.
type CommodityType string

// Available commodities.
const (
	WestTexasOil       CommodityType = "DCOILWTICO"
	BrentEuropeOil     CommodityType = "DCOILBRENTEU"
	HenryHubNG         CommodityType = "DHHNGSP"
	NYHeatingOil       CommodityType = "DHOILNYH"
	GulfCoastJetFuel   CommodityType = "DJFUELUSGULF"
	USDiesel           CommodityType = "GASDESW"
	USRegularGas       CommodityType = "GASREGCOVW"
	USMidgradeGas      CommodityType = "GASMIDCOVW"
	USPremiumGas       CommodityType = "GASPRMCOVW"
	MontBelvieuPropane CommodityType = "DPROPANEMBTX"
)

var commodityDescriptions = map[CommodityType]string{
	WestTexasOil:       "Crude Oil West Texas Intermediate ($USD/barrel)",
	BrentEuropeOil:     "Crude Oil Brent Europe ($USD/barrel)",
	HenryHubNG:         "Henry Hub Natural Gas Spot Price ($USD/million BTU)",
	NYHeatingOil:       "No. 2 Heating Oil New York Harbor ($USD/gallon)",
	GulfCoastJetFuel:   "Kerosene Type Jet Fuel US Gulf Coast ($USD/gallon)",
	USDiesel:           "US Diesel ($USD/gallon)",
	USRegularGas:       "US Regular Conventional Gas ($USD/gallon)",
	USMidgradeGas:      "US Midgrade Conventional Gas ($USD/gallon)",
	USPremiumGas:       "US Premium Conventional Gas ($USD/gallon)",
	MontBelvieuPropane: "Mont Belvieu Texas Propane ($USD/gallon)",
}

// String provides the Stringer interface for CommodityType.
func (ct CommodityType) String() string {
	return commodityDescriptions[ct]
}

// CommodityPrice returns the price for the given commodity not seasonally
// adjusted.
func (c Client) CommodityPrice(ctx context.Context, ct CommodityType) (float64, error) {
	// By using an explicit type conversion to string we get the commodity symbol
	// instead of the description, which we would get if we utilized the Stringer
	// interface.
	return c.DataPointNumber(ctx, "market", string(ct))
}

//////////////////////////////////////////////////////////////////////////////
//
// Economic Data Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// CDRateType indicates the type of CD Rate.
type CDRateType string

// Available CD Rates.
const (
	NonJumboCD CDRateType = "MMNRNJ"
	JumboCD    CDRateType = "MMNRJD"
)

var cdRateDescriptions = map[CDRateType]string{
	NonJumboCD: "CD Rate Non-Jumbo less than $100,000 money market",
	JumboCD:    "CD Rate Jumbo more than $100,000 money market",
}

// String provides the Stringer interface for CDRateType.
func (cd CDRateType) String() string {
	return cdRateDescriptions[cd]
}

// CDRate returns the price for the given commodity not seasonally
// adjusted.
func (c Client) CDRate(ctx context.Context, cd CDRateType) (float64, error) {
	// By using an explicit type conversion to string we get the CD Rate symbol
	// instead of the description, which we would get if we utilized the Stringer
	// interface.
	return c.DataPointNumber(ctx, "market", string(cd))
}

// CPI returns the consumer price index for all urban consumers.
func (c Client) CPI(ctx context.Context) (float64, error) {
	return c.DataPointNumber(ctx, "market", "CPIAUCSL")
}

// CreditCardInterestRate returns the commercial bank credit card interest
// rate.
func (c Client) CreditCardInterestRate(ctx context.Context) (float64, error) {
	return c.DataPointNumber(ctx, "market", "TERMCBCCALLNS")
}

// FederalFundsRate returns the effective federal funds rate.
func (c Client) FederalFundsRate(ctx context.Context) (float64, error) {
	return c.DataPointNumber(ctx, "market", "FEDFUNDS")
}

//////////////////////////////////////////////////////////////////////////////
//
// Reference Data Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// CryptoSymbols returns a list of cryptocurrencies that are supported by IEX
// Cloud.
func (c Client) CryptoSymbols(ctx context.Context) ([]CryptoSymbol, error) {
	r := []CryptoSymbol{}
	endpoint := "/ref-data/crypto/symbols"
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

// IEXSymbols returns an array of symbols the Investors Exchange supports for
// trading. This list is updated daily as of 7:45 a.m. ET. Symbols may be added
// or removed by the Investors Exchange after the list was produced.
func (c Client) IEXSymbols(ctx context.Context) ([]TradedSymbol, error) {
	symbols := []TradedSymbol{}
	endpoint := "/ref-data/iex/symbols"
	err := c.GetJSON(ctx, endpoint, &symbols)
	return symbols, err
}

// MutualFundSymbols returns an array of mutual funds that IEX Cloud supports
// for API calls.
func (c Client) MutualFundSymbols(ctx context.Context) ([]Symbol, error) {
	r := []Symbol{}
	endpoint := "/ref-data/mutual-funds/symbols"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// OptionsSymbols returns a map keyed by symbol with the value of each symbol
// being an slice of available contract dates
func (c Client) OptionsSymbols(ctx context.Context) (map[string][]string, error) {
	r := map[string][]string{}
	endpoint := "/ref-data/options/symbols"
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

// Sectors returns an array of all sectors
func (c Client) Sectors(ctx context.Context) ([]Sector, error) {
	r := []Sector{}
	endpoint := "/ref-data/sectors"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// Symbols returns an array of symbols that IEX Cloud supports for API calls.
func (c Client) Symbols(ctx context.Context) ([]Symbol, error) {
	symbols := []Symbol{}
	endpoint := "/ref-data/symbols"
	err := c.GetJSON(ctx, endpoint, &symbols)
	return symbols, err
}

// SymbolsByExchange returns an array of symbols from the defined market that IEX Cloud supports for API calls.
func (c Client) SymbolsByExchange(ctx context.Context, exchange string) ([]Symbol, error) {
	symbols := []Symbol{}
	endpoint := "/ref-data/exchange/" + exchange + "/symbols"
	err := c.GetJSON(ctx, endpoint, &symbols)
	return symbols, err
}

// SymbolsByRegion returns an array of symbols from the defined region that IEX Cloud supports for API calls.
func (c Client) SymbolsByRegion(ctx context.Context, region string) ([]Symbol, error) {
	symbols := []Symbol{}
	endpoint := "/ref-data/region/" + region + "/symbols"
	err := c.GetJSON(ctx, endpoint, &symbols)
	return symbols, err
}

// Search returns an array of search results for the given symbol fragment.
func (c Client) Search(ctx context.Context, fragment string) ([]SearchResult, error) {
	searchResult := []SearchResult{}
	endpoint := fmt.Sprintf("/search/%s", fragment)
	err := c.GetJSON(ctx, endpoint, &searchResult)
	return searchResult, err
}

// Tags returns an array of tags.  Tags can
// be found for each on each company.
func (c Client) Tags(ctx context.Context) ([]Tag, error) {
	r := []Tag{}
	endpoint := "/ref-data/tags"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
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
	var r []TradeHolidayDate
	endpoint := "/ref-data/us/dates/trade/next/1"
	if err := c.GetJSON(ctx, endpoint, &r); err != nil {
		return TradeHolidayDate{}, err
	}
	return r[0], nil
}

// NextTradingDays returns the dates of the next trading days for the given
// number of days.
func (c Client) NextTradingDays(ctx context.Context, numDays int) ([]TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := fmt.Sprintf("/ref-data/us/dates/trade/next/%d", numDays)
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// NextHoliday returns the date of the next holiday.
func (c Client) NextHoliday(ctx context.Context) (TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/holiday/next/1"
	if err := c.GetJSON(ctx, endpoint, &r); err != nil {
		return TradeHolidayDate{}, err
	}
	return r[0], nil
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
	r := []TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/trade/last/1"
	err := c.GetJSON(ctx, endpoint, &r)
	return r[0], err
}

// PreviousHoliday returns the date of the previous holiday.
func (c Client) PreviousHoliday(ctx context.Context) (TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := "/ref-data/us/dates/holiday/last/1"
	err := c.GetJSON(ctx, endpoint, &r)
	return r[0], err
}

// Holidays returns the last or next dates of holidays, for the
// given number of days, from the given start date.
func (c Client) Holidays(ctx context.Context, dir string, last int, startDate time.Time) ([]TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := fmt.Sprintf("/ref-data/us/dates/holiday/%s/%d/%s", dir, last, startDate.Format("20060102"))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// TradingDays returns the last or next dates of trading days, for the
// given number of days, from the given start date.
func (c Client) TradingDays(ctx context.Context, dir string, last int, startDate time.Time) ([]TradeHolidayDate, error) {
	r := []TradeHolidayDate{}
	endpoint := fmt.Sprintf("/ref-data/us/dates/trade/%s/%d/%s", dir, last, startDate.Format("20060102"))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// ISINMapping convert ISIN to IEX Cloud symbols.
func (c Client) ISINMapping(ctx context.Context, symbol string) ([]SymbolDetails, error) {
	sd := []SymbolDetails{}
	endpoint := fmt.Sprintf("/ref-data/isin?isin=%s", symbol)
	err := c.GetJSON(ctx, endpoint, &sd)
	return sd, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Investors Exchange Data Endpoints
//
//////////////////////////////////////////////////////////////////////////////

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

// StatsIntraday retrieves the intraday stats on IEX.
func (c Client) StatsIntraday(ctx context.Context) (IntradayStats, error) {
	r := IntradayStats{}
	endpoint := "/stats/intraday"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// IntradayStats retrieves the intraday stats on IEX. Deprecated. Use
// StatsIntraday instead.
func (c Client) IntradayStats(ctx context.Context, symbol string) (IntradayStats, error) {
	// FIXME(mdr): symbol isn't used, so this is a bad method. Need to delete.
	r := IntradayStats{}
	endpoint := "/stats/intraday"
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

//////////////////////////////////////////////////////////////////////////////
//
// Premium Data Endpoints
//
//////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////
//
// Refinitiv Endpoints
//
//////////////////////////////////////////////////////////////////////////////

// AnalystRecommendations pulls data from the last four months using premium
// data from Refinitiv.
func (c Client) AnalystRecommendations(ctx context.Context, symbol string) ([]Recommendation, error) {
	r := []Recommendation{}
	endpoint := fmt.Sprintf("/stock/%s/recommendation-trends", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &r)
	return r, err
}

// Earnings returns the specified number of most recent earnings data from the
// IEX Cloud endpoint for the given stock symbol.
func (c Client) Earnings(ctx context.Context, symbol string, num int) (Earnings, error) {
	earnings := Earnings{}
	endpoint := fmt.Sprintf("/stock/%s/earnings/%d", url.PathEscape(symbol), num)
	err := c.GetJSON(ctx, endpoint, &earnings)
	return earnings, err
}

// Estimates returns the latest consensue estimates for the next fiscal period.
func (c Client) Estimates(ctx context.Context, symbol string, num int) (Estimates, error) {
	estimates := Estimates{}
	endpoint := fmt.Sprintf("/stock/%s/estimates/%d", url.PathEscape(symbol), num)
	err := c.GetJSON(ctx, endpoint, &estimates)
	return estimates, err
}

// PriceTarget returns the latest average, high, and low analyst price target
// for a given stock symbol.
func (c Client) PriceTarget(ctx context.Context, symbol string) (PriceTarget, error) {
	pt := PriceTarget{}
	endpoint := fmt.Sprintf("/stock/%s/price-target", url.PathEscape(symbol))
	err := c.GetJSON(ctx, endpoint, &pt)
	return pt, err
}
