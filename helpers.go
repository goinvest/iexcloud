// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date models a report date
type Date time.Time

// UnmarshalJSON implements the Unmarshaler interface for ReportDate.
func (d *ReportDate) UnmarshalJSON(data []byte) error {
	var aux string
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return fmt.Errorf("error unmarshaling date to string: %s", err)
	}
	t, err := time.Parse("2006-01-02", aux)
	if err != nil {
		return fmt.Errorf("error converting string to date: %s", err)
	}
	*d = ReportDate(t)
	return nil
}
