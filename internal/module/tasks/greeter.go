package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

func HandleHealthCheck(_ context.Context, task *asynq.Task) error {
	var num int64
	if err := json.Unmarshal(task.Payload(), &num); err != nil {
		return err
	}
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
