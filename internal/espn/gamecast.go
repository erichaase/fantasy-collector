package espn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

type gamecastClient struct {
	httpClient *http.Client
	baseURL    string
}

func newGamecastClient(httpClient *http.Client, baseURL string) gamecastClient {
	return gamecastClient{httpClient: httpClient, baseURL: baseURL}
}

type GameLine struct {
	Id             int
	FirstName      string
	LastName       string
	PositionAbbrev string
	Jersey         string
	Active         string
	IsStarter      string
	Fg             string
	Ft             string
	Threept        string
	Rebounds       string
	Assists        string
	Steals         string
	Fouls          string
	Points         string
	Minutes        string
	Blocks         string
	Turnovers      string
	PlusMinus      string
	Dnp            bool
	EnteredGame    bool
}

func (c gamecastClient) gameLines(gid int) (lines []GameLine, err error) {
	res, err := c.makeRequest(gid)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer res.Body.Close()

	lines, err = c.parseResponse(res)
	if err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	return lines, nil
}

func (c gamecastClient) makeRequest(gid int) (*http.Response, error) {
	url, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %s: %w", c.baseURL, err)
	}

	url.Path = "nba/gamecast12/master"
	url.RawQuery = fmt.Sprintf("xhr=1&gameId=%d&lang=en&init=true&setType=true&confId=null", gid)

	res, err := c.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get request: %s: %w", url.String(), err)
	}

	return res, nil
}

func (c gamecastClient) parseResponse(res *http.Response) (lines []GameLine, err error) {
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response status not successful: %d", res.StatusCode)
	}

	// should use json.Decode() but the payload includes control chars which we need to strip
	// TODO: better way to do this?
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	bodyStripped := regexp.MustCompile("[[:^print:]]").ReplaceAllLiteralString(string(bodyRaw), "")
	bodyBytes := []byte(bodyStripped)

	var game struct {
		Gamecast struct {
			Stats struct {
				Player struct {
					Home []GameLine
					Away []GameLine
				}
			}
		}
	}

	err = json.Unmarshal(bodyBytes, &game)
	if err != nil {
		return nil, fmt.Errorf("parsing json: %w", err)
	}

	allLines := append(game.Gamecast.Stats.Player.Home, game.Gamecast.Stats.Player.Away...)
	for _, line := range allLines {
		if line.Id != 0 { // totals row has a value of 0 for the id field
			lines = append(lines, line)
		}
	}

	return lines, nil
}
