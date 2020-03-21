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
	rootCmd.AddCommand(availableDataPoints)
	rootCmd.AddCommand(dataPoint)
}

var availableDataPoints = &cobra.Command{
	Use:   "data-points [symbol]",
	Short: "Retrieve the available data points for the symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		dataPoints, err := client.AvailableDataPoints(context.Background(), symbol)
		if err != nil {
			log.Fatalf("Error getting available data points: %s", err)
		}
		b, err := json.MarshalIndent(dataPoints, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Available Data Points ##")
		fmt.Println(string(b))
	},
}

var dataPoint = &cobra.Command{
	Use:   "data-point [symbol] [key]",
	Short: "Retrieve the data point for the symbol & key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		key := args[1]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		dataPoint, err := client.DataPoint(context.Background(), symbol, key)
		if err != nil {
			log.Fatalf("Error getting data point (symbol = %s / key = %s): %s", symbol, key, err)
		}
		fmt.Printf("%s %s = %s\n", symbol, key, dataPoint)
	},
}
