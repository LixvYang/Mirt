package framework

import (
	"net/http"
	"log"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
	log.Println("core.serverHTTP")
	ctx := NewContext(request, response)

	// TODO
	router := c.router["foo"]
	if router == nil {
		return
	}

	log.Println("core.router")
	router(ctx)
}