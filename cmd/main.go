// Package main Searcher.
//
//		Schemes: http
//		Host: localhost:8083
//	 	BasePath: /
//		Version: 0.0.1
//
//		Consumes:
//		- application/json
//		Produces:
//		- application/json
//
// swagger:meta
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vitokz/ton-nft-searcher/internal/searcher"
)

func main() {
	service, err := searcher.New()
	if err != nil {
		log.Fatal(err)
	}

	service.Start()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit

	service.Stop()
}
