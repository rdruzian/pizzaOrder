package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"pizzaOrder/api/types"
	"pizzaOrder/customer"
	"pizzaOrder/database/sqlclient"
	"pizzaOrder/pizza"
	"strconv"
	"strings"
)

// loadJson load a json file and avoid to repeat code
func loadJson(file string) []byte {
	userFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error to open customer json %v", err)
		return nil
	}
	defer userFile.Close()

	fileOpen, err := ioutil.ReadAll(userFile)
	if err != nil {
		fmt.Println("Error to read json %v", err)
		return nil
	}

	return fileOpen
}

// saveJson save a json file and avoid to repeat code
func saveJson(file string, v []byte) error {
	err := ioutil.WriteFile(file, v, 0777)
	if err != nil {
		fmt.Println("Error to write json file %v", err)
		return err
	}

	return err
}

// Home receive the request to a home page
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

	pizzas, err := listFlavors(types.Flavor)
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
	menu := loadJson(types.Flavor)

	if err = newFlavor(id, name, ingredient, price, menu); err != nil {
		fmt.Println("Error to save a new flavor!")
	}
}

// NewOrder receive the request for a new pizza order
func NewOrder(w http.ResponseWriter, r *http.Request) {}

// CreateUser receive the request to add an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Method not supported")
		return
	}
	// get values to fill the customer fields
	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	confirm := r.FormValue("confirm")

	if pass != confirm {
		fmt.Println("Password and confirm password must be equals")
		return
	}

	typeString := r.FormValue("type")
	// get address values
	typeAddress, err := strconv.ParseInt(typeString,10, 64)
	if err != nil {
		fmt.Println("Error to get type of address")
		return
	}
	pPlace := r.FormValue("publicPlace")
	num := r.FormValue("number")
	number, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		fmt.Println("Error to get number of house")
		return
	}
	complement := r.FormValue("comp")

	userAddress := customer.Address{
		Type: typeAddress,
		PublicPlace: pPlace,
		Number: number,
		Complement: complement,
	}

	user := customer.Customer{
		Name: name,
		Address: userAddress,
		Login: email,
		Password: pass,
	}


	err = saveUser(user)
	if err != nil {
		fmt.Println("Error to save a new user")
		return
	}
}

// Login receive the request to make login
func Login(w http.ResponseWriter, r *http.Request) {}

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

	err = saveJson(types.Flavor, menu)
	if err != nil {
		fmt.Println("Error to write json file %v", err)
		return err
	}
	err = sqlclient.SaveFlavor(pizzas)
	if err != nil {
		fmt.Println("Erro to save ondatabase")
		return err
	}

	fmt.Println("Flavor added with success!")
	return nil
}

//listFlavors list all the flavor on database
func listFlavors(filePath string) ([]pizza.Pizza, error) {
	var pizzas []pizza.Pizza
	menu := loadJson(filePath)
	err := json.Unmarshal(menu, &pizzas)
	if err != nil {
		fmt.Println("Error to unmarshal json %v", err)
		return pizzas, err
	}
	fmt.Println("Success to list flavors!")
	return pizzas, err
}

// saveUSer save a new user on a json file
func saveUser(user customer.Customer) error {
	var userJson customer.Customer
	file := loadJson(types.User)

	err := json.Unmarshal(file, &userJson)
	if err != nil {
		fmt.Println("error to load json file")
		return err
	}

	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error to marshal json %v", err)
		return err
	}

	err = saveJson(types.User, u)
	if err != nil {
		fmt.Println("saveUser error to save file")
	}
	err = sqlclient.SaveCustomer(user)
	if err != nil {
		fmt.Println("saveUser Error to save on database %v", err)
		return err
	}

	fmt.Println("User save succesfully!")
	return err
}