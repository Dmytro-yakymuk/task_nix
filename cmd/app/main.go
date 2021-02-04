package main

import (
	_ "github.com/Dmytro-yakymuk/task_nix/docs"
	"github.com/Dmytro-yakymuk/task_nix/internal/app"
)

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
