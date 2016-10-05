package main

import (
    "errors"
    "fmt"
    "net/http"
)

type Auth struct {
    URL string
    Username string
    Password string
}

func (a *Auth) Valid() error {
    var req *http.Request
    var resp *http.Response
    var err error

    client := &http.Client{}
    if req, err = http.NewRequest("GET", a.URL, nil); err != nil {
        return err
    }

    req.SetBasicAuth(a.Username, a.Password)

    if resp, err = client.Do(req); err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return errors.New(fmt.Sprintf("Authenticating against %s returned '%s'", a.URL, resp.Status))
    }

    return nil
}
