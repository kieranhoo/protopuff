package server

import (
	"protopuff/internal/config"
	"protopuff/internal/gateway"
	"protopuff/internal/module/service"
	"protopuff/internal/module/tasks"
	"protopuff/internal/proto/gen/v1/greeter"
	"protopuff/pkg/worker"

	"google.golang.org/grpc"
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

func registerGrpcServer(s *grpc.Server) {
	greeter.RegisterGreeterServer(s, service.NewGreeter())
}

func APIGateway() error {
	gate := gateway.New(config.HttpUri, config.RpcUri)
	return gate.ServeGateway(registerGrpcServer)
}
