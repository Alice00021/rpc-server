package v1

import (
	"context"
	"errors"
	"github.com/Alice00021/rpc-server/internal/controller/rpc/v1/request"
	"github.com/Alice00021/rpc-server/internal/entity"
	"github.com/Alice00021/rpc-server/internal/usecase"
	"github.com/Alice00021/test_common/pkg/logger"
)

type RandomListService struct {
	uc usecase.RandomList
	l  logger.Interface
}

func NewRandomListService(uc usecase.RandomList, l logger.Interface) *RandomListService {
	return &RandomListService{uc: uc, l: l}
}

func (s *RandomListService) GetRandomList(req request.FibRequest, resp *[]int64) error {
	result, err := s.uc.GetRandomList(context.Background(), req.N)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidNumber) {
			return err
		}

		s.l.Error(err, "rpc - v1 - GetRandomList")
		return err
	}

	resp = &result
	return nil
}
