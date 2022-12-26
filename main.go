package main

import (
	"net/http"

	"github.com/DavidGenZ/fazt-go-api/db"
	"github.com/DavidGenZ/fazt-go-api/models"
	"github.com/DavidGenZ/fazt-go-api/routes"
	"github.com/gorilla/mux"
)


func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router:= mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}