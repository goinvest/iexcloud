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
	"strconv"

	iex "github.com/goinvest/iexcloud/v2"
	"github.com/goinvest/iexcloud/v2/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(earningsCmd)
}

var earningsCmd = &cobra.Command{
	Use:   "earnings [stock] [num]",
	Short: "Retrieve the specified num of earnings for stock symbol",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		num, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Provided %s. Needed an int: %s", args[1], err)
		}
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		earnings, err := client.Earnings(context.Background(), stock, num)
		if err != nil {
			log.Fatalf("Error getting earnings: %s", err)
		}
		b, err := json.MarshalIndent(earnings, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Earnings ##")
		fmt.Println(string(b))
	},
}
