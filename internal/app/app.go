package app

import (
	"github.com/Alice00021/rpc-server/config"
	"github.com/Alice00021/rpc-server/internal/controller/rpc"
	"github.com/Alice00021/rpc-server/internal/di"
	commonServer "github.com/Alice00021/test_common/pkg/json-rpc/server"
	"github.com/Alice00021/test_common/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.NewMultipleWriter(
		logger.Level(cfg.Log.Level),
		logger.FileName(cfg.Log.FileName),
	)

	// Use-Case
	uc := di.NewUseCase(l)

	// RPC server
	srv := commonServer.NewServer(cfg.HTTP.Port)
	rpc.NewRouter(srv, uc, l)
	log.Printf("%s v%s running on port %s", cfg.App.Name, cfg.App.Version, cfg.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func() {
		if err := srv.StartTCP(); err != nil {
			errChan <- err
		}
	}()

	select {
	case s := <-interrupt:
		l.Info("signal received: %s", s.String())

	case err := <-errChan:
		l.Error("RPC server error: %v", err)
	}

	if err := srv.Shutdown(); err != nil {
		l.Error("RPC shutdown error: %v", err)
	}
}
