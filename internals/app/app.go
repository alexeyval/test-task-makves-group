package app

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/alexeyval/test-task-makves-group/api"
	"github.com/alexeyval/test-task-makves-group/api/middleware"
	"github.com/alexeyval/test-task-makves-group/internals/app/handlers"
	"github.com/alexeyval/test-task-makves-group/internals/app/processors"
	"github.com/alexeyval/test-task-makves-group/internals/app/storage"
	"github.com/alexeyval/test-task-makves-group/internals/cfg"
)

type Server struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
}

func NewServer(ctx context.Context, config cfg.Cfg) *Server {
	return &Server{
		ctx:    ctx,
		config: config,
	}
}

func (server *Server) Serve() {
	log.Printf("Starting server on port %s", server.config.Port)

	usersStorage := storage.NewStorage(server.config.BatchSize)
	usersStorage.Upload(server.config.FileName)

	usersProcessor := processors.NewUsersProcessor(usersStorage)

	userHandler := handlers.NewUsersHandler(usersProcessor)

	routes := api.CreateRoutes(userHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:              ":" + server.config.Port,
		Handler:           routes,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Println("Server started")

	err := server.srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func (server *Server) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(server.ctx, 5*time.Second)

	if err := server.srv.Shutdown(ctxShutDown); err != nil {
		cancel()
		log.Fatalf("server Shutdown Failed: %s", err)
	}

	cancel()
	log.Println("server exited properly")
}
