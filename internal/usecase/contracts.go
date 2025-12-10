package usecase

import "context"

type (
	Fibonacci interface {
		CalculateFibonacci(context.Context, int64) (int64, error)
	}
	RandomList interface {
		GetRandomList(context.Context, int64) ([]int64, error)
	}
)
