package domain

import "context"

type Response struct {
	Message string `json:"message"`
}

type Usecase interface {
	Healthy(ctx context.Context) error
}
