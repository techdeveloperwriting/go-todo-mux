package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBDetails struct {
	Name           string
	CollectionName string
}

func ConnectToDB() *mongo.Client {
	// TODO: Need to read mongo uril form config file
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Print("Error while connecting mongodb")
	}
	//log.Print("DBConnection Created", client)
	//log.Print("DBConnection Created &&&&&", &client)
	//log.Print("DBConnection Created ****", *client)
	return client
}

func GetTasksDBDetails() *DBDetails {
	db := &DBDetails{Name: "todo", CollectionName: "tasks"}
	return db // TODO: Need to read it from config file
}

func GetUserDBDetails() *DBDetails {
	db := &DBDetails{Name: "todo", CollectionName: "users"}
	return db // TODO: Need to read it from config file
}
