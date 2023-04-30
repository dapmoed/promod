package main

import (
	"promod/internal/usecase"
	"promod/pkg/app"
	"promod/pkg/pm_servers"
)

func main() {
	application := app.New()
	application.AddServer(
		pm_servers.NewGinServer(
			usecase.NewPing(),
		),
	)
	application.Run()
}

func new()
