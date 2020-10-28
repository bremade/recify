package model

type Time struct {
    Worktime     int          `bson:"worktime" json:"worktime"`
    Resttime     int          `bson:"resttime" json:"resttime"`
    Cooktime     int          `bson:"cooktime" json:"cooktime"`
}

type Tag struct {
    Name         string       `bson:"name" json:"name"`
}

type Step struct {
    Description  string       `bson:"description" json:"description"`
}

type Ingredient struct {
    Name         string       `bson:"name" json:"name"`
    Amount       float32      `bson:"amount" json:"amount"`
    Unit         string       `bson:"unit" json:"unit"`
}

type Recipe struct {
    Id           string       `bson:"_id" json:"id"`
    Creators     []string     `bson:"creators" json:"creators"` 
    Title        string       `bson:"title" json:"title"`
    Description  string       `bson:"description" json:"description"`
    Tags         []Tag        `bson:"tags" json:"tags"`
    Time         Time         `bson:"time" json:"time"`
    Price        float32      `bson:"price" json:"price"`
    Servings     int          `bson:"servings" json:"servings"`
    Steps        []Step       `bson:"steps" json:"steps"`
    Ingredients  []Ingredient `bson:"ingredients" json:"ingredients"`
}
