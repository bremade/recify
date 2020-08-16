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

func (as *AuthService) RegisterUser(username string, password string) error {
    passwordHash := hashPassword(password)

    user := model.User{
        Id: "",
        Name: username,
        PasswordHash: passwordHash,
    }

    return as.db.CreateUser(user)
}

func (as *AuthService) CheckUser(username string, password string) (bool, int) {
    passwordHash := hashPassword(password)

    // username:password
    if username == "username" && passwordHash == "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8" {
        return true, 1
    } else {
        return false, 0
    }
}

func hashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}
