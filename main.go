package main

import (
	"log"
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/db"
	"github.com/user123-boop/fin_proj/pkg/server"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))

	//err := db.Init("scheduler.db")
	//if err != nil {
	//	log.Fatalf("Ошибка создания БД: %v", err)
	//}
	db.Init("scheduler.db")
	defer db.Close()

	err := server.Run()
	if err != nil {
		log.Println("server run fail", err)
	}
}
