package main

import (
    "flag"
    "log"
    "net/http"
)

var clusterid *string
var flexUrl *string

func init() {
    clusterid = flag.String("cluster", "", "flex cluster to run under")
    flexUrl = flag.String("flex-api", "https://flex.example.com/api", "Flex API Url to validate creds against")

    flag.Parse()

    if *clusterid == "" {
        log.Fatal("No flex cluster defined, please specify on the cli")
    }
}

func main() {
    log.Printf( "Creating projects under the %s cluster", *clusterid )
    log.Printf( "Authenticating against %s", *flexUrl )
    log.Printf( "Listening on port %d", 8000)

    http.HandleFunc("/", Router)
    http.ListenAndServe(":8000", nil)
}
