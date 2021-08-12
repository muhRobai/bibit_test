package domain

import (
	"context"

	"github.com/bibit/microservice/model"
)

// this third party domian
type MovieAPI interface {
	GetListMovie(ctx context.Context, req *model.GetMovieRequest) (*model.GetMovieResponse, error)
}
