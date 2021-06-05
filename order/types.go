package order

import "time"

type Order struct {
	Id int64 `json:"orderId"`
	CustomerID int64 `json:"customer"`
	Flavor []flavor `json:"flavor"`
	Date time.Time `json:"dateTime"`
	TotalPrice float64 `json:"total"`
}

type flavor struct {
	FlavorId int64 `json:"flavorId"`
	PizzaEdge string `json:"edge"`
	Units int64 `json:"units"`
	Details string `json:"details"`
	Price float64 `json:"price"`
}