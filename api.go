package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "time"
)

type Response struct {
    Referer string
    Time string
    Status int
    Body ResponseBody
}

type ResponseBody struct {
    Success bool
    Message string
}

const (
    cookieAge = 43200
    cookieName = "s3-club-7"
)


func Router(w http.ResponseWriter, r *http.Request) {
    loggedIn, username := isLoggedIn(r)
    LogRequest(r, username)

    var resp Response
    resp.Status = http.StatusOK
    resp.Time = time.Now().Format(time.RFC3339)
    resp.Body.Success = true

    if r.Referer() == ""{
        resp.Referer = "null"
    } else {
        resp.Referer = refererDomain(r.Referer())
    }

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
            return
        } else {
            http.SetCookie(w, setLogin(a.Username))
            resp.Body.Message = "logged in"
        }


    case r.Method == "GET" && r.URL.Path == "/session":
        if loggedIn {
            resp.Body.Message = "logged in"
        } else {
            resp.Status = http.StatusUnauthorized
            resp.Body.Message = "not logged in"
            resp.Body.Success = false
        }


    case r.Method == "POST" && r.URL.Path == "/upload":
        if !loggedIn {
            resp.Status = http.StatusUnauthorized
            resp.Body.Message = "not logged in"
            resp.Body.Success = false

            resp.respond(w)
            return
        }

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

func LogRequest(r *http.Request, username string) {
    log.Printf( "%s@%s :: %s %s",
        username,
        r.RemoteAddr,
        r.Method,
        r.URL.Path)
}

func (r Response) respond (w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "requested-with, Content-Type, origin, authorization, accept, client-security-token, cache-control, Set-Cookie, Cookie")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
    w.Header().Set("Access-Control-Allow-Origin", r.Referer)
    w.Header().Set("Access-Control-Max-Age", "10000")
    w.Header().Set("Cache-Control", "no-cache")

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(r.Status)
    j,_ := json.Marshal(r)

    fmt.Fprintf(w, string(j))
}

func setLogin(username string)(c *http.Cookie) {
    value := map[string]string{
        "loggedin": "true",
        "username": username,
    }
    if encoded, err := CookieStore.Encode(cookieName, value); err == nil {
        c = &http.Cookie{
            MaxAge: cookieAge,
            Name:  cookieName,
            Path:  "/",
            Secure: !*development,
            Value: encoded,
        }
    } else {
        log.Println(err)
    }

    return
}

func isLoggedIn(r *http.Request)(loggedIn bool, username string) {
    if cookie, err := r.Cookie(cookieName); err == nil {
        value := make(map[string]string)

        if err = CookieStore.Decode(cookieName, cookie.Value, &value); err == nil {
            if value["loggedin"] == "true" {
                return true, value["username"]
            }
        }
    }

    return
}

func refererDomain(s string)(domain string) {
    var u *url.URL
    if u, err = url.Parse(s); err != nil {
        return s
    }

    return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}
