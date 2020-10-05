package main

import (
    "fmt"
    "net/http"
    "log"
    "time"

    "github.com/gorilla/mux"

    "github.com/DeathWish546/nfl-rushing/lib/test"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", sayHello)
    r.HandleFunc("/1", sayHello2)

    srv := &http.Server {
        Handler: r,
        Addr: ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World asdflkjasdf")
    log.Println("said hello")
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
    res := test.Wow()
    fmt.Fprint(w, res)
    log.Println("said hello 2")
}
