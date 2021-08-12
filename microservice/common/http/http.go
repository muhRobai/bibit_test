package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	resty "github.com/go-resty/resty/v2"
)

type RestyRequest struct {
	Method         string
	Headers        map[string]string
	Params         map[string]string
	FormData       map[string]string
	Client         *resty.Client
	Ctx            context.Context
	URL            string
	Body           interface{}
	FilterResponse *ResponseFilter
	Files          []*resty.File
}

type ResponseFilter struct {
	Filter func(req *resty.Response) (*resty.Response, error)
}

var (
	//client3second is httpclient with 3 second timeout set.
	client3second *http.Client
)

const (
	defaultTimeout = 3 * time.Second
)

// Client http implementation wrapper
type Client struct {
	restyClient *resty.Client
}

func NewClient() *Client {
	return &Client{
		restyClient: resty.New(),
	}
}

func (c *Client) Get() *resty.Client {
	return c.restyClient
}

// RestyRequest http request wrapper via resty client
func HTTPRestyRequest(req *RestyRequest) (*resty.Response, error) {
	client := req.Client
	if req.Client == nil {
		client = resty.New().SetTimeout(defaultTimeout)
	}

	log.Println("[body request]", fmt.Sprintf("%v", req.Body))
	restyReq := client.R().
		SetHeaders(req.Headers).
		SetQueryParams(req.Params).
		SetFormData(req.FormData).
		SetContext(req.Ctx).
		SetBody(req.Body)

	for _, v := range req.Files {
		restyReq.SetFileReader(v.ParamName, v.Name, v.Reader)
	}

	resp, err := restyReq.
		Execute(req.Method, req.URL)

	if err != nil {
		log.Println("[failed to geting data]", err)
		return nil, err
	}

	if req.FilterResponse != nil {
		resp, err = req.FilterResponse.Filter(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
