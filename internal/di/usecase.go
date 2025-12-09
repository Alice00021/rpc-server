package di

import (
	"github.com/Alice00021/rpc-server/internal/usecase"
	"github.com/Alice00021/rpc-server/internal/usecase/fibonacci"
	"github.com/Alice00021/test_common/pkg/logger"
)

type UseCase struct {
	Fibonacci usecase.Fibonacci
}

func NewUseCase(
	l logger.Interface,
) *UseCase {
	FibonacciUc := fibonacci.New(l)

	return &UseCase{
		Fibonacci: FibonacciUc,
	}
}
