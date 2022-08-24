// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

//go:build integration
// +build integration

package iex_test

import (
	"context"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	iex "github.com/goinvest/iexcloud/v2"
)

// Config contains the configuration information needed to program and test the
// adapaters.
type Config struct {
	Token   string
	BaseURL string
}

// ReadConfig will read the TOML config file.
func readConfig(configFile string) (Config, error) {

	var cfg Config

	// Read config file
	f, err := os.Open(configFile)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return cfg, err
	}
	err = toml.Unmarshal(buf, &cfg)
	return cfg, err
}

func TestIntegrationTests(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(
		cfg.Token,
		iex.WithBaseURL(cfg.BaseURL),
		iex.WithRateLimiter(time.Second, 10),
	)
	t.Run("Annual Balance Sheet", testIntegrationAnnualBalanceSheets(client))
}

func testIntegrationAnnualBalanceSheets(client *iex.Client) func(*testing.T) {
	return func(t *testing.T) {
		bs, err := client.AnnualBalanceSheets(context.Background(), "aapl", 4)
		if err != nil {
			log.Fatalf("Error getting annual balance sheets: %s", err)
		}
		assertString(t, "symbol", bs.Symbol, "AAPL")
		assertInt(t, "number of years", len(bs.Statements), 4)
		y1 := bs.Statements[0]
		assertString(t, "filing type", y1.FilingType, "10-K")
		assertInt(t, "fiscal quarter", y1.FiscalQuarter, 0)
		isPositiveInt(t, "fiscal year", y1.FiscalYear)
		assertString(t, "currency", y1.Currency, "USD")
	}
}

func testIntegrationQuarterlyBalanceSheets(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(
		cfg.Token,
		iex.WithBaseURL(cfg.BaseURL),
		iex.WithRateLimiter(time.Second, 10),
	)
	bs, err := client.QuarterlyBalanceSheets(context.Background(), "aapl", 4)
	if err != nil {
		log.Fatalf("Error getting balance sheets: %s", err)
	}
	assertString(t, "symbol", bs.Symbol, "AAPL")
	assertInt(t, "number of quarters", len(bs.Statements), 4)
	q1 := bs.Statements[0]
	assertString(t, "filing type", q1.FilingType, "10-Q")
	isPositiveInt(t, "fiscal quarter", q1.FiscalQuarter)
	isPositiveInt(t, "fiscal year", q1.FiscalYear)
	assertString(t, "currency", q1.Currency, "USD")
}

func testIntegrationAnnualIncomeStatements(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(
		cfg.Token,
		iex.WithRateLimiter(time.Second, 10),
		iex.WithBaseURL(cfg.BaseURL),
	)
	is, err := client.AnnualIncomeStatements(context.Background(), "f", 4)
	if err != nil {
		log.Fatalf("Error getting annual income statements: %s", err)
	}
	assertString(t, "symbol", is.Symbol, "F")
	if len(is.Statements) != 4 {
		t.Errorf("\ngot = %d %s\nwant = %d", len(is.Statements), "number of years", 4)
		t.FailNow()
	}
	assertInt(t, "number of years", len(is.Statements), 4)
	y2 := is.Statements[1]
	assertString(t, "filing type", y2.FilingType, "10-K")
	isPositiveInt(t, "fiscal year", y2.FiscalYear)
	assertString(t, "currency", y2.Currency, "USD")
}

func testIntegrationQuarterlyIncomeStatements(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(
		cfg.Token,
		iex.WithRateLimiter(time.Second, 10),
		iex.WithBaseURL(cfg.BaseURL),
	)
	is, err := client.QuarterlyIncomeStatements(context.Background(), "f", 4)
	if err != nil {
		log.Fatalf("Error getting quarterly income statements: %s", err)
	}
	assertString(t, "symbol", is.Symbol, "F")
	if len(is.Statements) != 4 {
		t.Errorf("\ngot = %d %s\nwant = %d", len(is.Statements), "number of quarters", 4)
		t.FailNow()
	}
	q3 := is.Statements[2]
	isPositiveInt(t, "fiscal quarter", q3.FiscalQuarter)
	isPositiveInt(t, "fiscal year", q3.FiscalYear)
	assertString(t, "currency", q3.Currency, "USD")
}

