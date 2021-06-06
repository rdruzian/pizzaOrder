package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnDb(useDbName bool) *sql.DB {
	var (
		db *sql.DB
		err error
	)

	if useDbName {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s sslmode=disable", host, port, user, password)
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Successfully connected!")
	return db
}
