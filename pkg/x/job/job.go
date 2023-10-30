package job

import (
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type Job struct {
	function map[string]string
}

func New() *Job {
	return &Job{
		function: make(map[string]string),
	}
}

func getName(input string) string {
	paths := strings.Split(input, "/")
	return paths[len(paths)-1]
}

func (job *Job) Scheduler(fn func(), _time time.Duration) error {
	job.function[getName(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())] = _time.String()
	ticker := time.NewTicker(_time)
	go func() {
		for range ticker.C {
			fn()
		}
	}()
	return nil
}

func (job *Job) Launch() error {
	log.Info("Launching a job scheduler with the following settings:")
	for k, v := range job.function {
		log.Info("-", "function", k, "schedule", v)
	}
	return nil
}
