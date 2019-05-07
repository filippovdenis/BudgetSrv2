package controllers

import (
	"net/http"
	"encoding/json"
	"BudgetSrv2/models"
	"strconv"
)

//func CardGetHandler(){
func ShopsGetHandler(w http.ResponseWriter, r *http.Request){
	println("Method Get")
	db = getDB()
	from := -1
	to := -1

	if len(r.URL.RawQuery) > 0 {
		str := r.URL.Query().Get("From")
		f, fErr := strconv.Atoi(str)
		if fErr == nil {
			from = f
		}
		str = r.URL.Query().Get("To")
		t, tErr := strconv.Atoi(str)
		if tErr == nil {
			to = t
		}
	}
	println(from)
	println(to)
	var shops []models.Shop

	db.Offset(from).Limit(to).Find(&shops)

	enc := json.NewEncoder(w)
	enc.Encode(&shops)
}
