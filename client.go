// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"net/http"
)

const apiURL = "https://cloud.iexapis.com"

// Client models a client to consume the IEX Cloud API.
type Client struct {
	baseURL    string
	bearer     string
	httpClient *http.Client
}

// NewClient creates a client with the given authorization toke.
func NewClient(token string) *Client {
	return &Client{
		baseURL:    apiURL,
		bearer:     "Bearer " + token,
		httpClient: &http.Client{},
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(endpoint string, v interface{}) error {
	uri := c.baseURL + endpoint
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", c.bearer)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}
