package main

import (
	"net/http"
	"./controller" 
	"./database" 
)

func main() {
	db := database.InitDB("storage.db")
	database.CreateTable(db)

	http.HandleFunc("/login", controller.LoginHandler(db))
	http.HandleFunc("/register", controller.RegisterHandler(db))

	http.ListenAndServe(":8080", nil)
}
