package sqlclient

import (
	"fmt"
	"pizzaOrder/database"
)

func CreateDataBase() error {
	db := database.ConnDb(false)
	defer db.Close()
	_, err := db.Exec(`create database pizza`)
	if err != nil {
		fmt.Println("Error to create database", err)
		return err
	}
	fmt.Println("Database created successfully")

	return nil
}

func CreateCustomer() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists customer(
		customerId SERIAL not null primary key,
		name varchar(150) not null,
		favorite integer,
		login varchar(50) not null,
		pass varchar(25) not null
	);`)
	if err != nil {
		fmt.Println("Error to create table customer", err)
		return err
	}
	fmt.Println("Table customer created successfully!")

	return nil
}

func CreateAddress() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists address(
		  addressID SERIAL not null primary key,
		  type varchar(15) not null,
		  publicPlace varchar(100) not null,
		  number integer not null,
		  complement varchar(100)
	);`)
	if err != nil {
		fmt.Println("Error to create table address", err)
		return err
	}
	fmt.Println("Table address created successfully!")

	return nil
}

func CreateOrder() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists orders(
		orderID SERIAL not null primary key,
		customerID integer not null,
		flavor integer not null,
		quantity integer not null,
		details varchar(50),
		price decimal(4,2) not null,
		edge varchar(50),
		dateTime date not null,
		status integer not null,
		totalPrice decimal(5,2) not null
	);`)
	if err != nil {
		fmt.Println("Error to create table order", err)
		return err
	}
	fmt.Println("Table order created successfully!")

	return nil
}

func CreatePizza() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists pizza(
		pizzaID SERIAL not null primary key,
		nameFlavor varchar(50) not null,
		price decimal(4,2) not null
	);`)
	if err != nil {
		fmt.Println("Error to create table pizza", err)
		return err
	}
	fmt.Println("Table pizza created successfully!")

	return nil
}

func CreateCustomerAddress() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists customerAddress(
		customerID integer not null primary key references customer,
		addressID integer not null references address
	);`)
	if err != nil {
		fmt.Println("Error to create table customer/address", err)
		return err
	}
	fmt.Println("Table customer/address created successfully!")

	return nil
}

func CreateOrderCustomer() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists orderCustomer(
		cusotmerID integer not null primary key references customer,
		orderID integer not null references orders,
		total decimal(5,2) not null
	);`)
	if err != nil {
		fmt.Println("Error to create table customer/order", err)
		return err
	}
	fmt.Println("Table customer/order created successfully!")

	return nil
}

func CreateIngredients() error {
	db := database.ConnDb(true)
	defer db.Close()
	_, err := db.Exec(`create table if not exists ingredients(
		ingredientID SERIAL not null primary key,
		nameIngredients varchar(50) not null,
		pizzaID integer not null references pizza
	);`)
	if err != nil {
		fmt.Println("Error to create table ingredients", err)
		return err
	}
	fmt.Println("Table ingredients created successfully!")

	return nil
}