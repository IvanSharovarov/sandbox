package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var users []User

func main() {
	u := User{
		ID:        "1",
		Name:      "Ivan",
		Email:     "ivan@mail.com",
		Password:  "password",
		CreatedAt: "31.10.1995",
		UpdatedAt: "15.07.2020",
	}

	users = append(users, u)

	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	u.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, u := range users {
		if u.ID == p["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(u)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	var updatedUser User
	for i, u := range users {
		if u.ID == p["id"] {
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			updatedUser.ID = p["id"]
			users[i] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := mux.Vars(r)["id"]
	for i, u := range users {
		if u.ID == ID {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
