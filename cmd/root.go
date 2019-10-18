package cmd

import (
	flag "github.com/spf13/pflag"
	"go-hodl/ticker"
	"os"
)

func Execute() int {
	var currency string
	var limit int
	var filters string
	flags := flag.NewFlagSet("hodl", flag.ExitOnError)
	flags.StringVarP(&currency, "currency", "c", "usd", "Currency prices to given currency")
	flags.StringVarP(&filters, "filter", "f", "", "Filter data to specific coins")
	flags.IntVarP(&limit, "limit", "l", 10, "Limit the number of cryptocurrencies is shown according to the CoinMarketCap ranking")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return 1
	}

	params := ticker.InitParams(limit, currency, filters)
	t := ticker.New(params)
	t.UpdateData()
	t.PrintTable()

	return 0
}
