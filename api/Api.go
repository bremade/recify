package api

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "github.com/bremade/recify/persistence"
    "github.com/bremade/recify/auth"
)

type Api struct {
    auth *auth.AuthService
}

func New(db *persistence.DB) *Api {
    authService := auth.New(db)
    return &Api{ auth: authService }
}

func (api *Api) Status(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{ "status": "ok" })
}
