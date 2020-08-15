package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/bremade/recify/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBError struct {
	message string
}

func (err DBError) Error() string {
	return err.message
}

type DB struct {
	dbClient *mongo.Client
	database *mongo.Database
	ctx      context.Context
}

func (db *DB) Open(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println("error creating mongodb client", err)
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("error connecting to mongodb", err)
		return err
	}

	db.dbClient = client
	db.ctx = ctx

	fmt.Println("mongodb opened")
	return nil
}

func (db *DB) Setup() error {
	database := db.dbClient.Database("recify")
	db.database = database

	collection := database.Collection("Test")

	test := model.Test{
		Id:   1,
		Text: "12431244123rfweffsasfd",
	}

	_, err := collection.InsertOne(
		db.ctx, test)
	if err != nil {
		fmt.Println("error while inserting: ", err)
		return err
	}
	return nil
}

func (db *DB) QueryTest(id int) (model.Test, error) {
	collection := db.database.Collection("Test")
	var result model.Test

	err := collection.FindOne(db.ctx, bson.D{
		{Key: "id", Value: id},
	}).Decode(&result)

	fmt.Println("result:", result)

	return result, err
}
