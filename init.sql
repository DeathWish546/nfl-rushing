DROP DATABASE IF EXISTS nfl;

CREATE DATABASE nfl;

USE nfl;

CREATE TABLE IF NOT EXISTS `playerRushingStats` (
    id int(9) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    team varchar(3) NOT NULL,
    position varchar(3) NOT NULL,
    yards int(11) NOT NULL,
    longest int(11) NOT NULL,
    touchdowns int(5) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
