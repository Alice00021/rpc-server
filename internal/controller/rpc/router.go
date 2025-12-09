package rpc

import (
	"github.com/Alice00021/rpc-server/internal/controller/rpc/v1"
	"github.com/Alice00021/rpc-server/internal/di"
	commonServer "github.com/Alice00021/test_common/pkg/json-rpc/server"
	"github.com/Alice00021/test_common/pkg/logger"
)

func NewRouter(srv *commonServer.Server, uc *di.UseCase, l logger.Interface) error {

	fibService := v1.NewFibService(uc.Fibonacci, l)

	return srv.Register(fibService)
}
