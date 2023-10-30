package config

var (
	WorkerQueue   = "default_queue"
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

func init() {
	if StageStatus != "prod" {
		WorkerQueue = "default_queue_dev"
		CriticalQueue = "critical_dev"
		DefaultQueue = "default_dev"
		LowQueue = "low_dev"
	}
}
