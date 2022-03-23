package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kuppa/todo/db"
	"github.com/Kuppa/todo/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var connection *mongo.Database

func init() {
	client := db.ConnectToDB()
	dbDetails := db.GetTasksDBDetails()

	connection = client.Database(dbDetails.Name)
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	log.Print("Welcome!!")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	task.Status = "Pending"
	task.CreatedDate = time.Now()

	json.NewDecoder(r.Body).Decode(&task)

	taskCollection := connection.Collection(db.GetTasksDBDetails().CollectionName)
	result, error := taskCollection.InsertOne(context.TODO(), task)
	if error != nil {
		log.Fatal("Error while creatring document")
	}
	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usr models.User

	json.NewDecoder(r.Body).Decode(&usr)

	userCollection := connection.Collection(db.GetUserDBDetails().CollectionName)
	result, error := userCollection.InsertOne(context.TODO(), usr)
	if error != nil {
		log.Fatal("Error while creating user")
	}

	json.NewEncoder(w).Encode(result)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []models.Task

	cursor, error := connection.Collection(db.GetTasksDBDetails().CollectionName).Find(context.TODO(), bson.M{})
	if error != nil {
		log.Fatal("Error while getting taks list")
	}

	if error = cursor.All(context.TODO(), &tasks); error != nil {
		panic(error)
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks []models.Task

	query := bson.M{"createdby": "User"}

	cursor, error := connection.Collection(db.GetTasksDBDetails().CollectionName).Find(context.TODO(), query)
	if error != nil {
		log.Fatal("Error while getting taks list")
	}

	fmt.Println("TTT::", cursor)

	if error = cursor.All(context.TODO(), &tasks); error != nil {
		panic(error)
	}
	fmt.Println("error::", error)
	fmt.Println("tasks::", tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks []models.Task

	groupStage := bson.D{{"$group", bson.D{{"_id", "$createddate"}}}}

	//group := bson.M{"$group": bson.M{"createddate": "$createddate", "count": bson.M{"$sum": 1}}}

	cursor, error := connection.Collection(db.GetTasksDBDetails().CollectionName).Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if error != nil {
		log.Fatal("Error while getting taks list")
	}

	fmt.Println("TTT::", cursor)

	if error = cursor.All(context.TODO(), &tasks); error != nil {
		panic(error)
	}
	fmt.Println("error::", error)
	fmt.Println("tasks::", tasks)
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var task models.Task
	task.Status = "Pending"

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&task)

	// prepare update model.
	update := bson.D{{"$set", task}}

	err := connection.Collection(db.GetTasksDBDetails().CollectionName).FindOneAndUpdate(context.TODO(), filter, update).Decode(&task)

	if err != nil {
		log.Fatal("Error while updating task")
	}
	json.NewEncoder(w).Encode(task)

}
