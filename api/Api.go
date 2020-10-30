package api

import (
    "net/http"
    "context"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/bremade/recify/persistence"
    "github.com/bremade/recify/auth"
    "golang.org/x/oauth2"
    oidc "github.com/coreos/go-oidc"
)

type Api struct {
    auth *auth.AuthService
    db *persistence.DB
    providerURL string
    oauth2Config *oauth2.Config
    provider *oidc.Provider
}

func New(db *persistence.DB, providerURL string) *Api {
    authService := auth.New(db)
    return &Api{ auth: authService, db: db, providerURL: providerURL }
}

func (api *Api) StartOAuth() {
    if api.oauth2Config == nil {
        ctx := context.Background()
        provider, err := oidc.NewProvider(ctx, api.providerURL)
        if err != nil {
            panic(err)
        }

        api.provider = provider

        api.oauth2Config = &oauth2.Config{
            ClientID:     "recify_app",
            RedirectURL:  "http://localhost:8000/api/v1/auth/token",
            Endpoint: provider.Endpoint(),
            Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
        }
    }
}

func (api *Api) Status(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{ "status": "ok" })
}

func (api *Api) checkLogin(c *gin.Context) (bool, string) {
    session := sessions.Default(c)
    loggedIn, userId := api.auth.GetSessionStatus(session)

    return loggedIn, userId
}
