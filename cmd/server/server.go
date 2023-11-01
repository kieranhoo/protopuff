package server

import (
	"google.golang.org/grpc"
	"protopuff/internal/config"
	"protopuff/internal/gateway"
	"protopuff/internal/module/service"
	"protopuff/internal/module/tasks"
	"protopuff/internal/proto/gen/v1/auth"
	"protopuff/internal/proto/gen/v1/greeter"
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

func registerGrpcServer(s *grpc.Server) {
	greeter.RegisterGreeterServer(s, service.NewGreeter())
	auth.RegisterAuthServer(s, service.NewAuth())
}

func APIGateway() error {
	gate := gateway.New()
	return gate.ServeGateway(registerGrpcServer)
}
