package espn

import "fmt"

func GetGameLines() ([]int, error) {
	sc := newScoreboardClient()
	gids, err := sc.gameIDs()
	if err != nil {
		return nil, fmt.Errorf("getting game IDs: %w", err)
	}
	return gids, nil
}
