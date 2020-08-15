package auth

import (
    "github.com/gin-contrib/sessions/memstore"
    "github.com/gin-contrib/sessions"
)

type SessionStatus struct {
    LoggedIn bool `json:"logged_in""`
    UserId int    `json:"user_id"`
}

func GetSessionStore() memstore.Store {
    return memstore.NewStore([]byte("secret_key"))
}

func Login(session sessions.Session, userId int) {
    session.Set("logged_in", true)
    session.Set("user_id", userId)
    session.Save()
}

func GetSessionStatus(session sessions.Session) SessionStatus {
    loggedIn, _ := session.Get("logged_in").(bool)
    userId, _   := session.Get("user_id").(int)
    
    return SessionStatus {
        LoggedIn: loggedIn,
        UserId: userId,
    }
}
