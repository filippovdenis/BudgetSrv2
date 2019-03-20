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
	db = getDB()
	var cards []models.Card

	db.Find(&cards)

	enc := json.NewEncoder(w)
	enc.Encode(&cards)
}

//func CardGetHandler(){
func CardGetHandler(w http.ResponseWriter, r *http.Request){
	println("Method Get")
	db = getDB()
	vars := mux.Vars(r)
	id := vars["id"]
	println(id)
	var card models.Card

	db.Where("card_name = $1", id).Find(&card)

	enc := json.NewEncoder(w)
	enc.Encode(&card)
}

func CardPutHandler(w http.ResponseWriter, r *http.Request){
	println("Method Put")
	db = getDB()
	dec := json.NewDecoder(r.Body)
	card := &models.Card{}
	dec.Decode(card)
	println(card.CardName)
	println(card.Description)
	db.Create(card)
}

func CardPostHandler(w http.ResponseWriter, r *http.Request){
	println("Method Post")
	db = getDB()
	vars := mux.Vars(r)
	id := vars["id"]
	println(id)
	dec := json.NewDecoder(r.Body)
	card := &models.Card{}
	db.Debug().Where("card_name = $1", id).Debug().Find(&card)
	if card.CardName == "" {
		return
	}

	dec.Decode(card)
	println(card.CardName)
	db.Debug().Save(card)
}

func CardDeleteHandler(w http.ResponseWriter, r *http.Request){
	println("Method Delete")
	db = getDB()
	vars := mux.Vars(r)
	id := vars["id"]
	println(id)
	db.Table("cards").Debug().Where("card_name= ?", id).Debug().Delete(&models.Card{})
}
