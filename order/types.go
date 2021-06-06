package order

import "time"

const (
	WaitingPayment = iota
	Payed
	Doing
	WaitingDelivery
	Delivering
	Deliveried


type Order struct {
	Id int64 `json:"orderId"`
	CustomerID int64 `json:"customer"`
	Flavor []flavor `json:"flavor"`
	Date time.Time `json:"dateTime"`
	TotalPrice float64 `json:"total"`
	Status int64 `json:"status"`
}

type flavor struct {
	FlavorId int64 `json:"flavorId"`
	PizzaEdge string `json:"edge"`
	Units int64 `json:"units"`
	Details string `json:"details"`
	Price float64 `json:"price"`
}