package main

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"go-hodl/ticker"
)

const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
)

func main() {
	t := ticker.New(10, "USD")
	t.UpdateData()

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Rank"},
			{Align: simpletable.AlignCenter, Text: "Coin"},
			{Align: simpletable.AlignCenter, Text: "Price (USD)"},
			{Align: simpletable.AlignCenter, Text: "Change 1H"},
			{Align: simpletable.AlignCenter, Text: "Change 24H"},
			{Align: simpletable.AlignCenter, Text: "Change 7D"},
			{Align: simpletable.AlignCenter, Text: "Market Cap (USD)"},
		},
	}

	for _, r := range t.Records {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: r.Rank},
			{Align: simpletable.AlignLeft, Text: r.Coin},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.8f", r.Price)},
			{Align: simpletable.AlignRight, Text: colorizePercentage(r.Change1H)},
			{Align: simpletable.AlignRight, Text: colorizePercentage(r.Change24H)},
			{Align: simpletable.AlignRight, Text: colorizePercentage(r.Change7D)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.0f", r.MarketCap)},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}

func colorizePercentage(val float64) string {
	if val < 0 {
		return fmt.Sprintf("%s%.2f%%%s", ColorRed, val, ColorDefault)
	}
	return fmt.Sprintf("%s%.2f%%%s", ColorGreen, val, ColorDefault)
}
