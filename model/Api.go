package model

type UsernameAndPassword struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type SessionStatus struct {
    Username string `json:"username"`
    LoggedIn bool   `json:"logged_in"`
}
