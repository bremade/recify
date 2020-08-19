package model

type User struct {
    Id           string  `bson:"_id"`
    Name         string  `bson:"name"`
    PasswordHash string  `bson:"passwordhash"`
}
