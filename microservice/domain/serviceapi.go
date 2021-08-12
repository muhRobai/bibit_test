package domain

import (
	"context"

	"github.com/bibit/microservice/model"
)

// this internal servic domain
type ServiceAPI interface {
	GetListMovie(ctx context.Context, req *model.GetMovieRequest) (*model.GetMovieResponse, error)
}
