package main

import "protopuff/internal/gateway"

func main() {
	gate := gateway.New("localhost:8000", "localhost:8080")
	gate.Serve()
}
