package app

import (
	"log"

	"github.com/Dmytro-yakymuk/task_nix/internal/config"
	"github.com/Dmytro-yakymuk/task_nix/internal/handler"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
	"github.com/Dmytro-yakymuk/task_nix/internal/server"
	"github.com/Dmytro-yakymuk/task_nix/internal/service"
	"github.com/Dmytro-yakymuk/task_nix/pkg/database/mysql"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db := mysql.ConnectDB(cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Name)

	// Services, Repos & API Handlers
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)
	handlers.Init()

	// HTTP Server
	srv := server.NewServer(cfg)
	log.Print("Server started")
	if err := srv.Run(); err != nil {
		log.Fatalf("error occurred while running http server: %s\n", err.Error())
	}

}
