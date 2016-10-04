package main

import (
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

type Uploader struct {
    Form string
    Password string
    Project *Project
    UUID string
    Username string

    tmpFile *os.File
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

func (u *Uploader) DumpToFS () (err error) {
    if u.tmpFile,err = ioutil.TempFile("", "aintnopartylikeansclubparty"); err != nil {
        return
    }
    defer u.tmpFile.Close()

    f, err := os.OpenFile(u.tmpFile.Name(), os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return
    }
    defer f.Close()

    io.Copy(f, u.tmpFile)

    return
}

func (u *Uploader) CleanUp() (err error) {
    _ = u.tmpFile.Close()    // Just to be safe
    os.Remove(u.tmpFile.Name())

    return

}
