package model

type Time struct {
    Worktime     int          `bson:"worktime"`
    Resttime     int          `bson:"resttime"`
    Cooktime     int          `bson:"cooktime"`
}

type Tag struct {
    Name         string       `bson:"name"`
}

type Step struct {
    Description  string       `bson:"description"`
}

type Ingredient struct {
    Name         string       `bson:"name"`
    Amount       float32      `bson:"amount"`
    Unit         string       `bson:"unit"`
}

type Recipe struct {
    Id           string       `bson:"_id"`
    Creators     []string     `bson:"creators"` 
    Title        string       `bson:"title"`
    Description  string       `bson:"description"`
    Tags         []Tag        `bson:"tags"`
    Time         Time         `bson:"time"`
    Price        float32      `bson:"price"`
    Servings     int          `bson:"servings"`
    Steps        []Step       `bson:"steps"`
    Ingredients  []Ingredient `bson:"ingredients"`
}
