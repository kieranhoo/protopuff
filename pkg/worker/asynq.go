package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/hibiken/asynq"
)

type Queue map[string]int
type Path map[string]func(context.Context, *asynq.Task) error
type Engine struct {
	server *asynq.Server
	mux    *asynq.ServeMux
	path   Path
	queue  Queue
}

var broker asynq.RedisClientOpt

func SetBroker(host, port, password string) {
	broker = asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", host, port), // Redis server address
		Password: password,                         // Redis password
	}
}

func NewServer(concurrency int, queue Queue) *Engine {
	return &Engine{
		server: asynq.NewServer(broker, asynq.Config{
			// Specify how many concurrent workers to use.
			Concurrency: concurrency,
			// Specify multiple queues with different priority.
			Queues: queue,
		}),
		mux:   asynq.NewServeMux(),
		queue: queue,
	}
}

func (w *Engine) HandleFunctions(path Path) {
	w.path = path
	for k, v := range path {
		w.mux.HandleFunc(k, v)
	}
}

func getName(input string) string {
	paths := strings.Split(input, "/")
	return paths[len(paths)-1]
}

func (w *Engine) Run() error {
	log.Info("Launching a asynchronous worker with the following settings:")
	log.Info("Broker:", "redis", broker.Addr)
	for q, p := range w.queue {
		log.Info("-", "queue", q, "priority", p)
	}
	log.Info("Handle Function: ")
	for types, handler := range w.path {
		log.Info("-", "typename", types, "handler", getName(runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()))
	}
	fmt.Println()
	if err := w.server.Run(w.mux); err != nil {
		return err
	}
	return nil
}

func Exec(queue string, task *asynq.Task) error {
	// Create a new Asynq client.
	client := asynq.NewClient(broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)
	// Process the task immediately in critical queue.
	_, err := client.Enqueue(
		task,               // task payload
		asynq.Queue(queue), // set queue for task
	)
	if err != nil {
		return err
	}
	return nil

}

func NewTask(typename string, data interface{}) *asynq.Task {
	payload, _ := json.Marshal(data)
	return asynq.NewTask(typename, payload)
}

func Delay(delay *time.Duration, queue string, task *asynq.Task) error {
	// Create a new Asynq client.
	client := asynq.NewClient(broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)
	if _, err := client.Enqueue(
		task,                    // task payload
		asynq.Queue(queue),      // set queue for task
		asynq.ProcessIn(*delay), // set time to process task
	); err != nil {
		return err
	}
	return nil
}
