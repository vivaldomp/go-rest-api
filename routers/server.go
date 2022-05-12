package routers

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/vivaldomp/go-rest-api/configuration"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server, error) {
	config := configuration.GetConfig()
	log.Info("Configuring server ...")
	router, err := Init(config)
	cobra.CheckErr(err)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{s}, nil
}

func (srv *Server) Start() {
	log.Info("starting server...")
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			//panic(err)
		}
	}()
	log.Infof("Listening on %s\n", srv.Addr)

	// Listen for the interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify use of shutdown
	stop()
	log.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Info("Server exiting")
}
