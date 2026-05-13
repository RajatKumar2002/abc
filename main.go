package main

import (
	"abc/data"
	"abc/routes"
	"log"
	"net/http"
)

func main() {
	data.InitDB()
	router := routes.RegisterRoutes()
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
