package app

import (
	"log"

	"github.com/Dmytro-yakymuk/task_nix/internal/config"
	"github.com/Dmytro-yakymuk/task_nix/internal/handler"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
	"github.com/Dmytro-yakymuk/task_nix/internal/service"
	"github.com/Dmytro-yakymuk/task_nix/pkg/database/mysql"
	"github.com/labstack/echo/v4"
)

// @title Task NIX
// @version 2.0
// @description This is REST API with echo framework.

// @host localhost:8080
// @BasePath /api/v1
func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db := mysql.ConnectDB(cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Name)
	e := echo.New()

	// Services, Repos & API Handlers
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)
	handlers.Init(e)

	// HTTP Server
	e.Logger.Fatal(e.Start(":" + cfg.HTTP.Port))

}
