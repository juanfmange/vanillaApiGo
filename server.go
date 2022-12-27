package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	Age       int64  `json:"age"`
	Message   string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/", GetUser).Methods("GET")
	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{"Juanito", "Mange", 25, "retomando go"}
	json.NewEncoder(w).Encode(user)
}
