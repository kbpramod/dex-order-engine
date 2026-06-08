package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	url := "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	var priceResp PriceResponse

	err = json.Unmarshal(body, &priceResp)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(priceResp)
}
