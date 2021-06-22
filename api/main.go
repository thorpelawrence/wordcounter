package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/url/*url", URLHandler)

	var r http.Handler
	r = handlers.CombinedLoggingHandler(os.Stdout, router)
	r = handlers.CompressHandler(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
