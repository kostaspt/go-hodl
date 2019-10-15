package main

import (
	"go-hodl/ticker"
)

func main() {
	t := ticker.New(10, "USD")
	t.UpdateData()
	t.PrintTable()
}
