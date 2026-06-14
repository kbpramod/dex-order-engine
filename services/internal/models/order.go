package models

type OrderRequest struct {
	Symbol   string  `json:"symbol"`
	Side     string  `json:"side"`
	Type     string  `json:"type"`
	Quantity float64 `json:"quantity"`
}

type OrderResponse struct {
	Id     int64  `json:"orderid"`
	Status string `json:"status"`
}