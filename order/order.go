package order

import (
	"gorm.io/gorm"
	"time"
)

const (
	WaitingPayment = iota
	Payed
	Doing
	WaitingDelivery
	Delivering
	Deliveried
)

type Order struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CustomerID int64 `json:"customer"`
	Flavor []flavor `json:"flavor"`
	Date time.Time `json:"dateTime"`
	TotalPrice float64 `json:"total"`
	Status int64 `json:"status"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type flavor struct {
	FlavorId int64 `json:"flavorId"`
	PizzaEdge string `json:"edge"`
	Units int64 `json:"units"`
	Details string `json:"details"`
	Price float64 `json:"price"`
}