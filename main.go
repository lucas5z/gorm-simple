package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"githup.com/lucas5z/mux-fast/db"
	"githup.com/lucas5z/mux-fast/models"
	"githup.com/lucas5z/mux-fast/routes"
)

func main() {
	//donde se agrupa la url
	db.BDconnex()
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Task{})

	r := mux.NewRouter()
	//user
	r.HandleFunc("/user", routes.Get_user).Methods("GET")
	r.HandleFunc("/user/{id}", routes.Get_user_id).Methods("GET")
	r.HandleFunc("/user", routes.Post_user).Methods("POST")
	r.HandleFunc("/user/{id}", routes.Put_user).Methods("PUT")
	r.HandleFunc("/user/{id}", routes.Delete_user).Methods("DELETE")
	//task
	r.HandleFunc("/task", routes.Get_task).Methods("GET")
	r.HandleFunc("/task", routes.Post_task).Methods("POST")
	r.HandleFunc("/task/{id}", routes.Put_task).Methods("PUT")
	r.HandleFunc("/task/{id}", routes.Delete_task).Methods("DELETE")
	log.Println("corriendo en el puerto:8080...")

	http.ListenAndServe(":8080", r)
}
