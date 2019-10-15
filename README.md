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
### Limit coin(s)
You can use the `--limit` flag to limit the number of cryptocurrencies is shown according to the CoinMarketCap ranking.
```shell script
$ ./bin/hodl --limit 20
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
