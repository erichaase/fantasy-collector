package espn

import (
	"net/http"
	"net/url"
	"time"
)

type scoreboardClient struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func newScoreboardClient() scoreboardClient {
	return scoreboardClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    &url.URL{Scheme: "http", Host: "site.api.espn.com"},
	}
}

func (c scoreboardClient) gameIDs() ([]int, error) {
	return []int{5}, nil
}
