package main

import (
	"BudgetSrv2/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/jinzhu/gorm"

	"BudgetSrv2/PgRepo"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"os"
	"encoding/json"
)

type Configuration struct {
	DBType             string
	ConnectionString   string
}


func initDB(db *gorm.DB) {
	db.AutoMigrate(&PgRepo.PgCard{})
	cardRepo := PgRepo.PgCardRepo{Db: db}
	holder := controllers.RepoHolder{CardRepo: cardRepo}
	controllers.SetRepoHolder(&holder)
}

func main() {
	//connStr := "user=postgres password=tamplier dbname=Budget sslmode=disable"
	var configuration Configuration
	filename := "config/config.json"
	file, errF := os.Open(filename)
	if errF != nil {
		log.Fatal(errF)
	}
	decoder := json.NewDecoder(file)
	errF = decoder.Decode(&configuration)
	if errF != nil {
		log.Fatal(errF)
	}


	db, err := gorm.Open(configuration.DBType, configuration.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}


	initDB(db)
	defer db.Close()

	holder := controllers.GetRepoHolder()

	cards := holder.CardRepo.GetCards()

	for _, card := range cards {
		println(card.CardName)
	}

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
