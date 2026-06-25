package dto

type StockRequest struct {
	Action   string `json:"action"`
	Quantity int64  `json:"quantity"`
}
