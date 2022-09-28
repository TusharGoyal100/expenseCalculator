package connector

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionUri string = "mongodb+srv://tushargoyal100:q5oMWlP06YpWZBXP@cluster0.omrstuo.mongodb.net/?retryWrites=true&w=majority"

var dbName = "ExpensesCollector"
var collectionName = "Expense"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionUri)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	//collection instance

	fmt.Println("Collection instance is ready")
}

func UpdateOne(expenseId string, expense interface{}) error {

	Id, _ := primitive.ObjectIDFromHex(expenseId)
	filter := bson.M{"_id": Id}
	_, err := collection.UpdateOne(context.Background(), filter, expense)

	if err != nil {
		return err
	}
	fmt.Println("Successfully updated One expense")
	return nil
}

func InsertOne(expense interface{}) error {
	inserted, err := collection.InsertOne(context.Background(), expense)

	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted one expense with ID:", inserted.InsertedID)
	return nil
}

func DeleteOne(id string) error {
	expenseId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": expenseId}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func GetOneResult(id string) (interface{}, error) {

	expenseId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": expenseId}

	res := collection.FindOne(context.Background(), filter)

	if res.Err() != nil {
		return nil, res.Err()
	}

	var expense bson.M
	res.Decode(&expense)

	return expense, nil
}

func GetAllResult() ([]primitive.M, error) {

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var expenses []primitive.M

	for cur.Next(context.Background()) {
		var expense bson.M
		if err := cur.Decode(&expense); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	defer cur.Close(context.Background())

	return expenses, nil
}

func DeleteAllResult() error {
	res, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return err
	}
	fmt.Printf("Deleted the %d documents\n", res.DeletedCount)
	return nil
}
