// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/goinvest/iexcloud/v2/test/fakeiexcloud"
)

const testToken = "not-a-real-token"

// func TestAnnualBalanceSheets(t *testing.T) {
// 	cfg, err := readConfig("config_test.toml")
// 	if err != nil {
// 		log.Fatalf("Error reading config file: %s", err)
// 	}
// 	client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
// 	bs, err := client.AnnualBalanceSheets(context.Background(), "aapl", 4)
// 	if err != nil {
// 		log.Fatalf("Error getting balance sheets: %s", err)
// 	}
// 	assertString(t, "symbol", bs.Symbol, "AAPL")
// 	assertInt(t, "number of years", len(bs.Statements), 4)
// 	q1 := bs.Statements[0]
// 	assertString(t, "filing type", q1.FilingType, "10-K")
// 	assertInt(t, "fiscal quarter", q1.FiscalQuarter, 0)
// 	isPositiveInt(t, "fiscal year", q1.FiscalYear)
// 	assertString(t, "currency", q1.Currency, "USD")
// }

// func TestQuarterlyBalanceSheets(t *testing.T) {
// 	cfg, err := readConfig("config_test.toml")
// 	if err != nil {
// 		log.Fatalf("Error reading config file: %s", err)
// 	}
// 	client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
// 	bs, err := client.QuarterlyBalanceSheets(context.Background(), "aapl", 2)
// 	if err != nil {
// 		log.Fatalf("Error getting balance sheets: %s", err)
// 	}
// 	assertString(t, "symbol", bs.Symbol, "AAPL")
// 	assertInt(t, "number of quarters", len(bs.Statements), 2)
// 	q1 := bs.Statements[0]
// 	assertString(t, "filing type", q1.FilingType, "10-K")
// 	isPositiveInt(t, "fiscal quarter", q1.FiscalQuarter)
// 	isPositiveInt(t, "fiscal year", q1.FiscalYear)
// 	assertString(t, "currency", q1.Currency, "USD")
// }

// func TestBook(t *testing.T) {
// 	cfg, err := readConfig("config_test.toml")
// 	if err != nil {
// 		log.Fatalf("Error reading config file: %s", err)
// 	}
// 	client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
// 	got, err := client.Book(context.Background(), "aapl")
// 	if err != nil {
// 		log.Fatalf("Error getting book: %s", err)
// 	}
// 	assertString(t, "symbol", got.Quote.Symbol, "AAPL")
// 	assertString(t, "company name", got.Quote.CompanyName, "Apple Inc")
// 	assertScrambledString(t, "primary exchange", got.Quote.PrimaryExchange, "NASDAQ")
// 	assertString(t, "calculation price", got.Quote.CalculationPrice, "close")
// 	isPositiveFloat64(t, "open", got.Quote.Open)
// 	assertScrambledString(t, "open source", got.Quote.OpenSource, "official")
// 	isPositiveFloat64(t, "latest price", got.Quote.LatestPrice)
// }

// Returns the base address for the test server.
func baseAddress(s *httptest.Server) ClientOption {
	return WithBaseURL("http://" + s.Listener.Addr().String())
}

func TestHistoricalPrices(t *testing.T) {
	fakeIEX := fakeiexcloud.FakeIEXCloud{}
	s := httptest.NewServer(http.HandlerFunc(fakeIEX.Handle))
	defer s.Close()
	client := NewClient(testToken, baseAddress(s))

	testCases := []struct {
		name string

		// These parameters will be used in the request.
		requestSymbol    string
		requestTimeframe HistoricalTimeFrame

		// These configure the fake response.
		responseJSON       string
		responseHTTPStatus int

		// These set our expectations for the test result.
		wantRequestPath string
		wantPrices      []HistoricalDataPoint
		wantErr         bool
	}{
		{
			name:             "nominal",
			requestSymbol:    "aapl",
			requestTimeframe: "1m",
			responseJSON: `[
				{
					"close": 161.84,
					"high": 164.96,
					"low": 159.72,
					"open": 164.02,
					"symbol": "AAPL",
					"volume": 118023116,
					"id": "HISTORICAL_PRICES",
					"key": "AAPL",
					"subkey": "12345",
					"date": "2021-12-03",
					"uOpen": 164.02,
					"uClose": 161.84,
					"uHigh": 164.96,
					"uLow": 159.72,
					"uVolume": 118023116,
					"change": -1.9199999999999875,
					"changePercent": -0.0117,
					"label": "Dec 3, 21",
					"changeOverTime": 0.07577771869183732
				}
			]`,
			wantRequestPath: "/stock/aapl/chart/1m",
			wantPrices: []HistoricalDataPoint{{
				Close:          161.84,
				High:           164.96,
				Low:            159.72,
				Open:           164.02,
				Symbol:         "AAPL",
				Volume:         118023116,
				ID:             "HISTORICAL_PRICES",
				Key:            "AAPL",
				Subkey:         "12345",
				Date:           Date(time.Date(2021, 12, 3, 0, 0, 0, 0, time.UTC)),
				UOpen:          164.02,
				UClose:         161.84,
				UHigh:          164.96,
				ULow:           159.72,
				UVolume:        118023116,
				Change:         -1.9199999999999875,
				ChangePercent:  -0.0117,
				Label:          "Dec 3, 21",
				ChangeOverTime: 0.07577771869183732,
			}},
		},
		{
			name:             "invalid time frame",
			requestTimeframe: "asdf",
			responseJSON:     "",
			wantErr:          true,
		},
		{
			name:               "server error",
			requestTimeframe:   "1m",
			responseHTTPStatus: http.StatusInternalServerError,
			wantErr:            true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the fake response.
			fakeIEX.ResponseJSON = tc.responseJSON
			fakeIEX.ResponseHTTPStatus = tc.responseHTTPStatus

			// Run the fetch.
			histPrices, err := client.HistoricalPrices(context.TODO(), tc.requestSymbol, tc.requestTimeframe, nil)

			// Compare the response with our test expectations.
			if err != nil {
				if tc.wantErr {
					return // error was expected
				}
				t.Fatalf("%s: Error getting historical prices: %s", tc.name, err)
			}
			if tc.wantErr {
				t.Fatalf("%s: Got nil error, want error", tc.name)
			}

			if diff := deep.Equal(histPrices, tc.wantPrices); diff != nil {
				t.Fatalf("%s: Got unexpected values:\n%s", tc.name, diff)
			}

			if got, want := fakeIEX.LastURLReceived.Path, tc.wantRequestPath; got != want {
				t.Errorf("%s: Got %q, want %q", tc.name, got, want)
			}
		})
	}
}
