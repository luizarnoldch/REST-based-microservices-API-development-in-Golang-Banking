package main

import (
	//"log"

	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking-lib/logger"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking/app"
)

func main() {
	//log.Println("Starting our application")
	logger.Info("Starting our application...")
	app.Start()
}
