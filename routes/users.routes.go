package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DavidGenZ/fazt-go-api/db"
	"github.com/DavidGenZ/fazt-go-api/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get one user"))
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //status 400
		w.Write([]byte(err.Error()))         //msg error
	}

	json.NewEncoder(w).Encode((&user))
}
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
