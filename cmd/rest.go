package cmd

import (
	"log"

	"github.com/imanudd/inventorySvc-clean-architecture/config"
	rest "github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/usecase"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "rest",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		pgDB := InitPostgreSQL(cfg)

		if cfg.LogMode {
			pgDB = pgDB.Debug()
		}

		// client := InitElastic(cfg)

		//init elasticsearch
		// es := elasticsearch.New(client)

		app := rest.NewRest(cfg)
		repo := repository.NewRepository(pgDB)
		useCase := usecase.NewUsecase(cfg, repo)

		route := &rest.Route{
			Config:     cfg,
			App:        app,
			UseCase:    useCase,
			Repository: repo,
		}

		route.RegisterRoutes()

		if err := rest.Serve(app, cfg); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}

	},
}
