package main

import (
	"log"

	"github.com/javadkavossi/Golange_Clean_webApi/src/api"
	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/data/cache"
	"github.com/javadkavossi/Golange_Clean_webApi/src/data/db"
)

func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}



	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}
	
	api.InitServer(cfg)
}
