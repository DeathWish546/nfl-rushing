package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"github.com/DeathWish546/nfl-rushing/lib/models"
	nflUtils "github.com/DeathWish546/nfl-rushing/lib/utils"
)

//Creates an APP that has both the router and DB attached for simplicity
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Opens connection to DB and gets the routes set up
func (a *App) initialize() {
	var err error
	a.DB, err = sql.Open("mysql", "root:password@tcp(db)/nfl")
	if err != nil {
		log.Fatal("Could not initate db: ", err.Error())
		return
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

//Initializes the server
func (a *App) run(addr string) {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

//Assigning all routes
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", sayHello)
	a.Router.HandleFunc("/players", a.uploadData).Methods("POST")
	a.Router.HandleFunc("/players", a.getAllPlayerData).Methods("GET")
	a.Router.HandleFunc("/players/delete", a.deleteAllPlayerData).Methods("DELETE")
}

// POST /players, insert data into DB
func (a *App) uploadData(w http.ResponseWriter, r *http.Request) {
	log.Println("Uploading player data")
	var allPlayerStats []models.PlayerStat

	//Parse the request body into a string
	postBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Could not read request body", err.Error())
		respondWithError(w, http.StatusBadRequest, "Could not read request body: "+err.Error())
		return
	}

	//Marshal the data into a list of player stat objects, or use rushing.json if no body was found
	allPlayerStats, err = nflUtils.ParsePlayerData(postBody)
	if err != nil {
		log.Println("Could not parse player data: ", err.Error())
		respondWithError(w, http.StatusBadRequest, "Could not parse player data: "+err.Error())
		return
	}

	//Insert all player stats into db, or return error if an empty list was given
	if len(allPlayerStats) > 0 {
		err = models.InsertPlayersIntoDB(a.DB, allPlayerStats)
		if err != nil {
			log.Println("Could not insert into db: ", err.Error())
			respondWithError(w, http.StatusInternalServerError, "Could not insert into db: "+err.Error())
			return
		}
	} else {
		respondWithError(w, http.StatusBadRequest, "No player data found")
		return
	}

	respondWithJSON(w, http.StatusOK, allPlayerStats)
}

// GET /players, get all existing players from db
func (a *App) getAllPlayerData(w http.ResponseWriter, r *http.Request) {
	log.Println("Retrieving all players")
	var allPlayerStats []models.PlayerStat
	var err error

	//retrieves all the players from the db
	allPlayerStats, err = models.GetAllPlayers(a.DB)
	if err != nil {
		log.Println("Could not retrieve player data: ", err.Error())
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve player data: "+err.Error())
		return
	}

	//if there are no players in the db, return no content
	if len(allPlayerStats) == 0 {
		log.Println("No players were found")
		respondWithError(w, http.StatusNoContent, "")
		return
	}

	respondWithJSON(w, http.StatusOK, allPlayerStats)
}

// DELETE /players/delete, clears DB of all player data
func (a *App) deleteAllPlayerData(w http.ResponseWriter, r *http.Request) {
	log.Println("WARNING: Deleting all user data")

	//Deletes all data from the player stats table
	queryStr := "DELETE FROM playerRushingStats;"
	res, err := a.DB.Query(queryStr)
	defer res.Close()
	if err != nil {
		log.Println("Could not delete from db: ", err.Error())
		respondWithError(w, http.StatusInternalServerError, "Could not delete from db: "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"success": "ok"})
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the NFL Rushing Data Service")
}

//Generic function to return whatever error was found as an error json
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

//Generic function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("Could not encode JSON properly")
		respondWithError(w, http.StatusBadRequest, "Could not properly encode JSON: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	log.Println("Starting NFL Rushing Service")

	a := App{}
	a.initialize()
	a.run(":8080")
}
