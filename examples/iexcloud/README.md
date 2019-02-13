# IEX Cloud CLI

This Go Command Line Interface (CLI) application retrieves various data
from the IEX Cloud API.

## Example commands

The `iexcloud` executable expects a `config.toml` file with the
`BaseURL` and your IEX Cloud `Token`. The full list of commands can be
seen be running `./iexcloud -h`. Some of the available commands include:

```bash
$ ./iexcloud company aapl
$ ./iexcloud price aapl
$ ./iexcloud quote appl
```
