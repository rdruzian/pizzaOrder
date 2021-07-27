package migrations

import (
	"gorm.io/gorm"
	"pizzaOrder/customer"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(customer.Customer{})
	//db.AutoMigrate(order.Order{})
	//db.AutoMigrate(pizza.Pizza{})
}