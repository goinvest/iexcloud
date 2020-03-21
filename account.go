// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

// AccountMetadata provides details about an IEX Cloud account.
type AccountMetadata struct {
	PayAsYouGo       bool       `json:"overagesEnabled"`
	EffectiveDate    EpochTime  `json:"effectiveDate"`
	EndDateEffective *EpochTime `json:"endDateEffective"`
	SubscriptionTerm string     `json:"subscriptionTermType"`
	TierName         string     `json:"tierName"`
	MessageLimit     int        `json:"messageLimit"`
	MessagesUsed     int        `json:"messagesUsed"`
}

// Usage provides current month usage for your account.
type Usage struct {
	MonthlyUsage      int            `json:"monthlyUsage"`
	MonthlyPayAsYouGo int            `json:"monthlyPayAsYouGo"`
	DailyUsage        map[string]int `json:"dailyUsage"`
	TokenUsage        map[string]int `json:"tokenUsage"`
	KeyUsage          map[string]int `json:"keyUsage"`
}
