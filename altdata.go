// Copyright (c) 2019-2024 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// CEOCompensation models the compensation for a company's CEO.
type CEOCompensation struct {
	Symbol              string `json:"symbol"`
	Name                string `json:"name"`
	Company             string `json:"companyName"`
	Location            string `json:"location"`
	Salary              int    `json:"salary"`
	Bonus               int    `json:"bonus"`
	StockAwards         int    `json:"stockAwards"`
	OptionAwards        int    `json:"optionAwards"`
	NonEquityIncentives int    `json:"nonEquityIncentives"`
	PensionAndDeferred  int    `json:"pensionAndDeferred"`
	OtherCompensation   int    `json:"otherComp"`
	Total               int    `json:"total"`
	Year                int    `json:"year"`
}
