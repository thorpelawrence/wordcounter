package main

import (
	"fmt"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	log.Println("listening on", addr)

	log.Fatal(http.ListenAndServe(addr, r))
}
