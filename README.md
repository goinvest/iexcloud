# iexcloud

Go library for accessing the IEX Cloud API.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]

## Overview

[iexcloud][] provides a Go interface to the [IEX Cloud API][1]. To
access the [IEX Cloud API][1] an account and token are required.

## Installation

```bash
$ go get github.com/goinvest/iexcloud
```

## Examples

See the [iexcloud CLI example README][2].

## Implementation Status

Below is a list of the APIs that have and have not been implemented. If
you want a particular API to be developed next, please open an issue.

### Account

- [x] Metadata
- [ ] Usage
- [ ] Pay as you go

### Stocks

- [x] Advanced Stats
- [x] Balance Sheet
- [ ] Batch Requests
- [x] Book
- [x] Cash Flow
- [x] Collections
- [x] Company
- [x] Delayed Quote
- [x] Dividends
- [x] Earnings
- [x] Earnings Today
- [x] Effective Spread (@FIXME: Has been removed from API.)
- [x] Estimates
- [x] Financials
- [ ] Financials As Reported
- [x] Fund Ownership
- [x] Nondynamic Historical Prices
- [ ] Dynamic Historical Prices
- [x] Income Statement
- [x] Insider Roster
- [x] Insider Summary
- [x] Insider Transactions
- [x] Institutional Ownership
- [ ] IPO Calendar
- [x] Key Stats
- [x] Largest Trades
- [x] List
- [x] Logo
- [x] Market Volume (U.S.)
- [x] News
- [x] OHLC
- [x] Open / Close Price
- [x] Peers
- [x] Previous Day Prices
- [x] Price
- [x] Price Target
- [x] Quote
- [x] Recommendation Trends
- [x] Relevant Stocks
- [x] Sector Performance
- [ ] Splits
- [ ] Volume by Venue

### Alternative Data

- [x] Crypto
- [ ] Social Sentiment
- [x] CEO Compensation

### Reference Data

- [x] Symbols
- [x] IEX Symbols
- [x] U.S. Exchanges
- [x] U.S. Holidays and Trading Days
- [x] Stock Tags
- [x] Stock Collections
- [x] Mutual Fund Symbols
- [x] OTC Symbols
- [x] FX Symbols
- [ ] Options Symbols
- [ ] Commodities Symbols
- [ ] Bonds Symbols
- [x] Crypto Symbols

### Forex / Currencies

- [x] Exchange Rates

### Investors Exchange Data

- [x] TOPS
- [x] Last
- [x] DEEP
- [ ] DEEP Auction
- [x] DEEP Book
- [ ] DEEP Operational Halt Status
- [ ] DEEP Official Price
- [ ] DEEP Security Event
- [ ] DEEP Short Sale Price Tst Status
- [ ] DEEP System Event
- [x] DEEP Trades
- [ ] DEEP Trade Break
- [ ] DEEP Trading Status
- [ ] Listed Regulation SHO Threshold Securities List
- [ ] Listed Short Interest List
- [ ] Stats Historical Daily
- [ ] Stats Historical Summary
- [x] Stats Intraday
- [ ] Stats Recent
- [ ] Stats Records

### API System Metadata

- [x] Status

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
