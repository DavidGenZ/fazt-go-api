package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DavidGenZ/fazt-go-api/db"
	"github.com/DavidGenZ/fazt-go-api/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("User not Found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
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
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("User not Found"))
		return
	}

	db.DB.Delete(&user) //cambiar de estado
	//db.DB.Unscoped().Delete(&user) //eliminar permanentemente
	w.WriteHeader(http.StatusOK)
}
