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
	rootCmd.AddCommand(accountMetadataCmd)
}

var accountMetadataCmd = &cobra.Command{
	Use:   "accountmetadata",
	Short: "Retrieve metadata information associated with secret key account",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		bs, err := client.AccountMetadata(context.Background())
		if err != nil {
			log.Fatalf("Error getting account metadata: %s", err)
		}
		b, err := json.MarshalIndent(bs, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling AccountMetadata into JSON: %s", err)
		}
		fmt.Println("## Account Metadata ##")
		fmt.Println(string(b))
	},
}
