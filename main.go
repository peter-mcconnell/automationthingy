package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/peter-mcconnell/automationthingy/api"
	"github.com/peter-mcconnell/automationthingy/web"
)

func main() {
	web_port := 8080
	api_port := 8081

	web_server, err := web.NewServer(log.Default(), http.NewServeMux())
	if err != nil {
		log.Fatal(err)
	}

	api_server, err := api.NewServer(log.Default(), http.NewServeMux())
	if err != nil {
		log.Fatal(err)
	}

	// run api server
	go http.ListenAndServe(":"+strconv.Itoa(api_port), api_server)

	// run web server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(web_port), web_server))
}
