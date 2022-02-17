// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/karagog/testutil-go/fakehttpserver"
)

const testToken = "not-a-real-token"

// Returns the base address for the given test server.
func withBaseAddress(s *httptest.Server) ClientOption {
	return WithBaseURL("http://" + s.Listener.Addr().String())
}

func TestURL(t *testing.T) {
	baseAddr := "http://base"
	c := NewClient("", WithBaseURL(baseAddr))

	for _, tc := range []struct {
		name        string
		endpoint    string
		queryParams map[string]string

		wantURL string
		wantErr bool
	}{
		{
			name:     "nominal",
			endpoint: "/path",
			wantURL:  "http://base/path",
		},
		{
			name:        "nominal w/ query params",
			endpoint:    "/path",
			queryParams: map[string]string{"param1": "value1", "param2": "value2"},
			wantURL:     "http://base/path?param1=value1&param2=value2",
		},
		{
			name:        "query params baked into endpoint",
			endpoint:    "/path?param1=value1",
			queryParams: map[string]string{"param2": "value2"},
			wantURL:     "http://base/path?param1=value1&param2=value2",
		},
		{
			name:     "error",
			endpoint: "not a valid url",
			wantErr:  true,
		},
	} {
		u, err := c.url(tc.endpoint, tc.queryParams)
		if err != nil {
			if tc.wantErr {
				continue // error was expected
			}
			t.Fatalf("%s: %s", tc.name, err)
		}
		if tc.wantErr {
			t.Errorf("%s: Got nil error, want error", tc.name)
		}
		if got, want := u.String(), tc.wantURL; got != want {
			t.Fatalf("%s: Got %v, want %v", tc.name, got, want)
		}
	}
}

