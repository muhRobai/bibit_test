package service

import (
	"context"
	"log"

	"github.com/bibit/microservice/domain"
	"github.com/bibit/microservice/model"
)

type serviceAPI struct {
	movieAPI domain.MovieAPI
}

func NewServiceAPI(movieAPI domain.MovieAPI) domain.ServiceAPI {
	return &serviceAPI{
		movieAPI: movieAPI,
	}
}

func (c serviceAPI) GetListMovie(ctx context.Context, req *model.GetMovieRequest) (*model.GetMovieResponse, error) {
	resp, err := c.movieAPI.GetListMovie(ctx, req)
	if err != nil {
		log.Println("[failed to get movie]", err)
		return nil, err
	}

	resp.Page = req.Page
	return resp, nil
}
