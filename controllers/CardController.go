package controllers

import (
	"net/http"
	"BudgetSrv2/models"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

var db *gorm.DB


func getDB() *gorm.DB {
	return db
}

func SetDB(d *gorm.DB) {
	db = d
}

//func CardGetHandler(){
func CardsGetHandler(w http.ResponseWriter, r *http.Request){
	println("Method Get")

	holder := GetRepoHolder()

	cardRepo := holder.CardRepo

	cards := cardRepo.GetCards()

	enc := json.NewEncoder(w)
	enc.Encode(&cards)
}

//func CardGetHandler(){
func CardGetHandler(w http.ResponseWriter, r *http.Request){
	println("Method Get")

	vars := mux.Vars(r)
	id := vars["id"]
	holder := GetRepoHolder()

	cardRepo := holder.CardRepo

	card := cardRepo.GetCardByName(id)

	enc := json.NewEncoder(w)
	enc.Encode(&card)
}

func CardPutHandler(w http.ResponseWriter, r *http.Request){
	println("Method Put")

	dec := json.NewDecoder(r.Body)
	card := &models.Card{}
	dec.Decode(card)

	holder := GetRepoHolder()
	cardRepo := holder.CardRepo

	cardRepo.InsertCard(card)

	enc := json.NewEncoder(w)
	enc.Encode(card)
}

func CardPostHandler(w http.ResponseWriter, r *http.Request){
	println("Method Post")
	db = getDB()
	vars := mux.Vars(r)
	id := vars["id"]
	println(id)
	dec := json.NewDecoder(r.Body)
	card := &models.Card{}
	dec.Decode(card)

	holder := GetRepoHolder()
	cardRepo := holder.CardRepo

	cardRepo.UpdateCardByName(id, card)

	enc := json.NewEncoder(w)
	enc.Encode(card)
}

func CardDeleteHandler(w http.ResponseWriter, r *http.Request){
	println("Method Delete")
	db = getDB()
	vars := mux.Vars(r)
	id := vars["id"]

	holder := GetRepoHolder()
	cardRepo := holder.CardRepo
	cardRepo.DeleteCardByName(id)
}
