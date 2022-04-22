package main

import (
	"contact/controllers"
	"contact/database"
	"contact/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" 
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllContact).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetContactByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateContactByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletContactByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3000",
			User:       "root",
			Password:   "root",
			DB:         "contacts",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Contact{})
}
