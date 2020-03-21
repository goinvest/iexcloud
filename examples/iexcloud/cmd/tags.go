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
	rootCmd.AddCommand(tagsCmd)
}

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Retrieve an array of all tags IEX Cloud supports",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient()
		tags, err := client.Tags(context.Background())
		if err != nil {
			log.Fatalf("Error getting tags: %s", err)
		}
		count := len(tags)
		fmt.Printf("IEX Cloud supports %d tags\n", count)
		for _, t := range tags {
			fmt.Println(t.Name)
		}
	},
}
