package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
)

type PlayerStat struct {
	Name             string      `json:"Player"` //Player's name
	Team             string      `json:"Team"`   //Player's team abreviation
	Position         string      `json:"Pos"`    //Player's position
	AttPerGameAvg    float64     `json:"Att/G"`  //Rushing Attempts Per Game Average
	Attempts         int         `json:"Att"`    //Rushing Attempts
	YardsRaw         interface{} `json:"Yds"`    //Some values are not ints
	Yards            int         `json:"-"`      //Total Rushing Yards
	AvgYardsPerAtt   float64     `json:"Avg"`    //Rushing Average Yards Per Attempt
	YardsPerGame     float64     `json:"Yds/G"`  //Rushing Yards Per Game
	Touchdowns       int         `json:"TD"`     //Total Rushing Touchdowns
	LongestRaw       interface{} `json:"Lng"`    //Some values aren't strings
	Longest          int         `json:"-"`      //Longest Rush -- a T represents a touchdown occurred
	LongestTouchdown bool        `json:"-"`      //Whether or not the longest rush has a touchdown
	FirstDowns       int         `json:"1st"`    //Rushing First Downs
	FirstDownPercent float64     `json:"1st%"`   //Rushing First Down Percentage
	Over20Yards      int         `json:"20+"`    //Rushing 20+ Yards Each
	Over40Yards      int         `json:"40+"`    //Rushing 40+ Yards Each
	Fumbles          int         `json:"FUM"`    //Rushing Fumbles
}

func InsertPlayersIntoDB(db *sql.DB, allPlayerData []PlayerStat) error {
	queryStr := `INSERT INTO playerRushingStats(
        name,
        team,
        position,
        attPerGameAvg,
        attempts,
        yards,
        yardPerAttAvg,
        yardPerGame,
        touchdowns,
        longest,
        longestTouchdown,
        firstDowns,
        firstDownPercent,
        over20Yards,
        over40Yards,
        fumbles
    ) VALUES `
	vals := []interface{}{}
	//               1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16
	const rowSQL = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var inserts []string
	for _, player := range allPlayerData {
		inserts = append(inserts, rowSQL)
		vals = append(vals,
			player.Name,             //1
			player.Team,             //2
			player.Position,         //3
			player.AttPerGameAvg,    //4
			player.Attempts,         //5
			player.Yards,            //6
			player.AvgYardsPerAtt,   //7
			player.YardsPerGame,     //8
			player.Touchdowns,       //9
			player.Longest,          //10
			player.LongestTouchdown, //11
			player.FirstDowns,       //12
			player.FirstDownPercent, //13
			player.Over20Yards,      //14
			player.Over40Yards,      //15
			player.Fumbles,          //16
		)
	}

	queryStr = queryStr + strings.Join(inserts, ",")

	res, err := db.Query(queryStr, vals...)
	defer res.Close()
	if err != nil {
		log.Println("Could not insert into DB")
		return err
	}

	return nil
}

func GetAllPlayers(db *sql.DB) ([]PlayerStat, error) {
	var allPlayerStats []PlayerStat
	queryStr := `SELECT
        name,
        team,
        position,
        attPerGameAvg,
        attempts,
        yards,
        yardPerAttAvg,
        yardPerGame,
        touchdowns,
        longest,
        longestTouchdown,
        firstDowns,
        firstDownPercent,
        over20Yards,
        over40Yards,
        fumbles
    FROM playerRushingStats;`

	rows, err := db.Query(queryStr)
	if err != nil {
		log.Println("Could not read from DB")
		return nil, err
	}

	for rows.Next() {
		var player PlayerStat
		err = rows.Scan(
			&player.Name,             //1
			&player.Team,             //2
			&player.Position,         //3
			&player.AttPerGameAvg,    //4
			&player.Attempts,         //5
			&player.Yards,            //6
			&player.AvgYardsPerAtt,   //7
			&player.YardsPerGame,     //8
			&player.Touchdowns,       //9
			&player.Longest,          //10
			&player.LongestTouchdown, //11
			&player.FirstDowns,       //12
			&player.FirstDownPercent, //13
			&player.Over20Yards,      //14
			&player.Over40Yards,      //15
			&player.Fumbles,          //16
		)

		player.YardsRaw = player.Yards

		if player.LongestTouchdown {
			player.LongestRaw = strconv.Itoa(player.Longest) + "T"
		} else {
			player.LongestRaw = player.Longest
		}

		allPlayerStats = append(allPlayerStats, player)
		if err != nil {
			log.Println("Could not properly scan row")
			return nil, err
		}
	}

	return allPlayerStats, nil
}
