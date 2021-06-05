package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"pizzaOrder/pizza"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Method not supported")
		return
	}

	fmt.Fprintf(w, "Pizza Project!")
}

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Method not supported")
		return
	}

	var pizzas []pizza.Pizza

	menuFile, err := os.Open("pizza.json")
	if err != nil {
		fmt.Println("Error to open json %v", err)
		return
	}

	defer menuFile.Close()

	menu, err := ioutil.ReadAll(menuFile)
	if err != nil {
		fmt.Println("Error to read json %v", err)
		return
	}
	err = json.Unmarshal(menu, &pizzas)
	if err != nil {
		fmt.Println("Error to unmarshal json %v", err)
		return
	}

	fmt.Fprintf(w, "Pizza Menu %v", pizzas)
}

func AddFlavor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Method not supported")
		return
	}

	stringId := r.FormValue("id")
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		fmt.Println("Error to get ID value %v", err)
		return
	}

	name := r.FormValue("flavorName")
	ingredients := r.FormValue("ingredients")
	stringPrice := r.FormValue("price")
	price, err := strconv.ParseFloat(stringPrice,10)
	if err != nil {
		fmt.Println("Error to get PRICE value %v", err)
		return
	}

	ingredient := strings.Split(ingredients, ",")

	menuFile, err := os.Open("pizza.json")
	if err != nil {
		fmt.Println("Error to open json %v", err)
		return
	}
	defer menuFile.Close()

	var pizzas []pizza.Pizza

	menu, err := ioutil.ReadAll(menuFile)
	if err != nil {
		fmt.Println("Error to read json %v", err)
		return
	}
	err = json.Unmarshal(menu, &pizzas)
	if err != nil {
		fmt.Println("Error to unmarshal json %v", err)
		return
	}

	pizzas = append(pizzas, pizza.Pizza{
		ID: id,
		FlavorName: name,
		Ingredients: ingredient,
		Price: price,
	})

	menu, err = json.Marshal(pizzas)
	if err != nil {
		fmt.Println("Error to marshal json %v", err)
		return
	}

	err = ioutil.WriteFile("pizza.json", menu, 0777)
	if err != nil {
		fmt.Println("Error to write json file %v", err)
		return
	}
}