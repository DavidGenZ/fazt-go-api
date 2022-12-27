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

	/* Users Routes */
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	/* Task Routes */

	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}