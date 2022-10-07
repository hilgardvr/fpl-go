package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type BootstrapResponse struct {
	TotalPlayers int64    `json:"total_players"`
	Elements     []Player `json:"elements"`
}

type Player struct {
	Code              int    `json:"code"`
	ElementType       int    `json:"element_type"`
	FirstName         string `json:"first_name"`
	Form              string `json:"form"`
	Id                int    `json:"id"`
	NowCost           int    `json:"now_cost"`
	PointsPerGame     string `json:"points_per_game"`
	SecondName        string `json:"second_name"`
	SelectedByPercent string `json:"selected_by_percent"`
	TotalPoints       int    `json:"total_points"`
	Minutes           int    `json:"minutes"`
}

type PlayerDetail struct {
	HistoryPast []HistoryPast `json:"history_past"`
}

type HistoryPast struct {
	SeasonName  string `json:"season_name"`
	TotalPoints int    `json:"total_points"`
	Minutes     int    `json:"minutes"`
}

const fplUrl = "https://fantasy.premierleague.com/api/bootstrap-static/"
const playerDetailUrl = "https://fantasy.premierleague.com/api/element-summary/"
const eventStatus = "https://fantasy.premierleague.com/api/event-status/"
const userAgent = "curl/7.82.0"

func BootstrapData() (BootstrapResponse, error) {
	request, err := http.NewRequest("GET", fplUrl, nil)
	request.Header.Add("user-agent", userAgent)
	if err != nil {
		return BootstrapResponse{}, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return BootstrapResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return BootstrapResponse{}, fmt.Errorf("Invalid fpl response code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BootstrapResponse{}, err
	}
	var response BootstrapResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal json: ", err)
		return BootstrapResponse{}, err
	}
	return response, nil
}

func GetPlayerDetail(playerId int) (PlayerDetail, error) {
	url := playerDetailUrl + strconv.Itoa(playerId) + "/"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request: ", err)
		return PlayerDetail{}, err
	}
	request.Header.Add("user-agent", userAgent)
	if err != nil {
		return PlayerDetail{}, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Error requesting player data: ", err)
		return PlayerDetail{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return PlayerDetail{}, fmt.Errorf("Invalid fpl response code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PlayerDetail{}, err
	}
	var response PlayerDetail
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal json: ", err)
		return PlayerDetail{}, err
	}
	return response, nil
}
