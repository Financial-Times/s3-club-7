package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
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
    resp.Time = time.Now().Format(time.RFC3339)
    resp.Body.Success = true

    r.ParseMultipartForm(32 << 20)

    switch {
    case r.Method == "OPTIONS":


    case r.Method == "POST" && r.URL.Path == "/session":
        a := Auth{Username: r.FormValue("username"), Password: r.FormValue("password"), URL: *flexUrl}
        if err := a.Valid(); err != nil {
            resp.Status = http.StatusUnauthorized
            resp.Body.Message = err.Error()
            resp.Body.Success = false

            resp.respond(w)
        } else {
            http.SetCookie(w, setLogin())
            resp.Body.Message = "logged in"
        }


    case r.Method == "GET" && r.URL.Path == "/session":
        if isLoggedIn(r) {
            resp.Body.Message = "logged in"
        } else {
            resp.Status = http.StatusUnauthorized
            resp.Body.Message = "not logged in"
            resp.Body.Success = false
        }


    case r.Method == "POST" && r.URL.Path == "/upload":
        r.ParseMultipartForm(32 << 20)

        uploadFile, handler, err := r.FormFile("upload")
        if err != nil {
            log.Println(err)
            return
        }
        defer uploadFile.Close()

        p := Project{FileName: handler.Filename, UUID: r.FormValue("uuid")}
        u := Uploader{Form: r.FormValue("form"), Project: &p, File: uploadFile}

        if err := u.DumpToFS(); err != nil {
            resp.Status = http.StatusInternalServerError
            resp.Body.Message = err.Error()
            resp.Body.Success = false

            resp.respond(w)
            return
        }

        defer u.CleanUp()

        if status, err := u.UploadData(); err != nil {
            resp.Status = http.StatusInternalServerError
            resp.Body.Message = err.Error()
            resp.Body.Success = false
        } else {
            resp.Body.Message = status
            resp.Body.Success = true
        }
        if err := u.CleanUp(); err != nil {
            log.Printf("Error cleaning up: %s", err.Error())
        }


    default:
        resp.Status = http.StatusNotFound
        resp.Body.Message = "Not found"
        resp.Body.Success = false
    }

    resp.respond(w)
    return
}

func LogRequest(r *http.Request) {
    log.Printf( "%s :: %s %s",
        r.RemoteAddr,
        r.Method,
        r.URL.Path)
}

func (r Response) respond (w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Headers", "requested-with, Content-Type, origin, authorization, accept, client-security-token, cache-control, x-api-key")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Max-Age", "10000")
    w.Header().Set("Cache-Control", "no-cache")

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(r.Status)
    j,_ := json.Marshal(r)

    fmt.Fprintf(w, string(j))
}

func setLogin()(c *http.Cookie) {
    value := map[string]string{
        "loggedin": "true",
    }
    if encoded, err := CookieStore.Encode("s3-club-7", value); err == nil {
        c = &http.Cookie{
            Name:  "s3-club-7",
            Value: encoded,
            Path:  "/",
        }
    }

    return c
}

func isLoggedIn(r *http.Request)(bool) {
    if cookie, err := r.Cookie("s3-club-7"); err == nil {
        value := make(map[string]string)

        if err = CookieStore.Decode("s3-club-7", cookie.Value, &value); err == nil {
            if value["loggedin"] == "true" {
                return true
            }
        }
    }

    return false
}
