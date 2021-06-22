package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	contentType := res.Header.Get("Content-Type")
	var parser Parser
	switch contentType {
	case "text/html":
		parser, err = NewHTMLParser(res.Body)
		if err != nil {
			log.Println(err)
		}
	case "text/plain":
		fallthrough
	default:
		parser, err = NewPlainTextParser(res.Body)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(parser.GetWordCounts())
}
