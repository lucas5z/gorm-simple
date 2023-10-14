package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"githup.com/lucas5z/mux-fast/db"
	"githup.com/lucas5z/mux-fast/models"
)

func Get_user(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func Get_user_id(w http.ResponseWriter, r *http.Request) {
	var user models.User
	us := mux.Vars(r)
	db.DB.First(&user, us["id"])
	if user.Id == 0 {
		w.WriteHeader(500)
		return
	}
	//db.Model(&usuario).Association("Tasks").Find(&tareas)
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&user)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func Post_user(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	db.DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func Put_user(w http.ResponseWriter, r *http.Request) {
	var user models.User
	us := mux.Vars(r)
	db.DB.First(&user, us["id"])
	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	db.DB.Save(&user)
	w.WriteHeader(http.StatusNoContent)

}
func Delete_user(w http.ResponseWriter, r *http.Request) {
	var user models.User

	us := mux.Vars(r)
	db.DB.First(&user, us["id"])
	if user.Id == 0 {
		w.WriteHeader(500)
		return
	}
	db.DB.Delete(&user)
	w.WriteHeader(http.StatusNoContent)

}
