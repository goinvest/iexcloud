# iexcloud

Go library for accessing the IEX Cloud API.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]

## Overview

[iexcloud][] provides a Go interface to the [IEX Cloud API][1]. To
access the [IEX Cloud API][1] an account and token are required. The goal is
for iexcloud to be compatible with the v1 version of the IEX Cloud API. There
were some changes from the beta version to v1 of the API, so things may still
be a in flux for this library.

## Installation

```bash
$ go get github.com/goinvest/iexcloud
```

## Examples

See the [iexcloud CLI example README][2].

## Implementation Status

Below is a list of the APIs that have and have not been implemented. If
you want a particular API to be developed next, please open an issue.

### Introduction

- [ ] Batch Requests

### Guides

- [ ] Time Series
- [ ] Calendar
- [x] Data Points
- [ ] Files

### Rules Engine

- [ ] Rules Schema
- [ ] Lookup Values
- [ ] Creating a Rule
- [ ] Pause and Unpause
- [ ] Edit an Existing Rule
- [ ] Delete a Rule
- [ ] Get Rule Info
- [ ] List All Rules
- [ ] Getting Logged Outputs

### Account

- [ ] Message Budget
- [x] Metadata
- [ ] Pay as you go
- [ ] Signed Requests
- [x] Usage

### API System Metadata

- [x] Status

### Stock Prices

- [x] Book
- [x] Charts — Use historical and intraday price endpoints.
- [x] Delayed Quote
- [x] Extended Hours Quote — Use Quote.
- [x] Historical Prices
- [x] Intraday Prices
- [x] Largest Trades
- [x] Open / Close Price - Use OHLC.
- [x] OHLC
- [x] Previous Day Price
- [x] Price Only
- [x] Quote
- [x] Real-time Quote — Use Quote.
- [x] Volume by Venue

### Stock Profiles

- [x] Company
- [x] Insider Roster
- [x] Insider Summary
- [x] Insider Transactions
- [x] Logo
- [x] Peer Groups

### Stock Fundamentals

- [x] Balance Sheet
- [x] Cash Flow
- [x] Dividends (Basic)
- [x] Earnings
- [x] Financials
- [ ] Financials As Reported
- [x] Income Statement
- [ ] SEC Filings — Use the Financials As Reported endpoint for raw SEC filings
      data.
- [ ] Splits (Basic)

### Stock Research

- [x] Advanced Stats
- [x] Analyst Recommendations
- [x] Estimates
- [x] Fund Ownership
- [x] Institutional Ownership
- [x] Key Stats
- [x] Price Target
- [ ] Technical Indicators

### Corporate Actions

- [ ] Bonus Issue
- [ ] Distribution
- [ ] Dividends
- [ ] Return of Capital
- [ ] Rights Issue
- [ ] Right to Purchase
- [ ] Security Reclassification
- [ ] Security Swap
- [ ] Spinoff
- [ ] Splits

### Market Info

- [x] Collections
- [x] Earnings Today
- [x] IPO Calendar
- [x] List
- [x] Market Volume (U.S.)
- [x] Sector Performance
- [x] Upcoming Events

### News

- [x] News
- [ ] Streaming News
- [ ] Historical News

### Cryptocurrency

- [ ] Cryptocurrency Book
- [ ] Cryptocurrency Events
- [ ] Cryptocurrency Price
- [ ] Cryptocurrency Quote

### Forex / Currencies

- [ ] Real-time Streaming
- [ ] Latest Currency Rates
- [ ] Currency Conversion
- [ ] Historical Daily

### Options

- [ ] End of Day Options

### Social Sentiment

- [ ] Social Sentiment

### CEO Compensation

- [x] CEO Compensation

### Treasuries

- [ ] Daily Treasury Rates

### Commodities

- [x] Oil Prices
- [x] Natural Gas Price
- [x] Heating Oil Prices
- [x] Jet Fuel Prices
- [x] Diesel Price
- [x] Gas Prices
- [x] Propane Prices

