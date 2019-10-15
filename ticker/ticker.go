package ticker

import (
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
)

type Ticker struct {
	ApiUrl  *url.URL
	Params  requestParams
	Records []record
	// TODO: Keep track of the last date it called the API
}

type requestParams struct {
	Limit    int
	Currency string
}

type record struct {
	Rank      string
	Coin      string
	Price     float64
	Change1H  float64
	Change24H float64
	Change7D  float64
	MarketCap float64
}

func New(limit int, currency string) *Ticker {
	params := initParams(limit, currency)
	apiUrl := generateUrl(params)
	return &Ticker{
		ApiUrl: apiUrl,
		Params: params,
	}
}

func (t *Ticker) UpdateData() {
	var respItems []map[string]interface{}
	respBody := t.callApi()
	defer respBody.Close()
	_ = json.NewDecoder(respBody).Decode(&respItems)

	priceKey := "price_"+strings.ToLower(t.Params.Currency)
	marketCapKey := "market_cap_"+strings.ToLower(t.Params.Currency)

	t.Records = nil
	for _, item := range respItems {
		price, _ := strconv.ParseFloat(item[priceKey].(string), 64)
		change1h, _ := strconv.ParseFloat(item["percent_change_1h"].(string), 64)
		change24h, _ := strconv.ParseFloat(item["percent_change_24h"].(string), 64)
		change7d, _ := strconv.ParseFloat(item["percent_change_7d"].(string), 64)
		marketCap, _ := strconv.ParseFloat(item[marketCapKey].(string), 64)
		r := record{
			Rank:      item["rank"].(string),
			Coin:      item["symbol"].(string),
			Price:     price,
			Change1H:  change1h,
			Change24H: change24h,
			Change7D:  change7d,
			MarketCap: marketCap,
		}
		t.Records = append(t.Records, r)
	}
}

func (t Ticker) callApi() io.ReadCloser {
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(t.ApiUrl.String())
	if err != nil {
		log.Fatal("Unable to fetch data from the API")
	}
	return resp.Body
}

func (t *Ticker) PrintTable() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Rank"},
			{Align: simpletable.AlignCenter, Text: "Coin"},
			{Align: simpletable.AlignCenter, Text: "Price (" + strings.ToUpper(t.Params.Currency) + ")"},
			{Align: simpletable.AlignCenter, Text: "Change 1H"},
			{Align: simpletable.AlignCenter, Text: "Change 24H"},
			{Align: simpletable.AlignCenter, Text: "Change 7D"},
			{Align: simpletable.AlignCenter, Text: "Market Cap (" + strings.ToUpper(t.Params.Currency) + ")"},
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

func initParams(limit int, currency string) requestParams {
	if limit == 0 {
		limit = 10
	}
	if currency == "" {
		currency = "USD"
	}
	return requestParams{
		Limit:    limit,
		Currency: currency,
	}
}

func generateUrl(params requestParams) *url.URL {
	u, _ := url.Parse("https://api.coinmarketcap.com/v1/ticker")
	q := u.Query()
	q.Set("limit", strconv.Itoa(params.Limit))
	q.Set("convert", params.Currency)
	u.RawQuery = q.Encode()
	return u
}

func colorizePercentage(val float64) string {
	if val < 0 {
		return fmt.Sprintf("%s%.2f%%%s", ColorRed, val, ColorDefault)
	}
	return fmt.Sprintf("%s%.2f%%%s", ColorGreen, val, ColorDefault)
}
