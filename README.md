# iexcloud

Go library for accessing the IEX Cloud API.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]

## Overview

[iexcloud][] provides a Go interface to the [IEX Cloud API][iexcloudio]. To
access the [IEX Cloud API][iexcloudio] an account and token are required. The
goal is for iexcloud to be compatible with the v1 version of the IEX Cloud API.
There were some changes from the beta version to v1 of the API, so things may
still be in flux for this library.

## Installation

```bash
$ go get github.com/goinvest/iexcloud/v2
```

## Examples

Examples are available at <https://github.com/goinvest/iexcloud-examples/>.

## Implementation Status

Please see [implementation.md][implementation] for the current implementation
status of the IEX Cloud API.

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

[iexcloudio]: https://iexcloud.io
[iexcloud]: https://github.com/goinvest/iexcloud
[godoc badge]: https://godoc.org/github.com/goinvest/iexcloud?status.svg
[godoc link]: https://godoc.org/github.com/goinvest/iexcloud
[implementation]: https://github.com/goinvest/iexcloud/blob/master/implementation.md
[LICENSE]: https://github.com/goinvest/iexcloud/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/goinvest/iexcloud
[report card]: https://goreportcard.com/report/github.com/goinvest/iexcloud
