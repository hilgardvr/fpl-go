package service

import (
	"fmt"
	"hilgardvr/go-fpl/clients"
	"log"
	"math"
	"sort"
	"strconv"
)

type PlayerStatsSeason struct {
	Id                      int
	Name                    string
	PlayerType              PlayerType
	TotalPoints             int
	PointsPerGame           float64
	Price                   float64
	MinutesPlayed           int
	TotalPointsPerMillion   float64
	PointsPerGamePerMillion float64
	PercentageOwned         float64
}

type PlayerStatsAllTime struct {
	SeasonStats                  PlayerStatsSeason
	AllTimePoints                int
	AllTimeMinutesPlayed         int
	AllTimePointsPer90           float64
	AllTimePointsPer90PerMillion float64
}

type allTimePlayerDetail struct {
	allTimePoints        int
	allTimeMinutesPlayed int
}

var allPlayersAllTime []PlayerStatsAllTime

func getPlayerHistory(playerId int) (allTimePlayerDetail, error) {
	details, err := clients.GetPlayerDetail(playerId)
	if err != nil {
		return allTimePlayerDetail{}, err
	}
	var pastHistoryTotal allTimePlayerDetail
	for _, detail := range details.HistoryPast {
		pastHistoryTotal.allTimeMinutesPlayed += detail.Minutes
		pastHistoryTotal.allTimePoints += detail.TotalPoints
	}
	fmt.Println(pastHistoryTotal.allTimePoints)
	return pastHistoryTotal, nil
}

func buildAllPlayerStats(bootstrapResponse clients.BootstrapResponse) error {
	var ctr int
	for _, player := range bootstrapResponse.Elements {
		if ctr >= 5 {
			break
		}
		ctr++
		ppg, err := strconv.ParseFloat(player.PointsPerGame, 64)
		if err != nil {
			log.Println("Could not parse points per game", err)
			return err
		}
		percentageSelected, err := strconv.ParseFloat(player.SelectedByPercent, 64)
		if err != nil {
			log.Println("Could not parse percentage selected", err)
			return err
		}
		playerHistory, err := getPlayerHistory(player.Id)
		if err != nil {
			log.Println("Could not get player history", err)
			return err
		}
		var allTimePPG float64
		allTimePoints := playerHistory.allTimePoints + player.TotalPoints
		allTimeMinutes := playerHistory.allTimeMinutesPlayed + player.Minutes
		if allTimeMinutes > 0 {
			allTimePPG = float64(allTimePoints) / float64(allTimeMinutes) * 90
		}
		price := float64(player.NowCost) / 10
		seasonStats := PlayerStatsSeason{
			Id:                      player.Id,
			Name:                    player.FirstName + " " + player.SecondName,
			PlayerType:              PlayerType(player.ElementType),
			TotalPoints:             player.TotalPoints,
			PointsPerGame:           ppg,
			Price:                   price,
			MinutesPlayed:           player.Minutes,
			TotalPointsPerMillion:   math.Round(float64(player.TotalPoints)/price*1000) / 1000,
			PointsPerGamePerMillion: math.Round(ppg/price*1000) / 1000,
			PercentageOwned:         percentageSelected,
		}
		allTimeStats := PlayerStatsAllTime{
			SeasonStats:                  seasonStats,
			AllTimePoints:                playerHistory.allTimePoints + player.TotalPoints,
			AllTimeMinutesPlayed:         playerHistory.allTimeMinutesPlayed + player.Minutes,
			AllTimePointsPer90:           math.Round(allTimePPG*1000) / 1000,
			AllTimePointsPer90PerMillion: math.Round(allTimePPG/float64(seasonStats.Price)*1000) / 1000,
		}
		allPlayersAllTime = append(allPlayersAllTime, allTimeStats)
	}
	return nil
}

func InitAllPlayerStats() error {
	if len(allPlayersAllTime) > 0 {
		return nil
	}
	bootstrapData, err := clients.BootstrapData()
	if err != nil {
		log.Println("Failed to bootstrap data:", err)
		return err
	}
	err = buildAllPlayerStats(bootstrapData)
	if err != nil {
		log.Println("Failed to build data:", err)
		return err
	}
	return nil
}

func GetAllPlayerStats() []PlayerStatsAllTime {
	return allPlayersAllTime
}

func FilterSortByRequest(req FilterSortRequest) []PlayerStatsAllTime {
	playerType, err := PlayerTypeFormString(req.PlayerType)
	if err != nil {
		log.Println("Could not determine player type:", err)
		return allPlayersAllTime
	}
	sortBy, err := SortByFormString(req.SortBy)
	if err != nil {
		log.Println("Could not determine sort by value:", err)
		return allPlayersAllTime
	}
	filteredByPlayerType := filterByPlayerType(playerType)
	sorted := sortBySelection(sortBy, filteredByPlayerType)
	return sorted
}

func filterByPlayerType(playerType PlayerType) []PlayerStatsAllTime {
	if playerType == PlayerType(All) {
		return allPlayersAllTime
	}
	var filtered []PlayerStatsAllTime
	for _, v := range allPlayersAllTime {
		if playerType == v.SeasonStats.PlayerType {
			fmt.Println(v)
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func sortBySelection(sortBy SortBy, players []PlayerStatsAllTime) []PlayerStatsAllTime {
	switch sortBy {
	case Price:
		sort.Slice(players, func(i, j int) bool {
			return players[i].SeasonStats.Price > players[j].SeasonStats.Price
		})
		return players
	}
	return players
}
