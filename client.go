// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const apiURL = "https://cloud.iexapis.com/beta"

// Client models a client to consume the IEX Cloud API.
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// NewClient creates a client with the given authorization toke.
func NewClient(token string, baseURL string) *Client {
	if baseURL == "" {
		baseURL = apiURL
	}
	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(endpoint string, v interface{}) error {
	address, err := c.addToken(endpoint)
	if err != nil {
		return err
	}
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

// GetJSONWithoutToken gets the JSON data from the given endpoint without
// adding a token to the URL.
func (c *Client) GetJSONWithoutToken(endpoint string, v interface{}) error {
	address := c.baseURL + endpoint
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) addToken(endpoint string) (string, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return "", err
	}
	v := u.Query()
	v.Add("token", c.token)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

// GetFloat64 gets the number from the given endpoint.
func (c *Client) GetFloat64(endpoint string) (float64, error) {
	address := c.baseURL + endpoint + "?token=" + c.token
	resp, err := http.Get(address)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(string(b), 64)

}
