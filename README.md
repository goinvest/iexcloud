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

### Account

- [ ] Message Budget
- [x] Metadata
- [ ] Pay as you go
- [ ] Signed Requests
- [ ] Usage

### API System Metadata

- [x] Status

### Time Series

- [ ] Time Series

### Calendar

- [ ] Calendar

### Data Points

- [ ] Data Points

### Stock Prices

- [x] Book
- [x] Charts — Use historical and intraday price endpoints.
- [x] Delayed Quote
- [ ] Extended Hours Quote — Use quote endpoint.
- [ ] Historical Prices
- [ ] Intraday Prices
- [x] Largest Trades
- [x] Open / Close Price
- [x] OHLC
- [x] Previous Day Price
- [x] Price Only
- [x] Quote
- [ ] Real-time Quote — Use quote endpoint.
- [ ] Volume by Venue

### Stock Profiles

- [x] Company
- [x] Insider Roster
- [x] Insider Summary
- [x] Insider Transactions
- [x] Logo
- [ ] Peer Groups

### Stock Fundamentals

- [x] Balance Sheet
- [x] Cash Flow
- [x] Dividends (Basic)
- [x] Earnings
- [x] Financials
- [ ] Financials As Reported
- [x] Income Statement
- [ ] SEC Filings
- [ ] Splits (Basic)

### Stock Research

- [x] Advanced Stats
- [ ] Analyst Recommendations
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
- [ ] IPO Calendar
- [x] List
- [x] Market Volume (U.S.)
- [x] Sector Performance
- [ ] Upcoming Events

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

- [ ] Oil Prices
- [ ] Natural Gas Price
- [ ] Heating Oil Prices
- [ ] Jet Fuel Prices
- [ ] Diesel Price
- [ ] Gas Prices
- [ ] Propane Prices

### Economic Data

- [ ] CD Rates
- [ ] Consumer Price Index
- [ ] Credit Card Interest Rate
- [ ] Federal Fund Rates
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
- [ ] Sectors
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

### Need to Reclassify/Remove

- [x] Effective Spread (@FIXME: Has been removed from API.)
- [x] Recommendation Trends
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
