package service

import "errors"

type FilterSortRequest struct {
	PlayerType string `json:"playerType"`
	SortBy     string `json:"sortBy"`
}

type PlayerType int32

const (
	All PlayerType = iota
	Keeper
	Defender
	Midfielder
	Forward
)

func (pt PlayerType) toString() (string, error) {
	switch pt {
	case Keeper:
		return "keeper", nil
	case Defender:
		return "defender", nil
	case Midfielder:
		return "midfielder", nil
	case Forward:
		return "forward", nil
	case All:
		return "all", nil
	}
	return "", errors.New("Unknown player type")
}

func PlayerTypeFormString(pt string) (PlayerType, error) {
	switch pt {
	case "keeper":
		return Keeper, nil
	case "defender":
		return Defender, nil
	case "midfielder":
		return Midfielder, nil
	case "forward":
		return Forward, nil
	case "all":
		return All, nil
	}

	return All, errors.New("Unknown player type")
}

type SortBy int32

const (
	Name SortBy = iota
	Price
	TotalPoints
	TotalPointsPerMillion
	PointsPerGame
	PointsPerGamePerMillion
	AllTimePoints
	AllTimeMinutesPlayed
	AllTimePointsPer90
	AllTimePointsPer90PerMillion
	Default
)

func (sortBy SortBy) toString() (string, error) {
	switch sortBy {
	case Name:
		return "name", nil
	case Price:
		return "price", nil
	case TotalPoints:
		return "totalPoints", nil
	case TotalPointsPerMillion:
		return "totalPointsPerMillion", nil
	case PointsPerGame:
		return "pointsPerGame", nil
	case PointsPerGamePerMillion:
		return "pointsPerGamePerMillion", nil
	case AllTimePoints:
		return "allTimePoints", nil
	case AllTimeMinutesPlayed:
		return "allTimeMinutesPlayed", nil
	case AllTimePointsPer90:
		return "allTimePointsPer90", nil
	case AllTimePointsPer90PerMillion:
		return "allTimePointsPer90PerMillion", nil
	case Default:
		return "defualt", nil
	}
	return "default", errors.New("Unknown player type")
}

func SortByFormString(sortBy string) (SortBy, error) {
	switch sortBy {
	case "name":
		return Name, nil
	case "price":
		return Price, nil
	case "totalPoints":
		return TotalPoints, nil
	case "totalPointsPerMillion":
		return TotalPointsPerMillion, nil
	case "pointsPerGame":
		return PointsPerGame, nil
	case "pointsPerGamePerMillion":
		return PointsPerGamePerMillion, nil
	case "allTimePoints":
		return AllTimePoints, nil
	case "allTimeMinutesPlayed":
		return AllTimeMinutesPlayed, nil
	case "allTimePointsPer90":
		return AllTimePointsPer90, nil
	case "allTimePointsPer90PerMillion":
		return AllTimePointsPer90PerMillion, nil
	case "default":
		return Default, nil
	}
	return Default, errors.New("Unknown player type")
}