func TestBalanceSheets(t *testing.T) {
	fakeIEX := fakehttpserver.FakeHTTPServer{}
	s := httptest.NewServer(http.HandlerFunc(fakeIEX.Handle))
	defer s.Close()
	client := NewClient(testToken, withBaseAddress(s))

	const nominalBalanceSheetJSON = `{
		"symbol": "AAPL",
		"balancesheet": [
			{
				"reportDate": "2020-10-17",
				"filingType": "10-K",
				"fiscalDate": "2020-09-13",
				"fiscalQuarter": 4,
				"fiscalYear": 2010,
				"currency": "USD",
				"currentCash": 25913000000,
				"shortTermInvestments": null,
				"receivables": 23186000000,
				"inventory": 3956000000,
				"otherCurrentAssets": 12087000000,
				"currentAssets": 131339000000,
				"longTermInvestments": 170799000000,
				"propertyPlantEquipment": 41304000000,
				"goodwill": null,
				"intangibleAssets": null,
				"otherAssets": 22283000000,
				"totalAssets": 365725000000,
				"accountsPayable": 55888000000,
				"currentLongTermDebt": null,
				"otherCurrentLiabilities": null,
				"totalCurrentLiabilities": 116866000000,
				"longTermDebt": 93735000000,
				"otherLiabilities": null,
				"minorityInterest": 0,
				"totalLiabilities": 258578000000,
				"commonStock": 40201000000,
				"retainedEarnings": 70400000000,
				"treasuryStock": null,
				"capitalSurplus": null,
				"shareholderEquity": 107147000000,
				"netTangibleAssets": 107147000000,
				"id": "BALANCE_SHEET",
				"key": "AAPL",
				"subkey": "quarterly",
				"date": 1635273127391,
				"updated": 1635273127391
			}
		]
	}`

	var nominalBalanceSheets = BalanceSheets{
		Symbol: "AAPL",
		Statements: []BalanceSheet{
			{
				ReportDate:              Date(time.Date(2020, 10, 17, 0, 0, 0, 0, time.UTC)),
				FilingType:              "10-K",
				FiscalDate:              Date(time.Date(2020, 9, 13, 0, 0, 0, 0, time.UTC)),
				FiscalQuarter:           4,
				FiscalYear:              2010,
				Currency:                "USD",
				CurrentCash:             25913000000,
				Receivables:             23186000000,
				Inventory:               3956000000,
				OtherCurrentAssets:      12087000000,
				CurrentAssets:           131339000000,
				LongTermInvestments:     170799000000,
				PropertyPlantEquipment:  41304000000,
				OtherAssets:             22283000000,
				TotalAssets:             365725000000,
				AccountsPayable:         55888000000,
				TotalCurrentLiabilities: 116866000000,
				LongTermDebt:            93735000000,
				MinorityInterest:        0,
				TotalLiabilities:        258578000000,
				CommonStock:             40201000000,
				RetainedEarnings:        70400000000,
				ShareholderEquity:       107147000000,
				NetTangibleAssets:       107147000000,
			},
		},
	}

	testCases := []struct {
		name string

		// These parameters will be used in the request.
		requestSymbol string
		requestPeriod string // annual/quarter
		requestNumber int

		// These configure the fake response.
		responseJSON       string
		responseHTTPStatus int

		// These set our expectations for the test result.
		wantRequestPath   string
		wantQueryParams   map[string][]string
		wantBalanceSheets BalanceSheets
		wantErr           bool
	}{
		{
			name:              "nominal - annual",
			requestSymbol:     "aapl",
			requestPeriod:     "annual",
			requestNumber:     1,
			responseJSON:      nominalBalanceSheetJSON,
			wantRequestPath:   "/stock/aapl/balance-sheet/1",
			wantQueryParams:   map[string][]string{"token": []string{testToken}, "period": []string{"annual"}},
			wantBalanceSheets: nominalBalanceSheets,
		},
		{
			name:              "nominal - quarterly",
			requestSymbol:     "goog",
			requestPeriod:     "quarter",
			requestNumber:     2,
			responseJSON:      nominalBalanceSheetJSON,
			wantRequestPath:   "/stock/goog/balance-sheet/2",
			wantQueryParams:   map[string][]string{"token": []string{testToken}, "period": []string{"quarter"}},
			wantBalanceSheets: nominalBalanceSheets,
		},
	}

	for _, tc := range testCases {
		fakeIEX.ResponseJSON = tc.responseJSON
		fakeIEX.ResponseHTTPStatus = tc.responseHTTPStatus

		bs, err := client.BalanceSheets(context.TODO(), tc.requestSymbol, tc.requestPeriod, tc.requestNumber)
		if err != nil {
			if tc.wantErr {
				return // error was expected
			}
			t.Fatalf("%s: Error getting balance sheets: %s", tc.name, err)
		}
		if tc.wantErr {
			t.Fatalf("%s: Got nil error, want error", tc.name)
		}

		if diff := deep.Equal(bs, tc.wantBalanceSheets); diff != nil {
			t.Fatalf("%s: Got unexpected values:\n%s", tc.name, diff)
		}

		if got, want := fakeIEX.LastURLReceived.Path, tc.wantRequestPath; got != want {
			t.Errorf("%s: Got %q, want %q", tc.name, got, want)
		}
		if diff := deep.Equal(fakeIEX.LastURLReceived.Query(), url.Values(tc.wantQueryParams)); diff != nil {
			t.Fatalf("%s: Got unexpected values:\n%s", tc.name, diff)
		}
	}
}

