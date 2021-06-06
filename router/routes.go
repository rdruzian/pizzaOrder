package router

import (
	"net/http"
	"pizzaOrder/api"
)

func Init(){
	http.HandleFunc("/", api.Home)
	http.HandleFunc("/list", api.List)
	http.HandleFunc("/addFlavor", api.AddFlavor)
	http.HandleFunc("/newOrder", api.NewOrder)
	http.HandleFunc("/createUSer", api.CreateUser)
	http.HandleFunc("/login", api.Login)
}