package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/status", getStatus)
	http.ListenAndServe(":8080", router)
}
