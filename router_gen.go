// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "hertz-m/biz/router"
)

// register registers all routers.
func Register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
