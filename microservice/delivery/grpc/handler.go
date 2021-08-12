package grpc

import (
	"context"
	"log"

	"github.com/bibit/microservice/domain"
	"github.com/bibit/microservice/model"
	"github.com/bibit/microservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcHandler struct {
	service domain.ServiceAPI
}

func NewGrpcHandler(svr *grpc.Server, svc domain.ServiceAPI) {
	grpcHandler := &grpcHandler{
		service: svc,
	}

	proto.RegisterMovieHandlerServer(svr, grpcHandler)
	reflection.Register(svr)
}

func (s *grpcHandler) GetListMovie(ctx context.Context, req *proto.ListMovieRequest) (*proto.MovieList, error) {
	resp, err := s.service.GetListMovie(ctx, &model.GetMovieRequest{
		Movie: req.Movie,
		Page:  req.Page,
	})

	if err != nil {
		log.Println("[failed to get movie]", err)
		return nil, err
	}

	var list []*proto.MovieItem
	for _, val := range resp.List {
		list = append(list, &proto.MovieItem{
			Title:    val.Title,
			Year:     val.Year,
			MovieId:  val.MovieID,
			Type:     val.Types,
			MovieUrl: val.MovieURL,
		})
	}

	return &proto.MovieList{
		List: list,
		Page: req.Page,
	}, nil
}
