package main

import (
    "errors"
    "fmt"
)

type Uploader struct {
    Username string
    Password string
    Form string
    UUID string
    Project *Project
}

func (u *Uploader) BucketName() (b string, err error) {
    var bucketItem string
    switch u.Form {
    case "ingest":
        bucketItem = "raw"
    case "publish":
        bucketItem = "edit"
    default:
        err = errors.New( fmt.Sprintf("Form: %s is not recognised", u.Form) )
        return
    }

    b = fmt.Sprintf("mio-%s-%s", bucketItem, clusterid)
    return
}
