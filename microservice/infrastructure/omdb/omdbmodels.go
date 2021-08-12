package omdb

type MovieItem struct {
	Title  string
	Year   string
	ImbdID string `json:"imdbID"`
	Types  string `json:"type"`
	Poster string
}

type MovieList struct {
	Search      []*MovieItem
	TotalResult string `json:"totalResults"`
	Response    string
}
