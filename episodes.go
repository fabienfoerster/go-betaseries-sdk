package betaseries

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Episode is an exported type TODO
type Episode struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	Show    struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"show"`
	Code string `json:"code"`
}

// SearchName is an exported function
func (e Episode) SearchName() string {
	tmp := fmt.Sprintf("%s %s 720p", e.Show.Title, e.Code)
	return tmp
}

type betaseriesResponse struct {
	Shows []struct {
		Unseen []Episode `json:"unseen"`
	} `json:"shows"`
	Errors []interface{} `json:"errors"`
}

func transformResponse(resp betaseriesResponse) []Episode {
	var episodes []Episode
	for _, show := range resp.Shows {
		episodes = append(episodes, show.Unseen...)
	}
	return episodes
}

//Episodes retrieve your unseen episode from betaseries
func (api BetaAPI) Episodes() []Episode {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.betaseries.com/episodes/list", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-BetaSeries-Version", "2.4")
	req.Header.Add("X-BetaSeries-Key", api.Key)
	req.Header.Add("X-BetaSeries-Token", api.Token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var betaResp betaseriesResponse
	err = json.NewDecoder(resp.Body).Decode(&betaResp)
	if err != nil {
		log.Fatal(err)
	}
	return transformResponse(betaResp)
}
