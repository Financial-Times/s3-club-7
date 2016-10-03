package main

import (
    "log"
    "net/http"
)

var clusterid string

func main() {
    clusterid = "jspcdev"
    log.Printf( "Creating projects under the %s cluster", clusterid )
    log.Printf( "Listening on port %d", 8001 )

    http.HandleFunc("/", Router)
    http.ListenAndServe(":8000", nil)
}