### Economic Data

- [x] CD Rates
- [x] Consumer Price Index
- [x] Credit Card Interest Rate
- [x] Federal Fund Rates
- [ ] Real GDP
- [ ] Institutional Money Funds
- [ ] Initial Claims
- [ ] Industrial Production Index
- [ ] Mortgage Rates
- [ ] Total Housing Starts
- [ ] Total Payrolls
- [ ] Total Vehicle Sales
- [ ] Retail Money Funds
- [ ] Unemployment Rates
- [ ] US Recession Probabilities

### Reference Data

- [ ] Search
- [x] Cryptocurrency Symbols
- [x] FX Symbols
- [x] IEX Symbols
- [ ] International Symbols
- [ ] International Exchanges
- [ ] ISIN Mapping
- [x] Mutual Fund Symbols
- [ ] Options Symbols
- [x] OTC Symbols
- [x] Sectors
- [x] Symbols
- [x] Tags
- [x] U.S. Exchanges
- [x] U.S. Holidays and Trading Days

### Investors Exchange Data

- [x] DEEP
- [ ] DEEP Auction
- [x] DEEP Book
- [ ] DEEP Operational Halt Status
- [ ] DEEP Official Price
- [ ] DEEP Security Event
- [ ] DEEP Short Sale Price Test Status
- [ ] DEEP System Event
- [x] DEEP Trades
- [ ] DEEP Trade Break
- [ ] DEEP Trading Status
- [x] Last
- [ ] Listed Regulation SHO Threshold Securities List
- [ ] Stats Historical Daily
- [ ] Stats Historical Summary
- [x] Stats Intraday
- [ ] Stats Recent
- [ ] Stats Records
- [x] TOPS

### Wall Street Horizon

- [ ] Analyst Days
- [ ] Board of Directors Meeting
- [ ] Business Updates
- [ ] Buybacks
- [ ] Capital Markets Day
- [ ] Company Travel
- [ ] Filing Due Dates
- [ ] Fiscal Quarter End
- [ ] Forum
- [ ] General Conference
- [ ] FDA Advisory Committee Meetings
- [ ] Holidays
- [ ] Index Changes
- [ ] IPOs
- [ ] Legal Actions
- [ ] Mergers & Acquisitions
- [ ] Product Events
- [ ] Research and Development Days
- [ ] Same Store Sales
- [ ] Secondary Offerings
- [ ] Seminars
- [ ] Shareholder Meetings
- [ ] Summit Meetings
- [ ] Trade Shows
- [ ] Witching Hours
- [ ] Workshops

### Fraud Factors

- [ ] Similarity Index
- [ ] Non-Timely Filings

### ExtractAlpha

### Precision Alpha

- [ ] Precision Alpha Price Dynamics

### BRAIN Company

### Kavout

- [ ] K Score for US Equities

### Audit Analytics

- [ ] Audit Analytics Director and Officer Changes
- [ ] Audit Analytics Accounting Quality and Risk Matrix

### ValuEngine

- [ ] ValuEngine Stock Research Report

### Need to Reclassify/Remove

- [x] Relevant Stocks

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check
```

To update and view the test coverage report:

```bash
$ make cover
```

## License

[iexcloud][] is released under the MIT license. Please see the
[LICENSE][] file for more information.

[1]: https://iexcloud.io
[2]: https://github.com/goinvest/iexcloud/blob/master/examples/iexcloud/README.md
[iexcloud]: https://github.com/goinvest/iexcloud
[godoc badge]: https://godoc.org/github.com/goinvest/iexcloud?status.svg
[godoc link]: https://godoc.org/github.com/goinvest/iexcloud
[LICENSE]: https://github.com/goinvest/iexcloud/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/goinvest/iexcloud
[report card]: https://goreportcard.com/report/github.com/goinvest/iexcloud
