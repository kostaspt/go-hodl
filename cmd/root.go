package cmd

import (
	flag "github.com/spf13/pflag"
	"go-hodl/ticker"
	"os"
)

func Execute() int {
	var currency string
	var limit int
	flags := flag.NewFlagSet("hodl", flag.ExitOnError)
	flags.StringVarP(&currency, "currency", "c", "USD", "Currency prices to given currency")
	flags.IntVarP(&limit, "limit", "l", 10, "Limit the number of cryptocurrencies is shown according to the CoinMarketCap ranking")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return 1
	}

	t := ticker.New(limit, currency)
	t.UpdateData()
	t.PrintTable()

	return 0
}
