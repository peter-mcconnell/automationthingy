package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/peter-mcconnell/automationthingy/api"
	"github.com/peter-mcconnell/automationthingy/logger"
	"github.com/peter-mcconnell/automationthingy/web"
)

const (
	web_port = 8080
	api_port = 8081
)

type flagStruct struct {
	ConfigPrint *bool
}

func flags() flagStruct {
	cmdflags := flagStruct{
		ConfigPrint: flag.Bool("configprint", false, "a bool"),
	}
	flag.Parse()
	return cmdflags
}

func main() {
	cmdflags := flags()
	logger, err := logger.Logger()
	if err != nil {
		panic(err)
	}
	api_server, err := api.NewServer(logger, http.NewServeMux())
	// handle -configprint=true
	if *cmdflags.ConfigPrint {
		fmt.Println(api_server.Config.GetConfigAsJson())
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}

	web_server, err := web.NewServer(log.Default(), http.NewServeMux())
	if err != nil {
		log.Fatal(err)
	}

	// run api server
	go http.ListenAndServe(":"+strconv.Itoa(api_port), api_server)

	// run web server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(web_port), web_server))
}
