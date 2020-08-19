package auth

import (
    "github.com/gin-contrib/sessions/memstore"
    "github.com/gin-contrib/sessions"
)

func GetSessionStore() memstore.Store {
    return memstore.NewStore([]byte("secret_key"))
}

func (as *AuthService) Login(session sessions.Session, userId string) {
    session.Set("logged_in", true)
    session.Set("user_id", userId)
    session.Save()
}

func (as *AuthService) Logout(session sessions.Session) {
    session.Set("logged_in", false)
    session.Delete("user_id")
    session.Save()
}

func (as *AuthService) GetSessionStatus(session sessions.Session) (bool, string) {
    loggedIn, _ := session.Get("logged_in").(bool)
    userId, _   := session.Get("user_id").(string)
    
    return loggedIn, userId
}
