package controller

import (
	"encoding/json"
	"expensecalculatorapi/connector"
	"expensecalculatorapi/model"
	"expensecalculatorapi/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOneExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetOneExpense")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	expenseId := params["id"]
	if expenseId == "" {
		json.NewEncoder(w).Encode("please gave a valid expense id")
		return
	}
	fmt.Println(expenseId)

	expense, err := connector.GetOneResult(expenseId)
	if err != nil {
		json.NewEncoder(w).Encode("no such expense is present")
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func CreateOneExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside CreateOneExpense")
	w.Header().Set("Content-Type", "application/json")

	var input model.Expense

	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode("Please send proper body")
		return
	}

	if r.Form["ItemType"] == nil {
		json.NewEncoder(w).Encode("Please gave item type")
		return
	} else {
		input.ItemType = r.Form["ItemType"][0]
		if !input.IsItemType() {
			json.NewEncoder(w).Encode("Item type is not same")
			return
		}
	}

	if r.Form["Name"] == nil {
		json.NewEncoder(w).Encode("Please gave Name")
		return
	} else {
		input.Item.Name = r.Form["Name"][0]
	}

	if r.Form["Price"] == nil {
		json.NewEncoder(w).Encode("Please gave Price")
		return
	} else {
		val, err := strconv.Atoi(r.Form["Price"][0])
		if err != nil {
			json.NewEncoder(w).Encode("Please gave valid price")
			return
		}
		input.Item.Price = val
	}

	if err := service.AddExpense(input); err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	json.NewEncoder(w).Encode(&input)

}

func GetAllExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetAllExpense")
	w.Header().Set("Content-Type", "application/json")

	var out []model.Expense

	out, err := service.GetAllExpense()
	if err != nil {
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode(out)
}

func DeleteOneExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside DeleteOneExpense")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	expenseId := params["id"]

	if !primitive.IsValidObjectID(expenseId) {
		json.NewEncoder(w).Encode("please gave a valid id")
		return
	}

	if err := service.DeleteExpense(expenseId); err != nil {
		json.NewEncoder(w).Encode("no such id is present")
		return
	}

	json.NewEncoder(w).Encode("Deleted")

}

func DeleteAllExpenses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside DeleteOneExpense")
	w.Header().Set("Content-Type", "application/json")

	if err := service.DeleteAllExpenses(); err != nil {
		json.NewEncoder(w).Encode("internal server error")
		return
	}

	json.NewEncoder(w).Encode("Deleted all the documents")

}
