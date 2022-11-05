package espn

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func GetGameLines() (allLines []GameLine, err error) {
	sc := newScoreboardClient(&http.Client{Timeout: 10 * time.Second}, "http://site.api.espn.com")
	gids, err := sc.gameIDs()
	if err != nil {
		return nil, fmt.Errorf("getting game IDs: %w", err)
	}

	gc := newGamecastClient(&http.Client{Timeout: 10 * time.Second}, "https://www.espn.com")
	for _, gid := range gids {
		lines, err := gc.gameLines(gid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: getting game lines: %d: %v\n", gid, err)
			continue
		}

		for _, line := range lines {
			allLines = append(allLines, line)
		}
	}

	return allLines, nil
}
