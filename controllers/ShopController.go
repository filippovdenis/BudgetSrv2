package controllers

import (
	"net/http"
	"encoding/json"
	"BudgetSrv2/models"
)

//func CardGetHandler(){
func ShopsGetHandler(w http.ResponseWriter, r *http.Request){
	println("Method Get")
	db = getDB()
	var shops []models.Shop

	db.Find(&shops)

	enc := json.NewEncoder(w)
	enc.Encode(&shops)
}
