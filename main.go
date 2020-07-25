package main

import (
	"log"
	"net/http"
)

func main() {
	r := CreateRouter()
	log.Fatal(http.ListenAndServe(":8090", r))
}
