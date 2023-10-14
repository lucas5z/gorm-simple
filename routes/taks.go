package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"githup.com/lucas5z/mux-fast/db"
	"githup.com/lucas5z/mux-fast/models"
)

func Get_task(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func Post_task(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	db.DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func Put_task(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	us := mux.Vars(r)
	db.DB.First(&task, us["id"])
	if task.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	db.DB.Save(&task)
	w.WriteHeader(http.StatusNoContent)

}
func Delete_task(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	us := mux.Vars(r)
	db.DB.First(&task, us["id"])
	if task.Id == 0 {
		w.WriteHeader(500)
		return
	}
	db.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)

}
