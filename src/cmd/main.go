package main

import (
	"github.com/javadkavossi/Golange_Clean_webApi/src/api"
	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/infra/cache"
	db "github.com/javadkavossi/Golange_Clean_webApi/src/infra/persistence/database"
	"github.com/javadkavossi/Golange_Clean_webApi/src/infra/persistence/migration"

	"github.com/javadkavossi/Golange_Clean_webApi/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
// @description Type "Bearer " followed by your token

func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	
	migration.Up1()

	api.InitServer(cfg)
}
