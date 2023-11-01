package main

import (
	"log"
	"os"
	"protopuff/cmd/server"
	"protopuff/internal/config"
	"protopuff/pkg/lib/worker"
	"sort"

	"github.com/urfave/cli/v2"
)

func init() {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

func NewClient() *cli.App {
	_app := &cli.App{
		Name:        "protopuff",
		Usage:       "grpc & http client",
		Version:     "0.0.1",
		Description: "gRPC & API server",
		Commands:    server.Command,
		// Flags:       module.Flag,
	}

	sort.Sort(cli.FlagsByName(_app.Flags))
	sort.Sort(cli.CommandsByName(_app.Commands))

	return _app
}

func main() {
	client := NewClient()

	if err := client.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
