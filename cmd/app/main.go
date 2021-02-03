package main

import "github.com/Dmytro-yakymuk/task_nix/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
