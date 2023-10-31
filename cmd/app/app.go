package app

import (
	"protopuff/internal/config"
	"protopuff/internal/gateway"
	"protopuff/internal/module/tasks"
	"protopuff/pkg/x/worker"
)

func AsyncWorker(concurrency int) error {
	w := worker.NewServer(concurrency, worker.Queue{
		config.CriticalQueue: 6, // processed 60% of the time
		config.DefaultQueue:  3, // processed 30% of the time
		config.LowQueue:      1, // processed 10% of the time
	})
	w.HandleFunctions(tasks.Path())
	return w.Run()
}

func APIGateway() error {
	gate := gateway.New("localhost:8000", "localhost:8080")
	return gate.Serve()
}
