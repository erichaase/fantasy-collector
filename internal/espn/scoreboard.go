package espn

import (
	"net/http"
	"net/url"
)

type scoreboardClient struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func newScoreboardClient(httpClient *http.Client, baseURL *url.URL) scoreboardClient {
	return scoreboardClient{httpClient: httpClient, baseURL: baseURL}
}

func (c scoreboardClient) gameIDs() ([]int, error) {
	return []int{5}, nil
}
