package migrations

import (
	"gorm.io/gorm"
	"pizzaOrder/customer"
	"pizzaOrder/order"
	"pizzaOrder/pizza"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(customer.Customer{})
	db.AutoMigrate(order.Order{})
	db.AutoMigrate(pizza.Pizza{})
}

