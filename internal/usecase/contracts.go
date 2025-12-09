package usecase

import "context"

type (
	Fibonacci interface {
		Calculate(context.Context, int) (int, error)
	}
)
