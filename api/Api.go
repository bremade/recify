package api

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "github.com/bremade/recify/persistence"
)

type Api struct {
    db *persistence.DB
}

func New(db *persistence.DB) *Api {
    return &Api{ db: db }
}

func (api *Api) Status(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{ "status": "ok" })
}
