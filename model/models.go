package model

import (
	"expensecalculatorapi/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemType string             `json:"item_type" bson:"item_type"`
	Item     Item               `json:"item" bson:"item"`
}

type Item struct {
	Name  string    `json:"name" bson:"name"`
	Price int       `json:"price" bson:"price"`
	Time  time.Time `json:"time" bson:"time"`
}

func (expense *Expense) IsItemType() bool {
	if expense.ItemType == constants.Grocery || expense.ItemType == constants.Food || expense.ItemType == constants.Gym || expense.ItemType == constants.Miscellaneous || expense.ItemType == constants.Rent || expense.ItemType == constants.Travel {
		return true
	}

	return false
}
