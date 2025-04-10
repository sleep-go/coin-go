package base

import (
	"log"
	"net/http"
	"net/url"
)

const (
	TestApi = "https://api.shasta.trongrid.io"
	ProdApi = "https://api.trongrid.io"
)

type Client struct {
	Debug   bool
	baseUrl string
	*http.Client
}

func (c *Client) Get(path string, values url.Values) (*http.Response, error) {
	urls := c.baseUrl + path + "?" + values.Encode()
	req, _ := http.NewRequest("GET", urls, nil)
	log.Println(urls)
	req.Header.Add("accept", "application/json")
	response, err := http.DefaultClient.Do(req)
	return response, err
}

func NewClient(httpClient *http.Client, debug bool) *Client {
	var baseUrl = TestApi
	if !debug {
		baseUrl = ProdApi
	}
	return &Client{
		Debug:   debug,
		baseUrl: baseUrl,
		Client:  httpClient,
	}
}
