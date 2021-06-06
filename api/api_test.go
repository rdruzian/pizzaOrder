package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pizzaOrder/pizza"
	"testing"
)

var expectedFlavors = []pizza.Pizza{
	{
		ID:          0,
		FlavorName:  "Portuguesa",
		Ingredients: []string{},
		Price:       26.9,
	},
	{
		ID:          1,
		FlavorName:  "Americana",
		Ingredients: []string{},
		Price:       26.9,
	},
	{
		ID:          2,
		FlavorName:  "Palmito",
		Ingredients: []string{},
		Price:       26.9,
	},
	{
		ID:         10,
		FlavorName: "Calabresa",
		Ingredients: []string{
			"Calabresa",
			"Muçarela",
			"Oregano",
			"Azeitona",
			"Molho de tomate",
			" Cebola",
		},
		Price: 26.84,
	},
}

func TestList(t *testing.T) {
	filepath, err := filepath.Abs("test/testdata.json")
	if err != nil {
		t.Error(err.Error())
	}

	dataTest, err := listFlavors(filepath)
	if err != nil {
		t.Error(err.Error())
	}
	dataLen := len(dataTest)
	if dataLen == 4 {
		fmt.Println("Quantity is correct: list %v expected 4", dataLen)
	}

	for dt := range dataTest {
		for expected := range expectedFlavors {
			if expected == dt {
				fmt.Println("Error to compare the values. Expected %v Read %v", expected, dt)
				continue
			}
		}
	}
}

func TestNewFlavor (t *testing.T) {
	testData, err := os.Open("test/testdata.json")
	if err != nil {
		t.Error(err.Error())
	}
	defer testData.Close()

	menu, err := ioutil.ReadAll(testData)
	if err != nil {
		t.Error(err.Error())
	}

	var id int64 = 100
	name := "Teste"
	ingredient := []string{"muçarela", "tomate", "cebola"}
	price := 23.70

	if err = newFlavor(id, name, ingredient, price, menu); err != nil {
		t.Error(err.Error())
	}
}