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
	rootCmd.AddCommand(dividendsCmd)
}

var dividendsCmd = &cobra.Command{
	Use:   "div [stock] [range]",
	Short: "Retrieve the dividends for stock symbol",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		r := args[1]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		// Make sure date range is valid.
		got, ok := iex.PathRanges[r]
		if !ok {
			log.Fatalf("Bad date range: %s", r)
		}
		log.Printf("Using date range = %s", got)
		dividends, err := client.Dividends(context.Background(), stock, got)
		if err != nil {
			log.Fatalf("Error getting dividends: %s", err)
		}
		b, err := json.MarshalIndent(dividends, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Dividends ##")
		fmt.Println(string(b))
	},
}
