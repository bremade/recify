package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/bremade/recify/persistence"
    "github.com/bremade/recify/auth"
)

type Api struct {
    auth *auth.AuthService
    db *persistence.DB
}

func New(db *persistence.DB) *Api {
    authService := auth.New(db)
    return &Api{ auth: authService, db: db }
}

func (api *Api) Status(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{ "status": "ok" })
}

func (api *Api) CheckLogin(c *gin.Context) (bool, string) {
    session := sessions.Default(c)
    loggedIn, userId := api.auth.GetSessionStatus(session)

    // Fetch username
    if loggedIn {
        user, err := api.db.GetUserById(userId)
        if err != nil {
            return false, ""
        }
        return true, user.Name
    } else {
        return false, ""
    }
}
