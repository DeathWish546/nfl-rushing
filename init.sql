-- DROP DATABASE IF EXISTS nfl;
-- CREATE DATABASE nfl;

USE nfl;

CREATE TABLE IF NOT EXISTS `playerRushingStats` (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    team varchar(3) NOT NULL,
    position varchar(3) NOT NULL,
    attPerGameAvg float NOT NULL,
    attempts int NOT NULL,
    yards int NOT NULL,
    yardPerAttAvg float NOT NULL,
    yardPerGame float NOT NULL,
    touchdowns int NOT NULL,
    longest int NOT NULL,
    longestTouchdown boolean NOT NULL,
    firstDowns int NOT NULL,
    firstDownPercent float NOT NULL,
    over20Yards int NOT NULL,
    over40Yards int NOT NULL,
    fumbles int NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

