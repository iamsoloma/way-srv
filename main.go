package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/TinajXD/way-srv/api"
	"github.com/TinajXD/way-srv/config"
)

func main() {
	fmt.Println("Starting...")

	fmt.Println("Reading a config file...")
	cfg := config.GetConf()

	fmt.Println("Preparing a workspace...")
	if _, err := os.Stat(cfg.StoragePath); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(cfg.StoragePath, 0764)
		if err != nil {
			log.Fatal("Can`t Create Storage Folder!")
		}
	}

	//start the api server
	apiServer := api.ApiServer{
		Addr: cfg.HttpServer.Address,
		StoragePath: cfg.StoragePath,
	}

	fmt.Println("API server listening on ", cfg.HttpServer.Address)
	err := apiServer.Start()
	if err != nil {
		log.Fatal("API Server is stoped!\n" + err.Error())
	}
}