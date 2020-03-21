// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"fmt"
	"log"

	iex "github.com/goinvest/iexcloud/v2"
	"github.com/goinvest/iexcloud/v2/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(symbolsCmd)
}

var symbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "Retrieve an array of symbols IEX Cloud supports",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		symbols, err := client.Symbols(context.Background())
		if err != nil {
			log.Fatalf("Error getting symbols: %s", err)
		}
		count := len(symbols)
		fmt.Printf("IEX Cloud supports %d symbols\n", count)
		for _, s := range symbols {
			disabled := ""
			if !s.IsEnabled {
				disabled = "(disabled) "
			}
			fmt.Printf("%s%s\t%s\t%s\n", disabled, s.Symbol, s.Type, s.Name)
		}
	},
}
