package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/peter-mcconnell/automationthingy/api"
	"github.com/peter-mcconnell/automationthingy/logger"
	"github.com/peter-mcconnell/automationthingy/web"
)

type flagStruct struct {
	ConfigPrint *bool
	ConfigPprof *bool
}

func flags() flagStruct {
	cmdflags := flagStruct{
		ConfigPrint: flag.Bool("configprint", false, "print config to stdout and exit"),
		ConfigPprof: flag.Bool("pprof", false, "create pprof file"),
	}
	flag.Parse()
	return cmdflags
}

func main() {
	ctx := context.Background()
	cmdflags := flags()
	logger, err := logger.Logger()
	if err != nil {
		panic(err)
	}
	if *cmdflags.ConfigPprof {
		logger.Info("enabling pprof. writing to automationthingy.pprof")
		f, err := os.Create("automationthingy.pprof")
		if err != nil {
			panic(err)
		}
		// runtime.SetCPUProfileRate(500)
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		defer f.Close()

		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		onKill := func(c chan os.Signal) {
			select {
			case <-c:
				defer f.Close()
				defer pprof.StopCPUProfile()
				defer os.Exit(0)
			}
		}
		go onKill(c)
	}

	logger.Info("instantiating servers")

	api_server, err := api.NewServer(logger, http.NewServeMux())
	// handle -configprint=true
	if *cmdflags.ConfigPrint {
		cfgJ, err := api_server.Config.GetConfigAsJson()
		if err != nil {
			panic(err)
		}
		fmt.Println(cfgJ)
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}

	web_server, err := web.NewServer(ctx, logger, http.NewServeMux())
	if err != nil {
		log.Fatal(err)
	}

	// run api server
	logger.Info("running api server")
	api_server.RunBackground()

	// run web server
	logger.Info("running web server")
	log.Fatal(http.ListenAndServe(":"+web_server.Config.General.Web.Port, web_server.Mux))
}
