package server

import (
	"fmt"
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/api"
)

func Run() error {

	port := ":7540"
	fmt.Println("Start server! port", port)

	api.Init()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}
