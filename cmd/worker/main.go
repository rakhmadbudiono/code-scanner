package main

import (
	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/controller"
)

type Worker struct {
	Controller controller.IController
	Config     *configs.Config
}

func (w *Worker) Start() {
	w.Controller.RunScanner()
}

func main() {
	cfg := configs.New()
	ctrl := controller.NewController(
		cfg,
		controller.WithDatabase(),
		controller.WithSubscriber(),
	)

	worker := &Worker{
		Config:     cfg,
		Controller: ctrl,
	}
	worker.Start()
}
