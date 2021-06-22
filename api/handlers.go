package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func URLHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url := strings.TrimPrefix(ps.ByName("url"), "/")
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var parser Parser

	switch res.Header.Get("Content-Type") {
	case "text/html":
		parser, err = NewHTMLParser(res.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "text/plain":
		fallthrough
	default:
		parser, err = NewPlainTextParser(res.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	wordCounts := parser.GetWordCounts()

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(wordCounts); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
