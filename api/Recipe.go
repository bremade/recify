package api

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/bremade/recify/model"
    "github.com/bremade/recify/util"
)

func (api *Api) RetrieveRecipes(c *gin.Context) {

    recipes, err := api.db.GetAllRecipes()

    if err != nil || len(recipes) <= 0 {
        c.String(http.StatusNotFound, "No recipes found")
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, recipes)
    }
}

func (api *Api) CreateRecipe(c *gin.Context) {

    ok, username := api.CheckLogin(c)

    if !ok {
        c.String(http.StatusForbidden, "User is not logged in")
        return
    }

    var recipeInput model.Recipe
    err := c.BindJSON(&recipeInput)

    newRecipeId := primitive.NewObjectID().Hex()

    // Assign new unique recipe id
    recipeInput.Id = newRecipeId

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    err = api.CheckTags(recipeInput.Tags)
    if err != nil {
        c.String(http.StatusBadRequest, "Error while creating tags")
        return
    }

    err =  api.CheckIngredients(recipeInput.Ingredients)
    if err != nil {
        c.String(http.StatusBadRequest, "Error while creating ingredients")
        return
    }

    recipeInput.Creators = append(recipeInput.Creators, username)

    err = api.db.CreateRecipe(recipeInput)

    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    } else {
        c.String(http.StatusCreated, "OK")
    }
}

func (api *Api) ReplaceRecipe(c *gin.Context) {

    ok, username := api.CheckLogin(c)

    if !ok {
        c.String(http.StatusForbidden, "User is not logged in")
        return
    }

    var recipeInput model.Recipe
    err := c.BindJSON(&recipeInput)

    if err != nil {
        c.String(http.StatusBadRequest, "Bad Request")
        return
    }

    ok, err = api.CheckAuthentication(recipeInput.Id, username)
    if err != nil {
        c.String(http.StatusNotFound, "Recipe with id %v was not found", recipeInput.Id)
        return
    } else if !ok {
        c.String(http.StatusForbidden, "User %v is not the creator of the recipe %v", username, recipeInput.Id)
        return
    }

    err = api.CheckTags(recipeInput.Tags)
    if err != nil {
        c.String(http.StatusBadRequest, "Error while creating tags")
        return
    }

    err =  api.CheckIngredients(recipeInput.Ingredients)
    if err != nil {
        c.String(http.StatusBadRequest, "Error while creating ingredients")
        return
    }

    if !util.ContainsString(recipeInput.Creators, username) {
        recipeInput.Creators = append(recipeInput.Creators, username)
    }

    modifiedCount, err := api.db.ReplaceRecipe(recipeInput)

    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    } else if modifiedCount != 1 {
        c.String(http.StatusInternalServerError, "Modified %v recipes instead of 1", modifiedCount)
    } else {
        c.String(http.StatusOK, "OK")
    }
}

func (api *Api) DeleteRecipe(c *gin.Context) {

    ok, username := api.CheckLogin(c)

    if !ok {
        c.String(http.StatusForbidden, "User is not logged in")
        return
    }

    id := c.Query("id")

    if id == "" {
        c.String(http.StatusInternalServerError, "No recipe id specified")
        return
    }

    ok, err := api.CheckAuthentication(id, username)
    if err != nil {
        c.String(http.StatusNotFound, "Recipe with id %v was not found", id)
        return
    } else if !ok {
        c.String(http.StatusForbidden, "User %v is not the creator of the recipe %v", username, id)
        return
    }

    deletedCount, err := api.db.DeleteRecipe(id)

    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    } else if deletedCount != 1 {
        c.String(http.StatusInternalServerError, "Deleted %v recipes instead of 1", deletedCount)
    } else {
        c.String(http.StatusOK, "OK")
    }
}

func (api *Api) RetrieveIngredients(c *gin.Context) {

    ingredients, err := api.db.GetAllIngredients()

    if err != nil || len(ingredients) <= 0 {
        c.String(http.StatusNotFound, "No ingredients found")
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, ingredients)
    }
}

func (api *Api) RetrieveTags(c *gin.Context) {

    tags, err := api.db.GetAllTags()

    if err != nil || len(tags) <= 0 {
        c.String(http.StatusNotFound, "No tags found")
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, tags)
    }
}

func (api *Api) checkAuthentication(id string, username string) (bool, error) {
    recipe, err := api.db.GetSingleRecipe(id)

    if err != nil {
        return false, err
    }

    return util.ContainsString(recipe.Creators, username), nil
}

func (api *Api) checkIngredients(ingredientsInput []model.Ingredient) error {
    ingredientsDB, err := api.db.GetAllIngredients()

    if err != nil {
        return err
    }

    for _, ingredient := range ingredientsInput {
        if !util.ContainsIngredient(ingredientsDB, ingredient) {
            err = api.db.CreateIngredient(ingredient)
        }
    }

    if err != nil {
        return err
    }

    return nil
}

func (api *Api) checkTags(tagsInput []model.Tag) error {
    tagsDB, err := api.db.GetAllTags()

    if err != nil {
        return err
    }

    for _, tag := range tagsInput {
        if !util.ContainsTag(tagsDB, tag) {
            err = api.db.CreateTag(tag)
        }
    }

    if err != nil {
        return err
    }

    return nil
}

