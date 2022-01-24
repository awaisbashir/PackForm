package models

// Order detail schema of the user table
type OrderDetail struct {
	CustomerCompany string `json:"CustomerCompany"`
	CustomerName     string `json:"CustomerName"`
	OrderName string `json:"OrderName"`
	CreatedAt string `json:"CreatedAt"`
	TotalAmount float64 `json:"TotalAmount"`
	DeliveredAmount float64 `json:"DeliveredAmount"`
	Total float64 `json:"total"`
}
