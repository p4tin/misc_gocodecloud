// Go program to illustrate
// how to compare two maps
package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	read_mongo()

	map_1 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}
	map_1_1 := map[int]string{
		205: "Rohit",
		200: "Anita",

		203: "Suman",
		204: "Robin",
		201: "Neha",
	}
	map_2 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
		206: "Sumit",
	}

	map_3 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}
	map_4 := map[string]int{

		"Anita": 200,
		"Neha":  201,
		"Suman": 203,
		"Robin": 204,
		"Rohit": 205,
	}

	// Comparing maps
	// Using DeepEqual() function
	res0 := reflect.DeepEqual(map_1, map_1_1)
	res1 := reflect.DeepEqual(map_1, map_2)
	res2 := reflect.DeepEqual(map_1, map_3)
	res3 := reflect.DeepEqual(map_1, map_4)
	res4 := reflect.DeepEqual(map_2, map_3)
	res5 := reflect.DeepEqual(map_3, map_4)
	res6 := reflect.DeepEqual(map_4, map_4)
	res7 := reflect.DeepEqual(map_2, map_4)

	// Displaying result
	fmt.Println("Is Map 1 is equal to Map 1_1: ", res0)
	fmt.Println("Is Map 1 is equal to Map 2: ", res1)
	fmt.Println("Is Map 1 is equal to Map 3: ", res2)
	fmt.Println("Is Map 1 is equal to Map 4: ", res3)
	fmt.Println("Is Map 2 is equal to Map 3: ", res4)
	fmt.Println("Is Map 3 is equal to Map 3: ", res5)
	fmt.Println("Is Map 4 is equal to Map 4: ", res6)
	fmt.Println("Is Map 2 is equal to Map 4: ", res7)

}

func read_mongo() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("CheckoutService")
	collection := db.Collection("orders")

	result := make(map[string]interface{})
	err = collection.FindOne(context.TODO(), bson.M{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found documents: %+v\n", result)
}
