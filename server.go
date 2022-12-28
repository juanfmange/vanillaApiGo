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
	r.HandleFunc("/user/post/", PostUser).Methods("POST")
	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{"Juanito", "Mange", 25, "retomando go"}
	json.NewEncoder(w).Encode(user)
	log.Println("GET /user")

}

func PostUser(w http.ResponseWriter, r *http.Request) {
	newUser := User{"paco", "mange", 30, "POST en go"}
	json.NewEncoder(w).Encode(newUser)
	log.Println("POST /newUser")

}
