package tasks

import "protopuff/pkg/x/worker"

const (
	WorkerHealthCheck string = "Worker.HealthCheck"
)

func Path() worker.Path {
	return worker.Path{
		WorkerHealthCheck: HandleHealthCheck,
	}
}
