package util

import "github.com/bremade/recify/model"

func ContainsString(slice []string, val string) bool {
    for _, tmp := range slice {
        if tmp == val {
            return true
        }
    }
    return false
}

func ContainsIngredient(slice []model.Ingredient, val model.Ingredient) bool {
    for _, tmp := range slice {
        if tmp.Name == val.Name {
            return true
        }
    }
    return false
}
