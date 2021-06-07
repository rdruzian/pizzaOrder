package sqlclient

import (
	"fmt"
	"pizzaOrder/customer"
	"pizzaOrder/database"
	"pizzaOrder/order"
	"pizzaOrder/pizza"
)

// SaveCustomer add a customer on database
func SaveCustomer(customer customer.Customer) error {
	db := database.ConnDb(true)
	defer db.Close()
	stmt, err := db.Prepare(`insert into customer (name, orders, favorite, login, pass) values (?,?,?,?,?)`)
	if err != nil {
		fmt.Println("Error to create save data on Customer table", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(customer.Name, customer.Orders, customer.Favorite, customer.Login, customer.Password)
	if err != nil {
		fmt.Println("Error to save customer")
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error to get a number of rows affected")
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error to get the last ID inserted")
		return err
	}
	err = saveAddress(lastID, customer.Address)
	if err != nil {
		fmt.Println("Error to insert ingredients")
		return err
	}

	fmt.Println("Customer saved successfully")
	fmt.Println("Rows affected", affected)

	return nil
}

// saveAddress save the address of the customer
func saveAddress(id int64, add customer.Address) error {
	var lastID int64

	db := database.ConnDb(true)
	defer db.Close()
	stmt, err := db.Prepare(`insert into address (type, publicPlace, number, complement) values (?,?,?,?)`)
	if err != nil {
		fmt.Println("Error to create save data on Customer table", err)
		return err
	}
	defer stmt.Close()
	switch add.Type{
	case customer.Street:
		res, err := stmt.Exec("Street", add.PublicPlace, add.Number, add.Complement)
		if err != nil {
			fmt.Println("Error to save address")
			return err
		}
		lastID, err = res.LastInsertId()
		if err != nil {
			fmt.Println("Error to get the last ID inserted")
			return err
		}
	case customer.Avenue:
		res, err := stmt.Exec("Avenue", add.PublicPlace, add.Number, add.Complement)
		if err != nil {
			fmt.Println("Error to save address")
			return err
		}
		lastID, err = res.LastInsertId()
		if err != nil {
			fmt.Println("Error to get the last ID inserted")
			return err
		}
	case customer.Road:
		res, err := stmt.Exec("Road", add.PublicPlace, add.Number, add.Complement)
		if err != nil {
			fmt.Println("Error to save address")
			return err
		}
		lastID, err = res.LastInsertId()
		if err != nil {
			fmt.Println("Error to get the last ID inserted")
			return err
		}
	default:
		res, err := stmt.Exec("Street", add.PublicPlace, add.Number, add.Complement)
		if err != nil {
			fmt.Println("Error to save address")
			return err
		}
		lastID, err = res.LastInsertId()
		if err != nil {
			fmt.Println("Error to get the last ID inserted")
			return err
		}
	}
	err = saveCustomerAddress(lastID, id)
	if err != nil {
		fmt.Println("Error to insert on customerAddress")
		return err
	}
	return err
}

func saveCustomerAddress(idAdd, id int64) error {
	db := database.ConnDb(true)
	defer db.Close()
	stmt, err := db.Prepare(`insert into customerAddress (customerID, addressID) values (?,?)`)
	if err != nil {
		fmt.Println("Error to create save data on Customer table", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, idAdd)
	if err != nil {
		fmt.Println("Error to save regitry")
		return err
	}
	return nil
}

// SaveFlavor add a new flavor on database
func SaveFlavor(flavor []pizza.Pizza) error {
	db := database.ConnDb(true)
	defer db.Close()
	stmt, err := db.Prepare(`insert into pizza (nameFlavor, price) values (?,?)`)
	if err != nil {
		fmt.Println("Error to create save data on Customer table", err)
		return err
	}
	defer stmt.Close()
	for f := range flavor {
		res, err := stmt.Exec(f.FlavorName, f.Price)
		if err != nil {
			fmt.Println("Error to save customer")
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			fmt.Println("Error to get a number of rows affected")
			return err
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			fmt.Println("Error to get the last ID inserted")
			return err
		}
		err = saveIngredients(lastID, f.Ingredients)
		if err != nil {
			fmt.Println("Error to insert ingredients")
			return err
		}
		fmt.Println("Customer saved successfully")
		fmt.Println("Rows affected", affected)
	}

	return nil
}

// saveIngredients save the ingredients of the new flavor
func saveIngredients(id int64, ingredients []string) error {
	db := database.ConnDb(true)
	defer db.Close()
	for i := range ingredients {
		stmt, err := db.Prepare( `insert into ingredients (pizzaId) values (?)`)
		if err != nil {
			fmt.Println("Error to create save data on Customer table", err)
			return err
		}
		defer stmt.Close()

		stmtSelect, err := db.Prepare( `select ingredientID from ingredients where nameIngredients = ?`)
		if err != nil {
			fmt.Println("Error to create save data on Customer table", err)
			return err
		}
		defer stmtSelect.Close()

		res, err := stmtSelect.Exec(i)
		if err != nil {
			fmt.Println("Error to get the name of ingredient")
			return err
		}
		_, err = stmt.Exec(res)
		if err != nil {
			fmt.Println("Error to get the name of ingredient")
			return err
		}
	}

	return nil
}

// SaveOrder registry the order on database
func SaveOrder(client customer.Customer, orders order.Order) error {

	return nil
}