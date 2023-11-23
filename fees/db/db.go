package db

import (
	"context"
	"fees/students"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() error {

	clientDummy, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://root:RR3TuipRLnX3M5pY@fees.bfdmqbj.mongodb.net/?retryWrites=true&w=majority"))
	if err == nil {
		fmt.Println("Connected to DB")
	}

	client = clientDummy

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Cannot ping server")
		return err
	} else {
		fmt.Println("Can be Pinged")
	}

	return nil
}

func GetClient() *mongo.Client {
	return client
}

func QueryClass(client *mongo.Client, class int) ([]students.Student, error) {

	collectionName := "fees"
	dbName := "feedb"

	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		log.Fatal("No such collection")
	}

	filter := bson.M{"Class": class}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var result []students.Student

	for cursor.Next(context.Background()) {
		var res students.Student
		err := cursor.Decode(&res)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, res)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Query results:")
	// for _, res := range result {
	// 	fmt.Printf("%+v\n", res)
	// }

	return result, err
}

func QueryStudent(client *mongo.Client, id int) (students.Student, error) {
	collectionName := "fees"
	dbName := "feedb"

	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		log.Fatal("No such collection")
	}

	filter := bson.M{"SID": id}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var result students.Student

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Query results:")
	// for _, res := range result {
	// 	fmt.Printf("%+v\n", res)
	// }

	return result, err
}

func QueryPayFee(client *mongo.Client, id int, amount string) (students.Student, error) {
	collectionName := "fees"
	dbName := "feedb"

	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		log.Fatal("No such collection")
	}

	filter := bson.M{"SID": id}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var result students.Student

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	st := result.PayFee(amount)
	update := bson.M{
		"$set": bson.M{
			"RemFee":     st.Remfee,
			"MonthArray": st.Montharray,
		},
	}

	options := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	res := collection.FindOneAndUpdate(context.Background(), filter, update, options).Decode(&st)
	if res != nil {
		log.Fatal(res)
	}

	fmt.Printf("Updated document: %+v\n", st)
	return st, err
}

func QueryAddStudent(client *mongo.Client, student students.Student) (students.Student, error) {
	collectionName := "fees"
	dbName := "feedb"

	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		log.Fatal("No such collection")
	}

	_, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
		return students.Student{}, err
	}

	return student, nil
}

func EmptyCollection(client *mongo.Client) error {
	collectionName := "fees"
	dbName := "feedb"

	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		log.Fatal("No such collection")
	}

	filter := bson.M{}
	_, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
