package omdb

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	commonHttp "github.com/bibit/microservice/common/http"
	"github.com/bibit/microservice/domain"
	"github.com/bibit/microservice/model"
)

type omdbService struct {
	BaseURL string
	Key     string
}

func NewOmdbService(
	baseURL, key string,
) domain.MovieAPI {
	return &omdbService{
		BaseURL: baseURL,
		Key:     key,
	}
}

var (
	errMissingMovieName = errors.New("missing movie name")
)

func (c *omdbService) GetListMovie(ctx context.Context, req *model.GetMovieRequest) (*model.GetMovieResponse, error) {
	page := "0"
	if item := parseStringToInt(req.Page); item != 0 {
		page = req.Page
	}

	if req.Movie == "" {
		return nil, errMissingMovieName
	}

	var queryParams = map[string]string{
		"s":    req.Movie,
		"page": page,
	}

	resp, err := commonHttp.HTTPRestyRequest(
		&commonHttp.RestyRequest{
			Params: queryParams,
			URL:    parseURL(c.BaseURL, c.Key),
			Ctx:    ctx,
		},
	)

	if err != nil {
		log.Println("[failed to get list movie]", err)
		return nil, err
	}

	var p MovieList
	err = json.Unmarshal(resp.Body(), &p)
	if err != nil {
		log.Println("[failed to unmarshal]", err)
		return nil, err
	}

	response := parseListMovieToModel(p)
	return &response, nil
}

func parseURL(baseurl, key string) string {
	return baseurl + "/?apikey=" + key
}

func parseStringToInt(number string) int {
	numb, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}

	return numb
}

func parseListMovieToModel(list MovieList) model.GetMovieResponse {
	var resp []*model.MovieItem
	for _, val := range list.Search {
		resp = append(resp, &model.MovieItem{
			Title:    val.Title,
			Year:     val.Year,
			MovieID:  val.ImbdID,
			Types:    val.Types,
			MovieURL: val.Poster,
		})
	}

	return model.GetMovieResponse{
		List: resp,
	}
}
