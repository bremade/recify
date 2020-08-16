package model

type User struct {
    Id           string  `bson:"_id"`
    Name         string
    PasswordHash string
}
