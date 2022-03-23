package main

import (
	"net/http"

	"github.com/Kuppa/todo/db"
	"github.com/Kuppa/todo/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	log1 := log.New()
	log1.Formatter = new(log.JSONFormatter)
	log.Info("To do App Initiated")

	db.ConnectToDB()

	router := routes.Routes()

	log1.Fatal(http.ListenAndServe(":8181", router))
}
