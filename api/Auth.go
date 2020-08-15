package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/bremade/recify/auth"
)

func (api *Api) Login(c *gin.Context) {
    credentials := make(map[string]string)
    err := c.BindJSON(&credentials)

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    ok, userId := auth.CheckUser(credentials["username"], credentials["password"])

    if !ok {
        c.String(http.StatusForbidden, "Login failed")
    } else {
        auth.Login(sessions.Default(c), userId)
        c.String(http.StatusOK, "OK")
    }
}

func (api *Api) AuthStatus(c *gin.Context) {
    session := sessions.Default(c)
    loggedIn, userId := auth.GetSessionStatus(session)
    
    c.JSON(http.StatusOK, gin.H { "logged_in" : loggedIn, "user_id" : userId })
}
