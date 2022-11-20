package main

import (
	"github.com/rakhmadbudiono/code-scanner/cmd/server/rest"
	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/controller"
)

func main() {
	cfg := configs.New()
	ctrl := controller.NewController(
		cfg,
		controller.WithDatabase(),
	)

	server := rest.NewServer(cfg, ctrl)
	server.Start()
}
