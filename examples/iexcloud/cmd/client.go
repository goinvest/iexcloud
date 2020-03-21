// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"log"

	iex "github.com/goinvest/iexcloud/v2"
	"github.com/goinvest/iexcloud/v2/examples/iexcloud/domain"
)

// GetClient uses the supplied token and base URL from
// the config file and returns a created client.
func GetClient() *iex.Client {
	cfg, err := domain.ReadConfig(configFileFlag)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	return iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
}
