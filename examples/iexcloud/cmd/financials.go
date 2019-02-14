// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	iex "github.com/goinvest/iexcloud"
	"github.com/goinvest/iexcloud/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(financialsCmd)
}

var financialsCmd = &cobra.Command{
	Use:   "financials [stock]",
	Short: "Retrieve the financials for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, cfg.BaseURL)
		financials, err := client.QuarterlyFinancials(stock, 4)
		if err != nil {
			log.Fatalf("Error getting quarterly financials: %s", err)
		}
		b, err := json.MarshalIndent(financials, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Quarterly financials ##")
		fmt.Println(string(b))
	},
}
