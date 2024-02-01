package main

import (
	"lrucache/application/services"
	"lrucache/application/storage"
	"lrucache/infrastructure/database/psql"
	"lrucache/interfaces/handlers"

	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	log.Info("LRU Cache Example")

	var args struct {
		storage.StorageLRUCfg
		handlers.HandlersCfg
		psql.PSQLCfg
	}

	arg.MustParse(&args)

	log.Printf("Start configuration %+v:", args)

	userRepo, err := psql.NewUserRepository(args.PSQLCfg)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorageLRU(args.StorageLRUCfg)

	services := services.New(storage, userRepo)

	handlers := handlers.NewHandlers(services)
	handlers.Run(args.HandlersCfg)

}
