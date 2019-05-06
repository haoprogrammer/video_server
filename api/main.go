package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func main() {
	r := RegisterHandlers()
	//监听端口
	http.ListenAndServe(":8000", r)
}



func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

