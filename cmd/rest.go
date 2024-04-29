package cmd

import (
	"log"

	"github.com/imanudd/clean-arch-pattern/config"
	rest "github.com/imanudd/clean-arch-pattern/internal/delivery/http"
	"github.com/imanudd/clean-arch-pattern/internal/repository"
	"github.com/imanudd/clean-arch-pattern/internal/usecase"
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

		app := rest.NewRest(cfg)

		//init repo
		userRepo := repository.NewUserRepository(pgDB)
		bookRepo := repository.NewBookRepository(pgDB)
		authorRepo := repository.NewAuthorRepository(pgDB)

		//init usecase
		authUseCase := usecase.NewAuthUseCase(cfg, userRepo)
		bookUseCase := usecase.NewBookUseCase(cfg, bookRepo, authorRepo)
		authorUseCase := usecase.NewAuthorUseCase(cfg, authorRepo, bookRepo)

		route := &rest.Route{
			Config:        cfg,
			App:           app,
			AuthUseCase:   authUseCase,
			BookUseCase:   bookUseCase,
			AuthorUseCase: authorUseCase,
			UserRepo:      userRepo,
		}

		route.RegisterRoutes()

		if err := rest.Serve(app, cfg); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}

	},
}
