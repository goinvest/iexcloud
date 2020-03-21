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
	rootCmd.AddCommand(oilCmd)
}

var oilCmd = &cobra.Command{
	Use:   "oil [type]",
	Short: "Retrieve oil type of either west or brent",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oilText := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		var oilType iex.CommodityType
		switch oilText {
		case "west":
			oilType = iex.WestTexasOil
		case "brent":
			oilType = iex.BrentEuropeOil
		default:
			oilType = iex.WestTexasOil
		}

		price, err := client.CommodityPrice(context.Background(), oilType)
		if err != nil {
			log.Fatalf("Error getting commodity price: %s", err)
		}
		fmt.Printf("%s oil price = %f\n", oilType, price)
	},
}
