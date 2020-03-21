// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	iex "github.com/goinvest/iexcloud/v2"
	"github.com/goinvest/iexcloud/v2/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deepCmd)
	rootCmd.AddCommand(deepBookCmd)
	rootCmd.AddCommand(deepTradesCmd)
}

var deepCmd = &cobra.Command{
	Use:   "deep [stock]",
	Short: "Retrieve the deep data for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		deep, err := client.DEEP(context.Background(), stock)
		if err != nil {
			log.Fatalf("Error getting deep: %s", err)
		}
		b, err := json.MarshalIndent(deep, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## DEEP ##")
		fmt.Println(string(b))
	},
}

var deepBookCmd = &cobra.Command{
	Use:   "deep-book [stock]",
	Short: "Retrieve the deep book data for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		s := []string{stock}
		deep, err := client.DEEPBook(context.Background(), s)
		if err != nil {
			log.Fatalf("Error getting deep book: %s", err)
		}
		b, err := json.MarshalIndent(deep, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## DEEP Book ##")
		fmt.Println(string(b))
	},
}

var deepTradesCmd = &cobra.Command{
	Use:   "deep-trades [stock]",
	Short: "Retrieve the deep trades data for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		s := []string{stock}
		deep, err := client.DEEPTrades(context.Background(), s)
		if err != nil {
			log.Fatalf("Error getting deep trades: %s", err)
		}
		b, err := json.MarshalIndent(deep, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## DEEP Trades ##")
		fmt.Println(string(b))
	},
}
