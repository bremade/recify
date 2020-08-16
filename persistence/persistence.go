package persistence

import (
    "github.com/bremade/recify/model"
)

type Persistence interface {
    Open(uri string) error
    Setup() error
    QueryTest(id int) model.Test

    CreateUser(user model.User) error
    CheckUser(name string, passwordHash string) (string, error)
    ContainsUsername(name string) bool
}
