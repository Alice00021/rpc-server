package app

import (
	"github.com/Alice00021/rpc-server/config"
	"github.com/Alice00021/rpc-server/internal/controller/rpc"
	"github.com/Alice00021/rpc-server/internal/di"
	commonServer "github.com/Alice00021/test_common/pkg/json-rpc/server"
	"github.com/Alice00021/test_common/pkg/logger"
	"log"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.NewMultipleWriter(
		logger.Level(cfg.Log.Level),
	)

	uc := di.NewUseCase(l)

	srv := commonServer.NewServer(cfg.HTTP.Port)

	rpc.NewRouter(srv, uc, l)

	log.Printf("%s v%s running on port %s", cfg.App.Name, cfg.App.Version, cfg.HTTP.Port)

	if err := srv.StartTCP(); err != nil {
		log.Fatalf("Failed to start RPC server: %v", err)
	}
}
