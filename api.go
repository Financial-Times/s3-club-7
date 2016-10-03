package main

import (
    "log"
    "net/http"
)

type Response struct {
    Time string
    Status int
    Body ResponseBody
}

type ResponseBody struct {
    Success bool
    Message string
}

func Router(w http.ResponseWriter, r *http.Request) {
    LogRequest(r)

    var resp Response
    resp.Status = http.StatusOK
    resp.Body.Success = true

    switch {
    case r.Method == "POST" && r.URL.Path == "/":
        r.ParseMultipartForm(32 << 20)

        uploadFile, handler, err := r.FormFile("upload")
        if err != nil {
            log.Println(err)
            return
        }
        defer uploadFile.Close()

        p := Project{FileName: handler.Filename, UUID: r.FormValue("uuid")}
        u := Uploader{Username: r.FormValue("username"), Password: r.FormValue("password"), Form: r.FormValue("form"), Project: &p}
        u.UploadData(uploadFile)

    default:
        resp.Status = http.StatusNotFound
        resp.Body.Message = "Not found"
        resp.Body.Success = false
    }
}

func LogRequest(r *http.Request) {
    log.Printf( "%s :: %s %s",
        r.RemoteAddr,
        r.Method,
        r.URL.Path)
}
