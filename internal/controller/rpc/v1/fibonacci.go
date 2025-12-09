package v1

import (
	"context"
	"errors"
	"github.com/Alice00021/rpc-server/internal/controller/rpc/v1/request"
	"github.com/Alice00021/rpc-server/internal/entity"
	"github.com/Alice00021/rpc-server/internal/usecase"
	"github.com/Alice00021/test_common/pkg/logger"
)

type FibService struct {
	uc usecase.Fibonacci
	l  logger.Interface
}

func NewFibService(uc usecase.Fibonacci, l logger.Interface) *FibService {
	return &FibService{uc: uc, l: l}
}

func (s *FibService) Calculate(req request.FibRequest, resp *entity.Fib) error {
	result, err := s.uc.Calculate(context.Background(), req.N)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidNumber) {
			return err
		}

		s.l.Error(err, "rpc - v1 - fib")
		return err
	}

	resp.Result = result
	return nil
}