func TestIncomeStatements(t *testing.T) {
	fakeIEX := fakehttpserver.FakeHTTPServer{}
	s := httptest.NewServer(http.HandlerFunc(fakeIEX.Handle))
	defer s.Close()
	client := NewClient(testToken, withBaseAddress(s))

	const nominalIncomeStatementsJSON = `{
		"symbol": "AAPL",
		"income": [
			{
				"reportDate": "2020-10-21",
				"filingType": "10-K",
				"fiscalDate": "2020-09-13",
				"fiscalQuarter": 4,
				"fiscalYear": 2010,
				"currency": "USD",
				"totalRevenue": 62681000000,
				"costOfRevenue": 39086000000,
				"grossProfit": 23595000000,
				"researchAndDevelopment": 3750000000,
				"sellingGeneralAndAdmin": 4216000000,
				"operatingExpense": 47052000000,
				"operatingIncome": 15629000000,
				"otherIncomeExpenseNet": 792000000,
				"ebit": 15629000000,
				"interestIncome": 868000000,
				"pretaxIncome": 16421000000,
				"incomeTax": 2296000000,
				"minorityInterest": 0,
				"netIncome": 14125000000,
				"netIncomeBasic": 14125000000
			}
		]
	}`

	var nominalIncomeStatements = IncomeStatements{
		Symbol: "AAPL",
		Statements: []IncomeStatement{
			{
				ReportDate:             Date(time.Date(2020, 10, 17, 0, 0, 0, 0, time.UTC)),
				FiscalDate:             Date(time.Date(2020, 9, 13, 0, 0, 0, 0, time.UTC)),
				Currency:               "USD",
				MinorityInterest:       0,
				TotalRevenue:           62681000000,
				CostOfRevenue:          39086000000,
				GrossProfit:            23595000000,
				ResearchAndDevelopment: 3750000000,
				SellingGeneralAndAdmin: 4216000000,
				OperatingExpense:       47052000000,
				OperatingIncome:        15629000000,
				OtherIncomeExpenseNet:  792000000,
				EBIT:                   15629000000,
				InterestIncome:         868000000,
				PretaxIncome:           16421000000,
				IncomeTax:              2296000000,
				NetIncome:              14125000000,
				NetIncomeBasic:         14125000000,
			},
		},
	}

	testCases := []struct {
		name string

		// These parameters will be used in the request.
		requestSymbol string
		requestPeriod string // annual/quarter
		requestNumber int

		// These configure the fake response.
		responseJSON       string
		responseHTTPStatus int

		// These set our expectations for the test result.
		wantRequestPath      string
		wantQueryParams      map[string][]string
		wantIncomeStatements IncomeStatements
		wantErr              bool
	}{
		{
			name:                 "nominal - annual",
			requestSymbol:        "aapl",
			requestPeriod:        "annual",
			requestNumber:        1,
			responseJSON:         nominalIncomeStatementsJSON,
			wantRequestPath:      "/stock/aapl/income/1",
			wantQueryParams:      map[string][]string{"token": {testToken}, "period": {"annual"}},
			wantIncomeStatements: nominalIncomeStatements,
		},
		{
			name:                 "nominal - quarterly",
			requestSymbol:        "goog",
			requestPeriod:        "quarter",
			requestNumber:        2,
			responseJSON:         nominalIncomeStatementsJSON,
			wantRequestPath:      "/stock/goog/income/2",
			wantQueryParams:      map[string][]string{"token": {testToken}, "period": {"quarter"}},
			wantIncomeStatements: nominalIncomeStatements,
		},
	}

	for _, tc := range testCases {
		fakeIEX.ResponseJSON = tc.responseJSON
		fakeIEX.ResponseHTTPStatus = tc.responseHTTPStatus

		incomeStatements, err := client.incomeStatements(context.TODO(), tc.requestSymbol, tc.requestPeriod, tc.requestNumber)
		if err != nil {
			if tc.wantErr {
				return // error was expected
			}
			t.Fatalf("%s: Error getting balance sheets: %s", tc.name, err)
		}
		if tc.wantErr {
			t.Fatalf("%s: Got nil error, want error", tc.name)
		}

		if diff := deep.Equal(incomeStatements, tc.wantIncomeStatements); diff != nil {
			t.Fatalf("%s: Got unexpected values:\n%s", tc.name, diff)
		}

		if got, want := fakeIEX.LastURLReceived.Path, tc.wantRequestPath; got != want {
			t.Errorf("%s: Got %q, want %q", tc.name, got, want)
		}
		if diff := deep.Equal(fakeIEX.LastURLReceived.Query(), url.Values(tc.wantQueryParams)); diff != nil {
			t.Fatalf("%s: Got unexpected values:\n%s", tc.name, diff)
		}
	}
}

