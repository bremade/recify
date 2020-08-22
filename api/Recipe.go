package api

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

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

    ok, _ := api.CheckLogin(c)

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

    err =  api.CheckIngredients(recipeInput.Ingredients)
    if err != nil {
        c.String(http.StatusBadRequest, "Error while creating ingredients")
        return
    }

    err = api.db.CreateRecipe(recipeInput)

    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    } else {
        c.JSON(http.StatusOK, "OK")
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

    err =  api.CheckIngredients(recipeInput.Ingredients)
    if err != nil {
        c.String(http.StatusBadRequest, "Error while creating ingredients")
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

    if !util.ContainsString(recipeInput.Creators, username) {
        recipeInput.Creators = append(recipeInput.Creators, username)
    }

    modifiedCount, err := api.db.ReplaceRecipe(recipeInput)

    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    } else if modifiedCount != 1 {
        c.String(http.StatusInternalServerError, "Modified %v recipes instead of 1", modifiedCount)
    } else {
        c.JSON(http.StatusOK, "OK")
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
        c.JSON(http.StatusOK, "OK")
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

func (api *Api) CheckLogin(c *gin.Context) (bool, string) {
    session := sessions.Default(c)
    loggedIn, userId := api.auth.GetSessionStatus(session)

    // Fetch username
    if loggedIn {
        user, err := api.db.GetUserById(userId)
        if err != nil {
            return false, ""
        }
        return true, user.Name
    } else {
        return false, ""
    }
}

func (api *Api) CheckAuthentication(id string, username string) (bool, error) {
    recipe, err := api.db.GetSingleRecipe(id)

    if err != nil {
        return false, err
    }

    return util.ContainsString(recipe.Creators, username), nil
}

func (api *Api) CheckIngredients(ingredientsInput []model.Ingredient) error {
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