func testIntegrationBook(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(
		cfg.Token,
		iex.WithRateLimiter(time.Second, 10),
		iex.WithBaseURL(cfg.BaseURL),
	)
	got, err := client.Book(context.Background(), "aapl")
	if err != nil {
		log.Fatalf("Error getting book: %s", err)
	}
	assertString(t, "symbol", got.Quote.Symbol, "AAPL")
	assertString(t, "company name", got.Quote.CompanyName, "Apple Inc")
	assertScrambledString(t, "primary exchange", got.Quote.PrimaryExchange, "NASDAQ")
	assertStringFromOptions(t, "calculation price", got.Quote.CalculationPrice,
		[]string{"tops", "sip", "previousclose", "close", "iexlasttrade"})
	isNotNegativeFloat64(t, "open", got.Quote.Open)
	assertScrambledString(t, "open source", got.Quote.OpenSource, "official")
	isPositiveFloat64(t, "latest price", got.Quote.LatestPrice)
}

func testIntegrationHistoricalPrices(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(
		cfg.Token,
		iex.WithRateLimiter(time.Second, 10),
		iex.WithBaseURL(cfg.BaseURL),
	)
	timeframe := iex.OneMonthHistorical
	histPrices, err := client.HistoricalPrices(context.Background(), "aapl", timeframe, nil)
	if err != nil {
		log.Fatalf("Error getting historical prices: %s", err)
	}
	got := histPrices[0]
	isPositiveFloat64(t, "close", got.Close)
	isPositiveFloat64(t, "high", got.High)
	isPositiveFloat64(t, "low", got.Low)
	isPositiveFloat64(t, "open", got.Open)
	assertString(t, "symbol", got.Symbol, "AAPL")
	isPositiveFloat64(t, "volume", got.Volume)
	assertScrambledString(t, "id", got.ID, "HISTORICAL_PRICES")
	assertScrambledString(t, "key", got.Key, "AAPL")
	assertString(t, "subkey", got.Subkey, "")
}

func assertInt(t *testing.T, label string, got, want int) {
	if got != want {
		t.Errorf("\ngot = %d %s\nwant = %d", got, label, want)
	}
}

func assertFloat64(t *testing.T, label string, got, want, tolerance float64) {
	if diff := math.Abs(want - got); diff >= tolerance {
		t.Errorf("\ngot = %f %s\ntwant = %f", got, label, want)
	}
}

func assertBool(t *testing.T, label string, got, want bool) {
	if got != want {
		t.Errorf("\ngot = %t %s\nwant = %t", got, label, want)
	}
}

func assertString(t *testing.T, label string, got, want string) {
	if got != want {
		t.Errorf("\ngot = %s %s\nwant = %s", got, label, want)
	}
}

func assertStringFromOptions(t *testing.T, label string, got string, options []string) {
	isAnOption := false
	for _, option := range options {
		if got == option {
			isAnOption = true
			break
		}
	}
	if isAnOption == false {
		t.Errorf("\ngot = %s %s\nwant one of %s", got, label, options)
	}
}

// IEX scrambles their responses when using the testing sandbox. Therefore, the
// best we can do is assert that all the letters are there even if scrambled.
func assertScrambledString(t *testing.T, label string, got, want string) {
	gotSorted := sortString(got)
	wantSorted := sortString(want)
	if gotSorted != wantSorted {
		t.Errorf("\n got = %s %s\nwant = %s", got, label, want)
	}
}

func isPositiveInt(t *testing.T, label string, got int) {
	if got <= 0 {
		t.Errorf("\n got = %d %s\nwant int > 0", got, label)
	}
}

func isPositiveFloat64(t *testing.T, label string, got float64) {
	if got <= 0.0 {
		t.Errorf("\n got = %f %s\nwant float64 > 0.0", got, label)
	}
}

func isNotNegativeFloat64(t *testing.T, label string, got float64) {
	if got < 0.0 {
		t.Errorf("\n got = %f %s\nwant float64 >= 0.0", got, label)
	}
}

func isString(t *testing.T, label string, got string) {
	if got == "" {
		t.Errorf("\n got = %s %s\nwant non-empty string", got, label)
	}
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
