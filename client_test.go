// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex_test

import (
	"context"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"testing"

	"github.com/BurntSushi/toml"
	iex "github.com/goinvest/iexcloud/v2"
)

// Config contains the configuration information neecded to program and test
// the adapaters.
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
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return cfg, err
	}
	err = toml.Unmarshal(buf, &cfg)
	return cfg, err
}

func TestHistoricalPrices(t *testing.T) {
	cfg, err := readConfig("config_test.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
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
	isPositiveInt(t, "volume", got.Volume)
	assertScrambledString(t, "id", got.ID, "HISTORICAL_PRICES")
	assertScrambledString(t, "key", got.Key, "AAPL")
	assertString(t, "subkey", got.Subkey, "")

}

func assertInt(t *testing.T, label string, got, want int) {
	if got != want {
		t.Errorf("\t got = %d %s\n\t\twant = %d", got, label, want)
	}
}

func assertFloat64(t *testing.T, label string, got, want, tolerance float64) {
	if diff := math.Abs(want - got); diff >= tolerance {
		t.Errorf("\t got = %f %s\n\t\t\twant = %f", got, label, want)
	}
}

func assertBool(t *testing.T, label string, got, want bool) {
	if got != want {
		t.Errorf("\t got = %t %s\n\t\t\twant = %t", got, label, want)
	}
}

func assertString(t *testing.T, label string, got, want string) {
	if got != want {
		t.Errorf("\t got = %s %s\n\t\t\twant = %s", got, label, want)
	}
}

// IEX scrambles their responses when using the testing sandbox. Therefore, the
// best we can do is assert that all the letters are there even if scrambled.
func assertScrambledString(t *testing.T, label string, got, want string) {
	gotSorted := sortString(got)
	wantSorted := sortString(want)
	if gotSorted != wantSorted {
		t.Errorf("\t got = %s %s\n\t\t\twant = %s", got, label, want)
	}
}

func isPositiveInt(t *testing.T, label string, got int) {
	if got <= 0 {
		t.Errorf("\t got = %d %s\n\t\twant int > 0", got, label)
	}
}

func isPositiveFloat64(t *testing.T, label string, got float64) {
	if got <= 0.0 {
		t.Errorf("\t got = %f %s\n\t\twant float64 > 0.0", got, label)
	}
}

func isString(t *testing.T, label string, got string) {
	if got == "" {
		t.Errorf("\t got = %s %s\n\t\twant non-empty string", got, label)
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
