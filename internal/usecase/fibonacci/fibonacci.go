package fibonacci

import (
	"context"
	"github.com/Alice00021/rpc-server/internal/entity"
	"github.com/Alice00021/test_common/pkg/logger"
)

type useCase struct {
	l logger.Interface
}

func New(l logger.Interface) *useCase {
	return &useCase{
		l: l,
	}
}

func (f *useCase) Calculate(ctx context.Context, n int) (int, error) {
	if n < 0 {
		return 0, entity.ErrInvalidNumber
	}

	return fib(n), nil
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
