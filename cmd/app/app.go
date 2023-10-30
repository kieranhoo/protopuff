package app

import "protopuff/internal/gateway"

func AsyncWorker(concurrency int) error {
	return nil
}

func APIGateway() error {
	gate := gateway.New("localhost:8000", "localhost:8080")
	return gate.Serve()
}
