package main

import (
	//"log"

	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking-Lib/logger"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking/app"
)

func main() {
	//log.Println("Starting our application")
	logger.Info("Starting our application...")
	app.Start()
}
