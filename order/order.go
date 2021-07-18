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
	Flavor []Flavor `json:"flavor"`
	Date time.Time `json:"dateTime"`
	TotalPrice float64 `json:"total"`
	Status int64 `json:"status"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Flavor struct {
	FlavorId string `json:"flavorId"`
	PizzaEdge string `json:"edge"`
	Units string `json:"units"`
	Details string `json:"details"`
	Price string `json:"price"`
}