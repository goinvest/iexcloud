// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sectorsCmd)
}

var sectorsCmd = &cobra.Command{
	Use:   "sectors",
	Short: "Retrieve an array of all sectors IEX Cloud supports",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient()
		sectors, err := client.Sectors(context.Background())
		if err != nil {
			log.Fatalf("Error getting sectors: %s", err)
		}
		count := len(sectors)
		fmt.Printf("IEX Cloud supports %d sectors\n", count)
		for _, s := range sectors {
			fmt.Println(s.Name)
		}
	},
}
