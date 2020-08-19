package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/bremade/recify/model"
)

func (api *Api) Login(c *gin.Context) {
    var credentials model.UsernameAndPassword
    err := c.BindJSON(&credentials)

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    ok, userId := api.auth.CheckUser(credentials.Username, credentials.Password)

    if !ok {
        c.String(http.StatusForbidden, "Login failed")
    } else {
        api.auth.Login(sessions.Default(c), userId)
        c.String(http.StatusOK, "OK")
    }
}

func (api *Api) Logout(c *gin.Context) {
    api.auth.Logout(sessions.Default(c))
    c.String(http.StatusOK, "OK")
}

func (api *Api) Register(c *gin.Context) {
    var credentials model.UsernameAndPassword
    err := c.BindJSON(&credentials)

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    taken, err := api.auth.RegisterUser(credentials.Username, credentials.Password)

    if taken {
        c.String(http.StatusConflict, "Username already taken")
        return
    }

    if err == nil {
        c.String(http.StatusOK, "Registration OK")
    } else {
        c.String(http.StatusInternalServerError, err.Error())
    }
}

func (api *Api) AuthStatus(c *gin.Context) {
    session := sessions.Default(c)
    loggedIn, userId := api.auth.GetSessionStatus(session)

    // Fetch username
    username := ""
    if loggedIn {
        user, err := api.db.GetUserById(userId)
        if err != nil {
            c.String(http.StatusInternalServerError, err.Error())
            return
        }
        username = user.Name
    }
    
    c.JSON(http.StatusOK, model.SessionStatus{
        LoggedIn: loggedIn,
        Username: username,
    })
}
