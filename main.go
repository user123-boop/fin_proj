package main

import (
	"log"
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/db"
	"github.com/user123-boop/fin_proj/pkg/server"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))

	var err_db error
	var err_server error

	err_db = db.Init("scheduler.db")
	if err_db != nil {
		log.Fatalf("Ошибка создания БД: %v", err_db)
	}
	db.Init("scheduler.db")
	defer db.Close()

	err_server = server.Run()
	if err_server != nil {
		log.Println("server run fail", err_server)
	}
}
