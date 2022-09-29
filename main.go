package main

import (
	"fmt"
	"net/http"
	"tushargoyal100/expensecalculatorapi/router"
)

func main() {

	fmt.Println("Starting ExpenseCalculator App....")
	r := router.Router()

	fmt.Println("Server is getting ready")
	fmt.Println("Listining on port 4000")

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		panic("Problem starting the server")
	}
}
