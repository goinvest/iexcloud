// Copyright (c) 2019 The iexcloud developers. All rights reserved.
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
	rootCmd.AddCommand(mostactiveCmd)
}

var mostactiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Retrieve quotes for the 10 most active stocks",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		q, err := client.MostActive(context.Background())
		if err != nil {
			log.Fatalf("Error getting quotes: %s", err)
		}
		b, err := json.MarshalIndent(q, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Most Active Quotes ##")
		fmt.Println(string(b))
	},
}
