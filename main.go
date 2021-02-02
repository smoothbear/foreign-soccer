package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"smooth-bear.live/lib/database"
	"smooth-bear.live/lib/database/access"
	"smooth-bear.live/lib/handler"
)

func main() {
	db, err := database.ConnectToMysql()

	if err != nil {
		log.Fatalf("db connect fail, error: %v", err)
	}

	database.Migrate(db)

	accessManage, err := database.NewAccessorManage(access.Default(db))
	if err != nil {
		log.Fatalf("db accessor create fail, error: %v", err)
	}

	// Router Initialize
	router := mux.NewRouter()
	defaultHandler := handler.NewDefault(&accessManage)

	router.HandleFunc("/user", defaultHandler.CreateNewUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
