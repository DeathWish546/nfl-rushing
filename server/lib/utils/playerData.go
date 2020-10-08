package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/DeathWish546/nfl-rushing/lib/models"
)

//Parses request body into player stats list
func ParsePlayerData(postBody []byte) ([]models.PlayerStat, error) {
	playerStats := []models.PlayerStat{}
	var data []byte
	var err error

	if len(postBody) > 0 {
		data = postBody
		//if no data body was found, use rushing.json as input
	} else {
		log.Println("No body found, using rushing.json file")
		data, err = ioutil.ReadFile("rushing.json")

		if err != nil {
			log.Println("ERROR reading file")
			return nil, err
		}
	}

	err = json.Unmarshal(data, &playerStats)
	if err != nil {
		log.Println("ERROR unmarshalling json")
		return nil, err
	}

	//Data can come in as non-expected formats, need to normalize first
	playerStats, err = normalizePlayerData(playerStats)
	if err != nil {
		log.Println("ERROR normalizing data")
		return nil, err
	}

	return playerStats, nil
}

func normalizePlayerData(playerStats []models.PlayerStat) ([]models.PlayerStat, error) {
	var normalizedStats []models.PlayerStat
	for _, player := range playerStats {
		//normalize Yards
		yards, err := normalizeYards(player.YardsRaw)
		if err != nil {
			return nil, err
		}
		player.Yards = yards

		//normalize Longest
		longest, td, err := normalizeLongest(player.LongestRaw)
		if err != nil {
			return nil, err
		}
		player.Longest = longest
		player.LongestTouchdown = td

		normalizedStats = append(normalizedStats, player)
	}
	return normalizedStats, nil
}

//yards can come in as string, need to make sure everything is int
func normalizeYards(yardsRaw interface{}) (int, error) {
	var yards int
	var err error
	switch yardsRaw.(type) {
	case string:
		yardsData := yardsRaw.(string)
		if strings.Contains(yardsData, ",") {
			yardsData = strings.ReplaceAll(yardsData, ",", "")
		}
		yards, err = strconv.Atoi(yardsData)
		if err != nil {
			return 0, err
		}
	default:
		yards = int(yardsRaw.(float64))
	}
	return yards, nil
}

//longest can be string and contain T, need to split it out
func normalizeLongest(longestRaw interface{}) (int, bool, error) {
	var longest int
	var td bool
	var err error
	switch longestRaw.(type) {
	case string:
		longestData := longestRaw.(string)
		if strings.Contains(longestData, "T") {
			longestData = strings.ReplaceAll(longestData, "T", "")
			td = true
		}
		longest, err = strconv.Atoi(longestData)
		if err != nil {
			return 0, td, err
		}
	default:
		longest = int(longestRaw.(float64))
	}
	return longest, td, nil
}
