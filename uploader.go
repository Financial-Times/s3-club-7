package main

import (
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "mime/multipart"
    "os"
)

type Uploader struct {
    Form string
    File multipart.File
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

    _,err = io.Copy(f, u.File)
    if err != nil {
        return
    }

    return
}

func (u *Uploader) CleanUp() (err error) {
    _ = u.tmpFile.Close()    // Just to be safe

    return os.Remove(u.tmpFile.Name())
}
