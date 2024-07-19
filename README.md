# iexcloud

Go library for accessing the IEX Cloud Legacy API.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]

## Archived

As of August, 31, 2024, all [IEX Cloud][iexweb] products will be retired.
Therefore, the [iexcloud][] Go library is archived and all development has
ceased.

## Overview

[iexcloud][] provides a Go interface to the [IEX Cloud Legacy API][iexlegacy].
To access the [IEX Cloud Legacy API][iexlegacy] an account and token are
required. The goal is for [iexcloud][] to be compatible with the v1 version of
the IEX Cloud Legacy API. There were some changes from the beta version to v1 of
the API, so things may still be in flux for this library.

- [IEX Cloud Legacy API][iexlegacy] uses <https://cloud.iexapis.com/> for its
  base URL.
- [IEX Cloud API][iexapi] uses <https://api.iex.cloud/v1/> for its base URL.

## Installation

```bash
$ go get github.com/goinvest/iexcloud/v2
```

## Examples

Examples are available at <https://github.com/goinvest/iexcloud-examples/>.

## Implementation Status

Please see [implementation.md][implementation] for the current implementation
status of the [IEX Cloud Legacy API][iexlegacy].

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check    # formats, vets, and unit tests the code
$ make lint     # lints code using staticcheck
```

To update and view the test coverage report:

```bash
$ make cover
```

#### Integration Testing

To perform the integration tests run:

```bash
$ make int
```

Prior to doing so, you'll need to create a `config_test.toml` file with your IEX
Cloud API Token and the base URL. It is recommended to use your sandbox token
and the sandbox URL, so as to not be charged credits when running the
integration tests. Sandbox tokens start with `Tpk_` instead of `pk_` for
non-sandbox tokens. Using the sandbox does make integration a little more
difficult, since results are scrambled in sandbox mode.

Example `config_test.toml` file:

```toml
Token = "Tpk_your_iexcloud_test_token"
BaseURL = "https://sandbox.iexapis.com/v1"
```

## License

[iexcloud][] is released under the MIT license. Please see the
[LICENSE][] file for more information.

[iexapi]: https://iexcloud.io/docs/
[iexcloud]: https://github.com/goinvest/iexcloud
[iexlegacy]: https://iexcloud.io/docs/api/
[iexweb]: https://iexcloud.io
[godoc badge]: https://godoc.org/github.com/goinvest/iexcloud?status.svg
[godoc link]: https://godoc.org/github.com/goinvest/iexcloud
[implementation]: https://github.com/goinvest/iexcloud/blob/master/implementation.md
[LICENSE]: https://github.com/goinvest/iexcloud/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/goinvest/iexcloud
[report card]: https://goreportcard.com/report/github.com/goinvest/iexcloud
