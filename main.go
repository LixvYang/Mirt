package main

import (
	"github.com/lixvyang/mirt/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr: "localhost:8080",
	}
	server.ListenAndServe()
}