func TestBook(t *testing.T) {
	fakeIEX := fakehttpserver.FakeHTTPServer{}
	s := httptest.NewServer(http.HandlerFunc(fakeIEX.Handle))
	defer s.Close()
	client := NewClient(testToken, withBaseAddress(s))

	testCases := []struct {
		name string

		// These parameters will be used in the request.
		requestSymbol string

		// These configure the fake response.
		responseJSON       string
		responseHTTPStatus int

		// These set our expectations for the test result.
		wantRequestPath string
		wantBook        Book
		wantErr         bool
	}{
		{
			name:          "nominal",
			requestSymbol: "aapl",
			responseJSON: `{
				"quote": { "symbol": "AAPL" },
				"bids": [{ "price": 120, "size": 10000, "timestamp": 1638820374000 }],
				"asks": [{ "price": 122, "size": 10030, "timestamp": 1638820374000 }],
				"trades": [{
					"price": 120,
					"size": 141432,
					"tradeId": 12,
					"isISO": true,
					"isOddLot": true,
					"isOutsideRegularHours": true,
					"isSinglePriceCross": true,
					"isTradeThroughExempt": true,
					"timestamp": 1638820374000
				}],
				"systemEvent": { "systemEvent": "foo", "timestamp": 1638820374000 }
			}`,
			wantRequestPath: "/stock/aapl/book",
			wantBook: Book{
				Quote: Quote{Symbol: "AAPL"},
				Bids: []BidAsk{{
					Price:     120,
					Size:      10000,
					Timestamp: EpochTime(time.Unix(1638820374, 0)),
				}},
				Asks: []BidAsk{{
					Price:     122,
					Size:      10030,
					Timestamp: EpochTime(time.Unix(1638820374, 0)),
				}},
				Trades: []Trade{{
					Price:                 120,
					Size:                  141432,
					TradeID:               12,
					IsISO:                 true,
					IsOddLot:              true,
					IsOutsideRegularHours: true,
					IsSinglePriceCross:    true,
					IsTradeThroughExempt:  true,
					Timestamp:             EpochTime(time.Unix(1638820374, 0)),
				}},
				SystemEvent: SystemEvent{
					SystemEvent: "foo",
					Timestamp:   EpochTime(time.Unix(1638820374, 0)),
				},
			},
		},
		{
			name:               "error",
			requestSymbol:      "aapl",
			responseHTTPStatus: http.StatusNotFound,
			wantRequestPath:    "/stock/aapl/book",
			wantErr:            true,
		},
	}

	for _, tc := range testCases {
		fakeIEX.ResponseJSON = tc.responseJSON
		fakeIEX.ResponseHTTPStatus = tc.responseHTTPStatus

		book, err := client.Book(context.TODO(), tc.requestSymbol)
		if err != nil {
			if tc.wantErr {
				return // error was expected
			}
			t.Fatalf("%s: Error getting book: %s", tc.name, err)
		}
		if tc.wantErr {
			t.Fatalf("%s: Got nil error, want error", tc.name)
		}

		if diff := deep.Equal(book, tc.wantBook); diff != nil {
			t.Fatalf("%s: Got unexpected values:\n%s", tc.name, diff)
		}

		if got, want := fakeIEX.LastURLReceived.Path, tc.wantRequestPath; got != want {
			t.Errorf("%s: Got %q, want %q", tc.name, got, want)
		}
	}
}

func TestHistoricalPrices(t *testing.T) {
	fakeIEX := fakehttpserver.FakeHTTPServer{}
	s := httptest.NewServer(http.HandlerFunc(fakeIEX.Handle))
	defer s.Close()
	client := NewClient(testToken, withBaseAddress(s))

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
					"minute": "9:30",
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
				Minute:         HourMinute(9*time.Hour + 30*time.Minute),
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
			// Sometimes the historical data returns a floating point number for some reason, which
			// can cause our JSON parsing to fail if we're not ready for it.
			name:             "corner case - floating point volume",
			requestSymbol:    "aapl",
			requestTimeframe: "1y",
			responseJSON:     `[{ "volume": 118023116.2342521 }]`,
			wantRequestPath:  "/stock/aapl/chart/1y",
			wantPrices:       []HistoricalDataPoint{{Volume: 118023116.2342521}},
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
