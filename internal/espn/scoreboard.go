package espn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type scoreboardClient struct {
	httpClient *http.Client
	baseURL    string
}

func newScoreboardClient(httpClient *http.Client, baseURL string) scoreboardClient {
	return scoreboardClient{httpClient: httpClient, baseURL: baseURL}
}

func (c scoreboardClient) gameIDs() ([]int, error) {
	url, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %s: %w", c.baseURL, err)
	}

	url.Path = "apis/site/v2/sports/basketball/nba/scoreboard"
	url.RawQuery = "lang=en&region=us&calendartype=blacklist&limit=100"
	// url.RawQuery = "lang=en&region=us&calendartype=blacklist&limit=100&dates=20211117"

	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("http request: %s: %w", url.String(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status not successful: %d", resp.StatusCode)
	}

	var sb struct {
		Events []struct {
			Id     string
			Status struct {
				Type struct {
					State string // values: "pre", "in", "post"
				}
			}
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&sb)
	if err != nil {
		return nil, fmt.Errorf("decoding response body: %w", err)
	}

	var ids []int
	for _, e := range sb.Events {
		if e.Status.Type.State != "pre" {
			id, err := strconv.Atoi(e.Id)
			if err != nil {
				// TODO: add a logger
				fmt.Fprintf(os.Stderr, "Warning: converting game ID to integer: %s: %v\n", e.Id, err)
				continue
			}
			ids = append(ids, id)
		}
	}

	return ids, nil
}
