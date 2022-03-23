package routes

import (
	//"log"

	"github.com/Kuppa/todo/handlers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	//log.Info("Router initiation")
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.CheckHealth).Methods("GET")

	router.HandleFunc("/createuser", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/create", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/taskslist", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/getbyuser", handlers.GetTaskByUser).Methods("GET")
	router.HandleFunc("/getbydate", handlers.GetByDate).Methods("GET")
	router.HandleFunc("/update/{id}", handlers.UpdateTask).Methods("PUT")

	return router
}
