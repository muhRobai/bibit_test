package model

type GetMovieRequest struct {
	Movie string `json:"movie"`
	Page  string `json:"page"`
}

type MovieItem struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	MovieID  string `json:"movie_id"`
	Types    string `json:"type"`
	MovieURL string `json:"movie_url"`
}

type GetMovieResponse struct {
	List []*MovieItem `json:"list"`
	Page string       `json:"page"`
}
