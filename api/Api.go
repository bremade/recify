package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{ "status": "ok" })
}
