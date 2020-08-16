package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)

func (api *Api) Login(c *gin.Context) {
    credentials := make(map[string]string)
    err := c.BindJSON(&credentials)

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    ok, userId := api.auth.CheckUser(credentials["username"], credentials["password"])

    if !ok {
        c.String(http.StatusForbidden, "Login failed")
    } else {
        api.auth.Login(sessions.Default(c), userId)
        c.String(http.StatusOK, "OK")
    }
}

func (api *Api) Register(c *gin.Context) {
    credentials := make(map[string]string)
    err := c.BindJSON(&credentials)

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    err = api.auth.RegisterUser(credentials["username"], credentials["password"])

    if err == nil {
        c.String(http.StatusOK, "Registration OK")
    } else {
        c.String(http.StatusInternalServerError, "Database error")
    }
}

func (api *Api) AuthStatus(c *gin.Context) {
    session := sessions.Default(c)
    loggedIn, userId := api.auth.GetSessionStatus(session)
    
    c.JSON(http.StatusOK, gin.H { "logged_in" : loggedIn, "user_id" : userId })
}
