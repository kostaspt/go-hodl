package main

import (
	"flag"
	"go-hodl/ticker"
	"os"
)

func main() {
	var currency string
	var limit int
	flags := flag.NewFlagSet("hodl", flag.ExitOnError)
	flags.StringVar(&currency, "currency", "USD", "Currency prices to given currency")
	flags.IntVar(&limit, "limit", 10, "Limit the number of cryptocurrencies is shown according to the CoinMarketCap ranking")

	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	t := ticker.New(limit, currency)
	t.UpdateData()
	t.PrintTable()

	os.Exit(0)
}
