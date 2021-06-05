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

//List receive the request to list all the flavors
func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Method not supported")
		return
	}

	pizzas, err := listFlavors("pizza.json")
	if err != nil {
		fmt.Println("Error to list flavors")
	}

	fmt.Fprintf(w, "Pizza Menu %v", pizzas)
}

// AddFlavor receive the request to save a new flavor on database
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
	ingredient := strings.Split(ingredients, ",")

	stringPrice := r.FormValue("price")
	price, err := strconv.ParseFloat(stringPrice,10)
	if err != nil {
		fmt.Println("Error to get PRICE value %v", err)
		return
	}

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

	if err = newFlavor(id, name, ingredient, price, menu); err != nil {
		fmt.Println("Error to save a new flavor!")
	}
}

// newFlavor add a new flavor on database
// TODO: verify to not add a duplicated flavor
func newFlavor(id int64, name string, ingredients []string, price float64, menu []byte) error {
	var pizzas []pizza.Pizza

	err := json.Unmarshal(menu, &pizzas)
	if err != nil {
		fmt.Println("Error to unmarshal json %v", err)
		return err
	}

	pizzas = append(pizzas, pizza.Pizza{
		ID: id,
		FlavorName: name,
		Ingredients: ingredients,
		Price: price,
	})

	menu, err = json.Marshal(pizzas)
	if err != nil {
		fmt.Println("Error to marshal json %v", err)
		return err
	}

	err = ioutil.WriteFile("pizza.json", menu, 0777)
	if err != nil {
		fmt.Println("Error to write json file %v", err)
		return err
	}

	fmt.Println("Flavor added with success!")
	return nil
}

//listFlavors list all the flavor on database
func listFlavors(filePath string) ([]pizza.Pizza, error) {
	var pizzas []pizza.Pizza
	menuFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error to open json %v", err)
		return pizzas, err
	}
	defer menuFile.Close()

	menu, err := ioutil.ReadAll(menuFile)
	if err != nil {
		fmt.Println("Error to read json %v", err)
		return pizzas, err
	}

	err = json.Unmarshal(menu, &pizzas)
	if err != nil {
		fmt.Println("Error to unmarshal json %v", err)
		return pizzas, err
	}

	fmt.Println("Success to list flavors!")

	return pizzas, err
}