package models

type PlayerStat struct {
	Name             string  `json:"Player"` //Player's name
	Team             string  `json:"Team"`   //Player's team abreviation
	Position         string  `json:"Pos"`    //Player's position
	AttPerGameAvg    float64 `json:"Att/G"`  //Rushing Attempts Per Game Average
	Attempts         int     `json:"Att"`    //Rushing Attempts
	Yards            int     `json:"Yds"`    //Total Rushing Yards
	AvgYardsPerAtt   float64 `json:"Avg"`    //Rushing Average Yards Per Attempt
	YardsPerGame     float64 `json:"Yds/G"`  //Rushing Yards Per Game
	Touchdowns       int     `json:"TD"`     //Total Rushing Touchdowns
	Longest          string  `json:"Lng"`    //Longest Rush -- a T represents a touchdown occurred
	FirstDowns       int     `json:"1st"`    //Rushing First Downs
	FirstDownPercent float64 `json:"1st%"`   //Rushing First Down Percentage
	Over20Yards      int     `json:"20+"`    //Rushing 20+ Yards Each
	Over40Yards      int     `json:"40+"`    //Rushing 40+ Yards Each
	Fumbles          int     `json:"FUM"`    //Rushing Fumbles
}
