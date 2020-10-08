# theScore "the Rush" Interview Challenge
At theScore, we are always looking for intelligent, resourceful, full-stack developers to join our growing team. To help us evaluate new talent, we have created this take-home interview question. This question should take you no more than a few hours.

**All candidates must complete this before the possibility of an in-person interview. During the in-person interview, your submitted project will be used as the base for further extensions.**

### Why a take-home challenge?
In-person coding interviews can be stressful and can hide some people's full potential. A take-home gives you a chance work in a less stressful environment and showcase your talent.

We want you to be at your best and most comfortable.

### A bit about our tech stack
As outlined in our job description, you will come across technologies which include a server-side web framework (like Elixir/Phoenix, Ruby on Rails or a modern Javascript framework) and a front-end Javascript framework (like ReactJS)

### Challenge Background
We have sets of records representing football players' rushing statistics. All records have the following attributes:
* `Player` (Player's name)
* `Team` (Player's team abbreviation)
* `Pos` (Player's postion)
* `Att/G` (Rushing Attempts Per Game Average)
* `Att` (Rushing Attempts)
* `Yds` (Total Rushing Yards)
* `Avg` (Rushing Average Yards Per Attempt)
* `Yds/G` (Rushing Yards Per Game)
* `TD` (Total Rushing Touchdowns)
* `Lng` (Longest Rush -- a `T` represents a touchdown occurred)
* `1st` (Rushing First Downs)
* `1st%` (Rushing First Down Percentage)
* `20+` (Rushing 20+ Yards Each)
* `40+` (Rushing 40+ Yards Each)
* `FUM` (Rushing Fumbles)

In this repo is a sample data file [`rushing.json`](/rushing.json).

##### Challenge Requirements
1. Create a web app. This must be able to do the following steps
    1. Create a webpage which displays a table with the contents of [`rushing.json`](/rushing.json)
    2. The user should be able to sort the players by _Total Rushing Yards_, _Longest Rush_ and _Total Rushing Touchdowns_
    3. The user should be able to filter by the player's name
    4. The user should be able to download the sorted data as a CSV, as well as a filtered subset

2. The system should be able to potentially support larger sets of data on the order of 10k records.

3. Update the section `Installation and running this solution` in the README file explaining how to run your code

### Submitting a solution
1. Download this repo
2. Complete the problem outlined in the `Requirements` section
3. In your personal public GitHub repo, create a new public repo with this implementation
4. Provide this link to your contact at theScore

We will evaluate you on your ability to solve the problem defined in the requirements section as well as your choice of frameworks, and general coding style.

### Help
If you have any questions regarding requirements, do not hesitate to email your contact at theScore for clarification.

# Installation and running this solution

## Prerequisites:
* [Docker](https://docs.docker.com/get-docker/)
* [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installing and setting up
```bash
$ git clone https://github.com/DeathWish546/nfl-rushing.git
$ cd nfl-rushing
$ docker-compose up
```
This should start up 3 containers, the frontend, backend, and db containers.

When running for the first time and you want to fill the database with the data from `rushing.json`, run the following curl request:

```bash
$ curl -X POST http://localhost:8080/players
```

## Usage
### Client
1. Once the docker containers are running, access the NFL Rushing Data viewer at http://localhost:3000/
2. Sort the table by *Total Yards*, *Touchdowns*, or *Longest Rush*
3. Search for a given player by entering their name, or part of their name in the search bar under *Player*
4. Modify how the pagination is controlled with the buttons below the table.
5. Download the data displayed on the page (i.e. however it's displayed after sorted, paginated, etc.) by clicking the *Download Page Data* button. It will save as **page_data.csv**
6. Download all the available data sorted by any of the above conditions or as is by pressing the *Download All Data* button. It will save as **all_data.csv**
7. Refresh the page to reload any data from the server

### Server
1. Can be accessed at http://localhost:8080
2. Stores all player data in a MySQL database
3. Retrieves all player data upon being requested by the client, or any request to the API as described below.

## API Endpoints
All API Endpoints are accessed via `http://localhost:8080`, and communicates using the Player Object:

#### `Player Object`
```
  {
    "Player": <string>,          //Player's Name
    "Team":   <string>,          //Player's team abbreviation
    "Pos":    <string>,          //Player's postion
    "Att/G":  <int>,             //Rushing Attempts Per Game Average
    "Att":    <int>,             //Rushing Attempts
    "Yds":    <int>,             //Total Rushing Yards
    "Avg":    <int>,             //Rushing Average Yards Per Attempt
    "Yds/G":  <int>,             //Rushing Yards Per Game
    "TD":     <int>,             //Total Rushing Touchdowns
    "Lng":    <string> or <int>, //Longest Rush -- a "T" represents a touchdown occurred
    "1st":    <int>,             //Rushing First Downs
    "1st%":   <int>,             //Rushing First Down Percentage
    "20+":    <int>,             //Rushing 20+ Yards Each
    "40+":    <int>,             //Rushing 40+ Yards Each
    "FUM":    <int>,             //Rushing Fumbles
  }
```
### `GET /players`
Retrieves all the available player data from the server
#### Response
* `200 OK` on a successful retrieval
  * Response Body is JSON: a list of `Player` objects
* `204 No Content` if no data is found
  * No response body
* `500 Internal Server Error` if something went wrong on the server side
  * Response Body is JSON: error message describing what went wrong

Sample Success Response Body:
```
[
    {
        "Player": "Joe Banyard",
        "Team": "JAX",
        "Pos": "RB",
        "Att/G": 2,
        "Att": 2,
        "Yds": 7,
        "Avg": 3.5,
        "Yds/G": 7,
        "TD": 0,
        "Lng": 7,
        "1st": 0,
        "1st%": 0,
        "20+": 0,
        "40+": 0,
        "FUM": 0
    },
    {
        "Player": "Shaun Hill",
        "Team": "MIN",
        "Pos": "QB",
        ...
    },
    ...
]

```

### `POST /players`
Adds player data to the existing data set. This will allow the player data to be duplicated (i.e. a player's data will appear twice if it's inserted twice)
#### Request
You can make this request with or without a body

* If you make it with a body, the body must be a JSON list of `Player` objects with all fields filled out.
* If you make it without a body, the server will insert another set of the default data from `rushing.json` into the db.

#### Response
* `200 OK` on a successful insert
  * Response Body is JSON: a list of all `Player` objects that were inserted into the db
* `400 Bad Request` if the body is malformed or missing data
  * Response Body is JSON: error message describing what went wrong
* `500 Internal Server Error` if something went wrong on the server side
  * Response Body is JSON: error message describing what went wrong

Sample Success Response Body:

* Identical to the request body (or `rushing.json` if no body was set)

### `DELETE /players/delete`
:warning: Deletes **ALL** existing player data from the database. Only use if you're certain you want to wipe the data for whatever reason.
#### Response
* `200 OK` on a successful delete
  * Response Body is JSON: message indicating success
* `500 Internal Server Error` if something went wrong on the server side
  * Response Body is JSON: error message describing what went wrong

Sample Success Response Body:
```json
{"success": "ok"}
```
