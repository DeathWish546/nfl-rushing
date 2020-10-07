package models

import ()

type PlayerStat struct {
	Name             string      `json:"Player"` //Player's name
	Team             string      `json:"Team"`   //Player's team abreviation
	Position         string      `json:"Pos"`    //Player's position
	AttPerGameAvg    float64     `json:"Att/G"`  //Rushing Attempts Per Game Average
	Attempts         int         `json:"Att"`    //Rushing Attempts
	YardsRaw         interface{} `json:"Yds"`    //Some values are not ints
	Yards            int         `json:"RY"`     //Total Rushing Yards
	AvgYardsPerAtt   float64     `json:"Avg"`    //Rushing Average Yards Per Attempt
	YardsPerGame     float64     `json:"Yds/G"`  //Rushing Yards Per Game
	Touchdowns       int         `json:"TD"`     //Total Rushing Touchdowns
	LongestRaw       interface{} `json:"Lng"`    //Some values aren't strings
	Longest          int         `json:"RL"`     //Longest Rush -- a T represents a touchdown occurred
	LongestTouchdown bool        `json:"-"`    //Whether or not the longest rush has a touchdown
	FirstDowns       int         `json:"1st"`    //Rushing First Downs
	FirstDownPercent float64     `json:"1st%"`   //Rushing First Down Percentage
	Over20Yards      int         `json:"20+"`    //Rushing 20+ Yards Each
	Over40Yards      int         `json:"40+"`    //Rushing 40+ Yards Each
	Fumbles          int         `json:"FUM"`    //Rushing Fumbles
}
