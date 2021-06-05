package api

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Method not supported")
		return
	}

	fmt.Fprintf(w, "Simple project!")
}