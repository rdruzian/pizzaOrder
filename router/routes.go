package router

import (
	"net/http"
	"pizzaOrder/api"
)

func Init(){
	http.HandleFunc("/", api.Home)
}