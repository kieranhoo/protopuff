package server

import (
	"protopuff/internal/config"
	"protopuff/internal/gateway"
	"protopuff/internal/mod/tasks"
	"protopuff/pkg/lib/worker"
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
	gate := gateway.New()
	return gate.ServeGateway()
}
