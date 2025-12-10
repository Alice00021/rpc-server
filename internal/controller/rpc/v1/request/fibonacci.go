package request

type FibRequest struct {
	N int64 `json:"n" validate:"required"`
}
