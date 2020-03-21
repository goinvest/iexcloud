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
	rootCmd.AddCommand(fedFundsCmd)
}

var fedFundsCmd = &cobra.Command{
	Use:   "fedfunds",
	Short: "Retrieve the effective federal funds rate",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		rate, err := client.FederalFundsRate(context.Background())
		if err != nil {
			log.Fatalf("Error getting federal funds rate: %s", err)
		}
		fmt.Printf("Effective federal funds rate = %f\n", rate)
	},
}
