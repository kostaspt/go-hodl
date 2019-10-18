# HODL
> Command-line tool to track cryptocurrency prices and more.

Hodl is heavily inspired by [Coinmon](https://github.com/bichenkk/coinmon).

For the time being, all data comes from [coinmarketcap.com](https://coinmarketcap.com/) API.

## Build
```shell script
$ make build
```

## Usage
To check the top 10 cryptocurrencies, run:
```shell script
$ ./bin/hodl
```

## Options
### Limit coins
You can use the `-l` (or `--limit`) flag to limit the number of
cryptocurrencies is shown according to the CoinMarketCap ranking.
```shell script
$ ./bin/hodl -l 20
$ ./bin/hodl --limit 20
```

### Convert to specific currency
You can use the `-c` (or `--currency`) flag to get prices converted to
another currency.

The default currency is USD and it supports AUD, BRL, CAD, CHF, CLP, CNY,
CZK, DKK, EUR, GBP, HKD, HUF, IDR, ILS, INR, JPY, KRW, MXN, MYR, NOK,
NZD, PHP, PKR, PLN, RUB, SEK, SGD, THB, TRY, TWD, ZAR.
```shell script
$ ./bin/hodl -c EUR
$ ./bin/hodl --currency EUR
```

### Filter specific coin(s)
You can use the `-f` (or `--filter`) flag to filter with coin symbol(s)
to search cryptocurrencies. You can add symbols separated by comma. 
```shell script
$ ./bin/hodl -f btc
$ ./bin/hodl --filter btc,eth
```

## License
This project is licensed under the MIT License - see the
[LICENSE.md](LICENSE) file for details.
