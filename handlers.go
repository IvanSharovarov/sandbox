package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
		checkErr(err)

		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var u User
	err = json.NewDecoder(r.Body).Decode(&u)
	checkErr(err)

	db.QueryRow(
		"INSERT INTO users(name, email, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5) returning id",
		u.Name,
		u.Email,
		u.Password,
		time.Now(),
		time.Now(),
	).Scan(&u.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var u User
	ID := mux.Vars(r)["id"]

	err = db.QueryRow("SELECT * FROM users WHERE id=$1", ID).
		Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
	checkErr(err)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var u User
	pID := mux.Vars(r)["id"]
	err = json.NewDecoder(r.Body).Decode(&u)
	checkErr(err)

	stmt, err := db.Prepare("update users set name=$1, email=$2, password=$3, updated_at=$4 where id=$5")
	checkErr(err)

	_, err = stmt.Exec(
		u.Name,
		u.Email,
		u.Password,
		time.Now(),
		pID,
	)
	checkErr(err)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	pID := mux.Vars(r)["id"]

	stmt, err := db.Prepare("delete from users where id=$1")
	checkErr(err)
	res, err := stmt.Exec(pID)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	w.WriteHeader(http.StatusOK)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
