package router

import (
	"tushargoyal100/expensecalculatorapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/expense/{id}", controller.GetOneExpense).Methods("GET")
	router.HandleFunc("/expense", controller.CreateOneExpense).Methods("POST")
	router.HandleFunc("/expenses", controller.GetAllExpense).Methods("GET")
	router.HandleFunc("/expense/{id}", controller.DeleteOneExpense).Methods("DELETE")
	router.HandleFunc("/expenses", controller.DeleteAllExpenses).Methods("DELETE")
	return router
}
