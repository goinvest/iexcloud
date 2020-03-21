// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iexcloud",
	Short: "Retrieve data from the IEX Cloud API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configFileFlag string

func init() {
	rootCmd.Flags().StringVarP(&configFileFlag, "config", "c", "config.toml", "config file")
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
