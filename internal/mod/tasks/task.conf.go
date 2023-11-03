package tasks

import (
	"protopuff/pkg/lib/worker"
)

const (
	WorkerHealthCheck string = "Worker.HealthCheck"
)

func Path() worker.Path {
	return worker.Path{
		WorkerHealthCheck: HandleHealthCheck,
	}
}
