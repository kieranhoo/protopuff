package service

import (
	"context"
	"protopuff/internal/config"
	"protopuff/internal/gen/v1/greeter"
	"protopuff/internal/module/tasks"
	"protopuff/pkg/x/worker"
	"time"
)

func NewGreeter() greeter.GreeterServer {
	return &Greeter{}
}

type Greeter struct {
	greeter.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(ctx context.Context, msg *greeter.MessageRequest) (*greeter.MessageReply, error) {
	return &greeter.MessageReply{
		Message:   "Hello " + msg.GetMessage(),
		Timestamp: time.Now().UTC().Unix(),
	}, nil
}

func (g *Greeter) HeathCheck(ctx context.Context, msg *greeter.StringMessage) (*greeter.StringMessage, error) {
	if err := worker.Exec(config.CriticalQueue, worker.NewTask(
		tasks.WorkerHealthCheck,
		1,
	)); err != nil {
		return nil, err
	}
	return &greeter.StringMessage{
		Value: msg.Value,
	}, nil
}
