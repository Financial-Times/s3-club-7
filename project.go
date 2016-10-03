package main

import (
    "fmt"
)

type Project struct {
    FileName string
    UUID string
}

func (p *Project)Key() string{
    return fmt.Sprintf("%s/%s", p.UUID, p.FileName)
}
