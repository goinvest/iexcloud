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
$ make check
```

To update and view the test coverage report:

```bash
$ make cover
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
