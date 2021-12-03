// Copyright (c) 2019-2022 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateToString(t *testing.T) {
	tests := []struct {
		input Date
		want  string
	}{
		{Date(time.Date(2021, 8, 1, 2, 3, 4, 5, time.UTC)), "2021-08-01"},
		{Date(time.Date(1999, 12, 31, 2, 3, 4, 5, time.UTC)), "1999-12-31"},
		{Date(time.Date(2024, 2, 29, 2, 3, 4, 5, time.UTC)), "2024-02-29"},
	}
	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf(
				"Error getting Date string\n\tgot %v; want %v",
				got,
				test.want,
			)
		}
	}
}

func TestDateUnmarhshalJSON(t *testing.T) {
	tests := []struct {
		input       []byte
		want        Date
		expectError bool
	}{
		{[]byte(`"2021-01-01"`), Date(time.Date(2021, 1, 1, 2, 3, 4, 5, time.UTC)), false},
		{[]byte(`"2024-02-29"`), Date(time.Date(2024, 2, 29, 2, 3, 4, 5, time.UTC)), false},
		{[]byte(`"foo"`), Date(time.Date(2021, 2, 1, 2, 3, 4, 5, time.UTC)), true},
		{[]byte(`"21-08-01"`), Date(time.Date(2021, 3, 1, 2, 3, 4, 5, time.UTC)), true},
		{[]byte(`""`), Date(time.Date(1929, 10, 24, 2, 3, 4, 5, time.UTC)), false},
	}
	for _, test := range tests {
		var got Date
		err := json.Unmarshal(test.input, &got)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error unmarshaling using %v", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error unmarshaling %s: %s", test.input, err)
			}
			if got.String() != test.want.String() {
				t.Errorf("got = %v / want = %v", got, test.want)
			}
		}
	}
}

func TestDateMarshalJSON(t *testing.T) {
	tests := []struct {
		input       Date
		want        []byte
		expectError bool
	}{
		{Date(time.Date(2021, 1, 1, 2, 3, 4, 5, time.UTC)), []byte(`"2021-01-01"`), false},
	}
	for _, test := range tests {
		got, err := json.Marshal(&test.input)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error marshaling using %v", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error marshaling %s: %s", test.input, err)
			}
			if string(got) != string(test.want) {
				t.Errorf("Error marshaling %v : got = %s / want = %s", test.input, got, test.want)
			}
		}
	}

}
