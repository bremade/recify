package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/bremade/recify/auth"
)

func Login(c *gin.Context) {
    auth.Login(sessions.Default(c), 0)
    c.JSONP(http.StatusOK, gin.H{ "status": "ok" })
}

func AuthStatus(c *gin.Context) {
    session := sessions.Default(c)
    c.JSON(http.StatusOK, auth.GetSessionStatus(session))
}

