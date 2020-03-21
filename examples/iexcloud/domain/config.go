// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package domain

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Config contains the configuration information neecded to program and test
// the adapaters.
type Config struct {
	Token   string
	BaseURL string
}

// ReadConfig will read the TOML config file.
func ReadConfig(configFile string) (Config, error) {

	var cfg Config

	// Read config file
	f, err := os.Open(configFile)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return cfg, err
	}
	err = toml.Unmarshal(buf, &cfg)
	return cfg, err
}
