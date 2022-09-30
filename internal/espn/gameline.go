package espn

import (
	"fmt"
	"net/http"
	"time"
)

func GetGameLines() ([]int, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	baseURL := "http://site.api.espn.com"
	sc := newScoreboardClient(httpClient, baseURL)

	gids, err := sc.gameIDs()
	if err != nil {
		return nil, fmt.Errorf("getting game IDs: %w", err)
	}

	return gids, nil
}
