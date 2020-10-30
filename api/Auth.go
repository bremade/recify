package api

import (
    "fmt"
    "context"

    "github.com/bremade/recify/model"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    oidc "github.com/coreos/go-oidc"
)

func (api *Api) CheckCode(c *gin.Context) {
    ctx := context.Background()

    api.StartOAuth()

    oidcConfig := &oidc.Config{
        ClientID: "recify_app",
    }
    
    verifier := api.provider.Verifier(oidcConfig)
    
    code := c.Query("code")

    fmt.Println("===== Get Auth Token =====")
    fmt.Println("code: " + code)

    oauth2Token, err := api.oauth2Config.Exchange(ctx, code)
    if err != nil {
        c.String(500, "Failed to exchange token: "+err.Error())
        return
    }
    rawIDToken, ok := oauth2Token.Extra("id_token").(string)
    if !ok {
        c.String(500, "No id_token field in oauth2 token.")
        return
    }
    idToken, err := verifier.Verify(ctx, rawIDToken)
    if err != nil {
        c.String(500, "Failed to verify ID Token: "+err.Error())
        return
    }

    var claims struct {
        Username string `json:"email"`
    }

    idToken.Claims(&claims)

    api.auth.Login(sessions.Default(c), claims.Username)

    c.Redirect(302, "/")
}

func (api *Api) AuthStatus(c *gin.Context) {
    session := sessions.Default(c)
    loggedIn, userId := api.auth.GetSessionStatus(session)
    
    c.JSON(200, model.SessionStatus{
        LoggedIn: loggedIn,
        Username: userId,
    })
}

func (api *Api) Logout(c *gin.Context) {
    api.auth.Logout(sessions.Default(c))
    c.Redirect(302, api.providerURL + "/protocol/openid-connect/logout?redirect_uri=http://localhost:8000")
}

func (api *Api) Login(c *gin.Context) {
    api.StartOAuth()
    c.Redirect(302, api.oauth2Config.AuthCodeURL("login-state"))
}

