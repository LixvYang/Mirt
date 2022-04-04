package main

import (
	"github.com/lixvyang/mirt/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: "localhost:8080",
	}
	server.ListenAndServe()
}