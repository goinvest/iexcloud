# iexcloud

Go library for accessing the IEX Cloud API.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]

## Overview

[iexcloud][] provides a Go interface to the [IEX Cloud API][1]. To
access the [IEX Cloud API][1] an account and token is required.

## Installation

```bash
$ go get github.com/goinvest/iexcloud
```

## Examples

See the [iexcloud CLI example README][2].

## Implementation Status

Below is a list of the APIs that have and have not been implemented.

### Stocks

- [x] Balance Sheet
- [ ] Batch Requests
- [ ] Book
- [x] Cash Flow
- [ ] Collections
- [x] Company
- [x] Delayed Quote
- [x] Dividends
- [x] Earnings
- [ ] Effective Spread
- [x] Estimates
- [x] Financials
- [ ] Historical Prices
- [x] Income Statement
- [ ] IPO Calendar
- [x] Key Stats
- [ ] Largest Trades
- [ ] List
- [x] Logo
- [ ] Market Value (U.S.)
- [ ] News
- [ ] OHLC
- [ ] Open / Close Price
- [x] Peers
- [x] Previous Day Prices
- [x] Price
- [ ] Price Target
- [x] Quote
- [ ] Relevant Stocks
- [ ] Sector Performance
- [ ] Splits
- [ ] Volume by Venue

### Alternative Data

- [ ] News
- [ ] Crypto

### Reference Data

- [x] Symbols
- [ ] IEX Symbols
- [ ] U.S. Exchanges
- [ ] U.S. Holidays and Trading Days
- [ ] Stock Tags
- [ ] Stock Collections
- [ ] Mutual Fund Symbols
- [ ] OTC Symbols
- [ ] Forex / Currency Symbols
- [ ] Options Symbols
- [ ] Commodities Symbols
- [ ] Bonds Symbols
- [ ] Crypto Symbols

### Investors Exchange Data

- [ ] TOPS
- [ ] Last
- [ ] DEEP
- [ ] DEEP Auction
- [ ] DEEP Book
- [ ] DEEP Operational Halt Status
- [ ] DEEP Official Price
- [ ] DEEP Security Event
- [ ] DEEP Short Sale Price Tst Status
- [ ] DEEP System Event
- [ ] DEEP Trades
- [ ] DEEP Trade Break
- [ ] DEEP Trading Status
- [ ] Listed Regulation SHO Threshold Securities List
- [ ] Listed Short Interest List
- [ ] Stats Historical Daily
- [ ] Stats Historical Summary
- [ ] Stats Intraday
- [ ] Stats Recent
- [ ] Stats Records

### API System Metadata

- [ ] Status

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


[1]:https://iexcloud.io
[2]: https://github.com/goinvest/iexcloud/blob/master/examples/iexcloud/README.md
[iexcloud]: https://github.com/goinvest/iexcloud
[godoc badge]: https://godoc.org/github.com/goinvest/iexcloud?status.svg
[godoc link]: https://godoc.org/github.com/goinvest/iexcloud
[LICENSE]: https://github.com/goinvest/iexcloud/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/goinvest/iexcloud
[report card]: https://goreportcard.com/report/github.com/goinvest/iexcloud
