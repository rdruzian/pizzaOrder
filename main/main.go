package main

import (
	"fmt"
	"log"
	"net/http"
	"pizzaOrder/database/sqlclient"
	"pizzaOrder/router"
)

func main() {
	// The routes initialization must be before initialize the web server
	router.Init()

	//err := sqlclient.CreateDataBase()
	//if err != nil {
	//	fmt.Println("Problem to create database")
	//	return
	//}

	err := sqlclient.CreateCustomer()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	err = sqlclient.CreateAddress()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	err = sqlclient.CreateOrder()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	err = sqlclient.CreatePizza()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	err = sqlclient.CreateCustomerAddress()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	err = sqlclient.CreateOrderCustomer()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	err = sqlclient.CreateIngredients()
	if err != nil {
		fmt.Println("Problem to create table")
		return
	}

	fmt.Println("Server is running on 9090 port!")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}