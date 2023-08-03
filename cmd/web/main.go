package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/alexeyval/test-task-makves-group/internals/app"
	"github.com/alexeyval/test-task-makves-group/internals/cfg"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, ctx)

	go func() {
		oscall := <-c
		log.Printf("system call: %+v\n", oscall)
		server.Shutdown()
		cancel()
	}()

	server.Serve()
}
