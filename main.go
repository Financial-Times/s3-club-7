package main

import (
    "flag"
    "log"
    "net/http"

    "github.com/gorilla/securecookie"
)

var blockKey []byte
var clusterid *string
var flexUrl *string
var hashKey []byte

var CookieStore *securecookie.SecureCookie

func init() {
    clusterid = flag.String("cluster", "", "flex cluster to run under")
    flexUrl = flag.String("flex-api", "https://flex.example.com/api", "Flex API Url to validate creds against")

    blockKeyString := flag.String("block-key", "", "AES Key to encrypt secure cookie with. 32 bytes suggested")
    hashKeyString := flag.String("hmac-key", "", "HMAC-specific key. 64 bytes suggested")

    flag.Parse()

    if *clusterid == "" {
        log.Fatal("No flex cluster defined, please specify on the cli")
    }

    if *blockKeyString == "" {
        if blockKey = securecookie.GenerateRandomKey(32); blockKey == nil {
            log.Fatal("Couldn't generate a block key. Suggest specifying one on the cli")
        }
    } else {
        blockKey = []byte(*blockKeyString)
    }

    if *hashKeyString == "" {
        if hashKey = securecookie.GenerateRandomKey(64); hashKey == nil {
            log.Fatal("Couldn't generate an HMACkey. Suggest specifying one on the cli")
        }
    } else {
        hashKey = []byte(*hashKeyString)
    }

    CookieStore = securecookie.New(hashKey, blockKey)
}

func main() {
    log.Printf( "Creating projects under the %s cluster", *clusterid )
    log.Printf( "Authenticating against %s", *flexUrl )
    log.Printf( "Listening on port %d", 8000)

    http.HandleFunc("/", Router)
    http.ListenAndServe(":8000", nil)
}
