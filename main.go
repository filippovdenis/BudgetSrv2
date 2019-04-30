package main

import (
	"BudgetSrv2/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=tamplier dbname=Budget sslmode=disable"

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	controllers.SetDB(db)

	router := mux.NewRouter()
	router.HandleFunc("/card", controllers.CardsGetHandler).Methods("GET")
	router.HandleFunc("/card/{id}", controllers.CardGetHandler).Methods("GET")
	router.HandleFunc("/card/{id}", controllers.CardPostHandler).Methods("POST")
	router.HandleFunc("/card", controllers.CardPutHandler).Methods("PUT")
	router.HandleFunc("/card/{id}", controllers.CardDeleteHandler).Methods("DELETE")

	router.HandleFunc("/shop", controllers.ShopsGetHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
	println("hello")
}
