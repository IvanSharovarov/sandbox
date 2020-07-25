package main

import "fmt"

const (
	DB_NAME     = "ghost"
	DB_USER     = "postgres"
	DB_PASSWORD = "123"
)

var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
