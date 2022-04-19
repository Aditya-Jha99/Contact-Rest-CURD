package controllers

import (
	"contact/database"
	"contact/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllContact(w http.ResponseWriter, r *http.Request) {
	var contacts []entity.Contact
	database.Connector.Find(&contacts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}

func GetContactByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var contact entity.Contact
	database.Connector.First(&contact, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var contact entity.Contact
	json.Unmarshal(requestBody, &contact)

	database.Connector.Create(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func UpdateContactByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var contact entity.Contact
	json.Unmarshal(requestBody, &contact)
	database.Connector.Save(&contact)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

func DeletContactByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var contact entity.Contact
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&contact)
	w.WriteHeader(http.StatusNoContent)
}
