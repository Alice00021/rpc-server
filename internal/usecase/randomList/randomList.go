package randomList

import (
	"context"
	"github.com/Alice00021/rpc-server/internal/entity"
	"github.com/Alice00021/test_common/pkg/logger"
	"math/rand"
)

type useCase struct {
	l logger.Interface
}

func New(l logger.Interface) *useCase {
	return &useCase{
		l: l,
	}
}

func (f *useCase) GetRandomList(ctx context.Context, n int64) ([]int64, error) {
	if n < 0 {
		return nil, entity.ErrInvalidNumber
	}

	result := make([]int64, n)
	for i := range result {
		result[i] = rand.Int63n(n)
	}

	return result, nil
}
