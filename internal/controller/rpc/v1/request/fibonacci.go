package request

type FibRequest struct {
	N int `json:"n" validate:"required"`
}
