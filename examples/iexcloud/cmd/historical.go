// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	iex "github.com/goinvest/iexcloud"
	"github.com/goinvest/iexcloud/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(historicalCmd)
	rootCmd.AddCommand(historicalByDayCmd)
	rootCmd.AddCommand(intradayHistoricalCmd)
	rootCmd.AddCommand(intradayHistoricalByDayCmd)
}

var historicalCmd = &cobra.Command{
	Use:   "historical [stock] [timeframe]",
	Short: "Retrieve the historical data for stock symbol in given timeframe",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		timeframe := iex.HistoricalTimeFrame(args[1])
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		histPrices, err := client.HistoricalPrices(stock, timeframe, nil)
		if err != nil {
			log.Fatalf("Error getting historical prices: %s", err)
		}
		b, err := json.MarshalIndent(histPrices, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Historical Prices ##")
		fmt.Println(string(b))
	},
}

var historicalByDayCmd = &cobra.Command{
	Use:   "historicalbyday [stock] [day]",
	Short: "Retrieve the historical data for stock symbol for given day yyyymmddd",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		day, err := time.Parse("20060102", args[1])
		if err != nil {
			log.Fatalf("Error parsing date: %s", err)
		}

		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		histPrices, err := client.HistoricalPricesByDay(stock, day, nil)
		if err != nil {
			log.Fatalf("Error getting intraday historical prices: %s", err)
		}
		b, err := json.MarshalIndent(histPrices, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Printf("## Historical Prices For %s ##\n", args[1])
		fmt.Println(string(b))
	},
}

var intradayHistoricalCmd = &cobra.Command{
	Use:   "intradayhistorical [stock]",
	Short: "Retrieve the intraday data for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		histPrices, err := client.IntradayHistoricalPrices(stock, nil)
		if err != nil {
			log.Fatalf("Error getting intraday historical prices: %s", err)
		}
		b, err := json.MarshalIndent(histPrices, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Intraday Historical Prices ##")
		fmt.Println(string(b))
	},
}

var intradayHistoricalByDayCmd = &cobra.Command{
	Use:   "intradayhistoricalbyday [stock] [day]",
	Short: "Retrieve the intraday data for stock symbol for given day yyyymmddd",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		day, err := time.Parse("20060102", args[1])
		if err != nil {
			log.Fatalf("Error parsing date: %s", err)
		}

		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		histPrices, err := client.IntradayHistoricalPricesByDay(stock, day, nil)
		if err != nil {
			log.Fatalf("Error getting intraday historical prices: %s", err)
		}
		b, err := json.MarshalIndent(histPrices, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Printf("## Intraday Historical Prices For %s ##\n", args[1])
		fmt.Println(string(b))
	},
}
