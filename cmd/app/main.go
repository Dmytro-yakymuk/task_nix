package main

import (
	_ "github.com/Dmytro-yakymuk/task_nix/docs"
	"github.com/Dmytro-yakymuk/task_nix/internal/app"
)

const configPath = "configs/main"

// @title Task NIX
// @version 2.0
// @description This is REST API with echo framework.

// @host localhost:8080
// @BasePath /api/v1
func main() {
	app.Run(configPath)
}
