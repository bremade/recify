package auth

import (
    "crypto/sha256"
    "encoding/hex"

    "github.com/bremade/recify/persistence"
    "github.com/bremade/recify/model"
)

type AuthService struct {
    db *persistence.DB
}

func New(db *persistence.DB) *AuthService {
    return &AuthService{ db: db }
}

func (as *AuthService) RegisterUser(username string, password string) (bool, error) {
    passwordHash := hashPassword(password)

    if as.db.ContainsUsername(username) {
        return true, nil
    }

    user := model.User{
        Id: "",
        Name: username,
        PasswordHash: passwordHash,
    }

    return false, as.db.CreateUser(user)
}

func (as *AuthService) CheckUser(username string, password string) (bool, string) {
    passwordHash := hashPassword(password)

    userId, err := as.db.CheckUser(username, passwordHash)

    if err != nil {
        return false, ""
    } else {
        return true, userId
    }
}

func hashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}
