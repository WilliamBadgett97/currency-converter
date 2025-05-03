package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)


type Currency struct {
	Rates map[string]float64 `json:"rates"`
}

var (
	from string
	to string
	amount string
)

func getCurrency() ([]byte, error) {
	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=36179924544d43028850a36d66169c43")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body, nil

}

func convertCurrency(amount, to string) {
	body, err := getCurrency()
	if err != nil {
		fmt.Println("Unable to fetch.")
	}
	var currencies Currency
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		fmt.Println("Error parsing Json.")
	}
	value, ok := currencies.Rates[to]
	if !ok {
		fmt.Println("Not a valid currency")
	}
	toFloat, _ := strconv.ParseFloat(strings.TrimSpace(amount), 64)
	converted := (toFloat / 1) * value
	fmt.Println(converted)

}


func main () {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Choose Base currency").Options(
				huh.NewOption("USD", "Usd"),
			).Value(&from),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Choose a currency to convert to").Options(
				huh.NewOption("AUD", "AUD"),
				huh.NewOption("EUR", "EUR"),
				huh.NewOption("GBP", "GBP"),
				huh.NewOption("JPY", "JPY"),
			).Value(&to),
		),

		huh.NewGroup(
			huh.NewInput().Title(fmt.Sprintf("Amount To convert")).Value(&amount),
		),
	)
	err := form.Run()
	convertCurrency(amount,to)
if err != nil {
    log.Fatal(err)
}
// convertCurrency()
}