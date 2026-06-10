package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"binance-service/models"
)

func GetPrice(symbol string) (*models.PriceResponse, error) {

	url := fmt.Sprintf(
		"https://api.binance.com/api/v3/ticker/price?symbol=%s",
		symbol,
	)

	requestedAt := time.Now().UnixMilli()

	resp, err := http.Get(url)

	receivedAt := time.Now().UnixMilli()

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"binance returned status %d",
			resp.StatusCode,
		)
	}

	var price models.PriceResponse

	err = json.NewDecoder(resp.Body).Decode(&price)

	price.RequestedAt = requestedAt
	price.ReceivedAt = receivedAt

	if err != nil {
		return nil, err
	}

	return &price, nil
}

func CreateOrder(data models.OrderRequest) (*models.OrderResponse, error) {

	fmt.Println("order recieved", data)

	var order models.OrderResponse

	order.Id = 1234
	order.Status = "received"

	return &order, nil
}