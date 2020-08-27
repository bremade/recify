package persistence

import (
    "context"
    "fmt"
    "time"

    "github.com/bremade/recify/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DBError struct {
    message string
}

func (err DBError) Error() string {
    return err.message
}

type DB struct {
    dbClient *mongo.Client
    database *mongo.Database
}

func (db *DB) Open(uri string) error {
    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        fmt.Println("error creating mongodb client", err)
        return err
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        fmt.Println("error connecting to mongodb", err)
        return err
    }

    db.dbClient = client

    fmt.Println("mongodb opened")
    return nil
}

func (db *DB) Setup(){
    db.database = db.dbClient.Database("recify")
}

func (db *DB) QueryTest(id int) (model.Test, error) {
    collection := db.database.Collection("Test")
    var result model.Test

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := collection.FindOne(ctx, bson.D{
        {Key: "id", Value: id},
    }).Decode(&result)

    fmt.Println("result:", result)

    return result, err
}

func (db *DB) CreateUser(user model.User) error {
    collection := db.database.Collection("User")
    newUserId := primitive.NewObjectID().Hex()

    // Assign new unique user id
    user.Id = newUserId

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    _, err := collection.InsertOne(ctx, user)

    if err != nil {
        return err
    } else {
        return nil
    }
}

func (db *DB) CheckUser(name string, passwordHash string) (string, error) {
    collection := db.database.Collection("User")

    var user model.User

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := collection.FindOne(ctx, bson.D{
        {Key: "name", Value: name},
        {Key: "passwordhash", Value: passwordHash},
    }).Decode(&user)

    if err != nil {
        return "", err
    } else {
        return user.Id, nil
    }
}

func (db *DB) ContainsUsername(name string) bool {
    collection := db.database.Collection("User")

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    cnt, _ := collection.CountDocuments(ctx, bson.D{
        {Key: "name", Value: name},
    })

    return cnt > 0
}

func (db *DB) GetUserById(id string) (model.User, error) {
    collection := db.database.Collection("User")

    var user model.User

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := collection.FindOne(ctx, bson.D{
        {Key: "_id", Value: id},
    }).Decode(&user)

    if err != nil {
        return user, err
    }

    return user, nil
}

func (db *DB) GetAllRecipes() ([]model.Recipe, error) {
    collection := db.database.Collection("Recipe")

    var recipes []model.Recipe

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})

    if err != nil {
        return recipes, err
    }

    if err = cursor.All(ctx, &recipes); err != nil {
        return recipes, err
    }

    return recipes, nil
}

func (db *DB) GetSingleRecipe(id string) (model.Recipe, error) {
    collection := db.database.Collection("Recipe")

    var recipe model.Recipe

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&recipe)

    if err != nil {
        return recipe, err
    }
    return recipe, nil
}

func (db *DB) CreateRecipe(recipe model.Recipe) error {
    collection := db.database.Collection("Recipe")

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    _, err := collection.InsertOne(ctx, recipe)

    if err != nil {
        return err
    } else {
        return nil
    }
}

func (db *DB) ReplaceRecipe(recipe model.Recipe) (int, error) {
    collection := db.database.Collection("Recipe")

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    result, err := collection.ReplaceOne(ctx, bson.M{"_id": recipe.Id}, recipe)

    if err != nil {
        return 0, err
    } else {
        return int(result.ModifiedCount), nil
    }
}

func (db *DB) DeleteRecipe(id string) (int, error) {
    collection := db.database.Collection("Recipe")

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    result, err := collection.DeleteOne(ctx, bson.M{"_id": id})

    if err != nil {
        return 0, err
    } else {
        return int(result.DeletedCount), nil
    }
}

func (db *DB) GetAllIngredients() ([]model.Ingredient, error) {
    collection := db.database.Collection("Ingredient")

    var ingredients []model.Ingredient

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})

    if err != nil {
        return ingredients, err
    }

    if err = cursor.All(ctx, &ingredients); err != nil {
        return ingredients, err
    }

    return ingredients, nil
}

func (db *DB) CreateIngredient(ingredient model.Ingredient) error {
    collection := db.database.Collection("Ingredient")

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    _, err := collection.InsertOne(ctx, ingredient)

    if err != nil {
        return err
    } else {
        return nil
    }
}

func (db *DB) GetAllTags() ([]model.Tag, error) {
    collection := db.database.Collection("Tag")

    var tags []model.Tag

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})

    if err != nil {
        return tags, err
    }

    if err = cursor.All(ctx, &tags); err != nil {
        return tags, err
    }

    return tags, nil
}

func (db *DB) CreateTag(tag model.Tag) error {
    collection := db.database.Collection("Tag")

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    _, err := collection.InsertOne(ctx, tag)

    if err != nil {
        return err
    } else {
        return nil
    }
}
