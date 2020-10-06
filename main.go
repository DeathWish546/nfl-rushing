package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	//  "sort"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"github.com/DeathWish546/nfl-rushing/lib/models"
	nflUtils "github.com/DeathWish546/nfl-rushing/lib/utils"
)

func main() {
	log.Println("Starting NFL Rushing Service")

	db, err := sql.Open("mysql", "root:password@tcp(db)/nfl")
	defer db.Close()
	if err != nil {
		log.Fatal("Could not initate db: ", err.Error())
		return
	}

	//query := "INSERT INTO playerRushingStats(name, team, position, yards, longest, touchdowns) VALUES (?, ?, ?, ?, ?, ?)"
	//res, err := db.Query(query, "Yes", "No", "Up", 123, 456, 5)
	//defer res.Close()
	//if err != nil {
	//    log.Fatal("Could not get results: ", err.Error())
	//}

	r := mux.NewRouter()

	r.HandleFunc("/", sayHello)
	r.HandleFunc("/1", sayHello2)
	r.HandleFunc("/upload", uploadData).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func uploadData(w http.ResponseWriter, r *http.Request) {
	var allPlayerStats []models.PlayerStat
	allPlayerStats = nflUtils.ParsePlayerData()

	//    sort.SliceStable(allPlayerStats, func(i, j int) bool {
	//        return allPlayerStats[i].Touchdowns < allPlayerStats[j].Touchdowns
	//    })

	w.Header().Set("Content-Type", "application/json")

	respondWithJSON(w, http.StatusOK, allPlayerStats)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World asdflkjasdf")
	log.Println("said hello")
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	res := nflUtils.Wow()
	fmt.Fprint(w, res)
	log.Println("said hello 2")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Could not encode JSON properly")
		respondWithError(w, http.StatusBadRequest, "Could not properly encode JSON: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
