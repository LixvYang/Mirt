// Package main provides 
package main

import (
	"github.com/lixvyang/mirt/framework"
)

func registerRouter(core *framework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	core.Get("foo", FooControllerHandler)
}
