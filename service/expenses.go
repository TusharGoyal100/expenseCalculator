package service

import (
	"encoding/json"
	"expensecalculatorapi/connector"
	"expensecalculatorapi/model"
)

func AddExpense(expense model.Expense) error {

	if err := connector.InsertOne(expense); err != nil {
		return err
	}

	return nil
}

func GetAllExpense() ([]model.Expense, error) {
	var out []model.Expense
	expensesMap, err := connector.GetAllResult()
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(expensesMap)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func DeleteExpense(expenseId string) error {

	if err := connector.DeleteOne(expenseId); err != nil {
		return err
	}

	return nil
}

func DeleteAllExpenses() error {

	if err := connector.DeleteAllResult(); err != nil {
		return err
	}

	return nil
}
