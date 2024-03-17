package main

import (
	"fmt"
	"log"

	"github.com/TinajXD/way-srv/api"
	"github.com/TinajXD/way-srv/config"
)

func main() {
	fmt.Println("Starting...")

	cfg := config.GetConf()

	//start the api server
	apiServer := api.ApiServer{
		Addr: cfg.Address,
	}


	err := apiServer.Start()
	if err != nil {
		log.Fatal("API Server is stoped!\n" + err.Error())
	}
}