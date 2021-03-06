import requests, os, pymongo, hashlib, pytest, json

BASE_URL = os.getenv("BASE_URL") + "/api/v1"

DB_URI = os.getenv("DATABASE_URI")

user = {
    "username": "Testuser",
    "password": "13245678",
}

tag = [
    {
    "name": "Testag"
    }
]

ingredient = [
    {
    "name": "Testingredient",
    "amount": 1,
    "unit": "mg"
    }
]

recipe = {
    "title": "Testtitle",
    "description": "Testdescription",
    "tags": tag,
    "time": {
        "worktime": 1,
        "resttime": 2,
        "cooktime": 3
    },
    "price": 1,
    "servings": 1,
    "steps": [
        {
        "description": "Teststep"
        }
    ],
    "ingredients": ingredient
}

session = requests.Session()

client = pymongo.MongoClient(DB_URI)
db = client["recify"]
recipeId = "tofill"

class TestAuth():

    def teardown_class(self):
        db["User"].drop()
        db["Recipe"].drop()
        db["Tag"].drop()
        db["Ingredient"].drop()
        session.cookies.clear()

    def setup_class(self):
        responseRegister = session.post(
            BASE_URL + "/auth/register",
            headers={"Content-Type": "application/json"},
            data=json.dumps(user)
        )
        assert responseRegister.status_code == 200

        responseLogin = session.post(
            BASE_URL + "/auth/login",
            headers={"Content-Type": "application/json"},
            data=json.dumps(user)
        )
        assert responseLogin.status_code == 200

    def testGetRecipesFail(self):
        response = session.get(
            BASE_URL + "/recipe"
        )
        assert response.status_code == 404
        assert response.text == "No recipes found"
    
    def testGetIngredientsFail(self):
        response = session.get(
            BASE_URL + "/ingredient"
        )
        assert response.status_code == 404
        assert response.text == "No ingredients found"
    
    def testGetTagsFail(self):
        response = session.get(
            BASE_URL + "/tag"
        )
        assert response.status_code == 404
        assert response.text == "No tags found"

    @pytest.mark.depends(on=['testGetRecipesFail', 'testGetIngredientsFail', 'testGetTagsFail'])
    def testCreateRecipeSuccess(self):
        response = session.post(
            BASE_URL + "/recipe",
            headers={"Content-Type": "application/json"},
            data=json.dumps(recipe)
        )
        assert response.status_code == 201
        assert response.text == "OK"

    @pytest.mark.depends(on=['testCreateRecipeSuccess'])
    def testGetRecipesSuccess(self):
        global recipeId

        response = session.get(
            BASE_URL + "/recipe"
        )
        body = response.json()

        recipeId = body[0]["id"]
        recipe["id"] = recipeId
        recipe["creators"] = [
            "Testuser"
        ]

        recipeList = [
            recipe
        ]

        assert response.status_code == 200
        assert body == recipeList
    
    @pytest.mark.depends(on=['testCreateRecipeSuccess'])
    def testGetIngredientsSuccess(self):
        response = session.get(
            BASE_URL + "/ingredient"
        )

        body = response.json()
        assert response.status_code == 200
        assert body == ingredient
    
    @pytest.mark.depends(on=['testCreateRecipeSuccess'])
    def testGetTagsSuccess(self):
        response = session.get(
            BASE_URL + "/tag"
        )

        body = response.json()
        assert response.status_code == 200
        assert body == tag

    @pytest.mark.depends(on=['testGetRecipesSuccess'])
    def testReplaceRecipeSuccess(self):
        global recipeId

        recipeCopy = recipe
        recipeCopy["id"] = recipeId
        recipeCopy["title"] = "UpdatedTitle"

        response = session.put(
            BASE_URL + "/recipe",
            headers={"Content-Type": "application/json"},
            data=json.dumps(recipe)
        )
        assert response.status_code == 200
        assert response.text == "OK"

    def testReplaceRecipeFail(self):
        response = session.put(
            BASE_URL + "/recipe",
            headers={"Content-Type": "application/json"},
            data=json.dumps(recipe)
        )
        assert response.status_code == 404
        assert response.text == "Recipe with id  was not found"

    @pytest.mark.depends(on=['testReplaceRecipeSuccess'])
    def testGetRecipesAfterReplace(self):
        global recipeId

        response = session.get(
            BASE_URL + "/recipe"
        )

        body = response.json()
        recipeCopy = recipe

        recipeId = body[0]["id"]
        recipeCopy["id"] = recipeId
        recipeCopy["title"] = "UpdatedTitle"
        recipeCopy["creators"] = [
            "Testuser"
        ]

        recipeList = [
            recipeCopy
        ]

        assert response.status_code == 200
        assert body == recipeList

    @pytest.mark.depends(on=['testReplaceRecipeSuccess', 'testGetRecipesAfterReplace'])
    def testDeleteRecipeSuccess(self):
        global recipeId

        response = session.delete(
            BASE_URL + "/recipe",
            params={"id": recipeId}
        )
        assert response.text == "OK"
        assert response.status_code == 200
        assert response.text == "OK"

    def testDeleteRecipeFail(self):
        response = session.delete(
            BASE_URL + "/recipe",
            params={"id": "badId"}
        )
        assert response.status_code == 404
        assert response.text == "Recipe with id badId was not found"

    @pytest.mark.depends(on=['testReplaceRecipeSuccess', 'testDeleteRecipeSuccess'])
    def testCreateRecipeNotLoggedIn(self):
        session.cookies.clear()
        response = session.post(
            BASE_URL + "/recipe",
            headers={"Content-Type": "application/json"},
            data=json.dumps(recipe)
        )
        assert response.status_code == 403
        assert response.text == "User is not logged in"

    @pytest.mark.depends(on=['testReplaceRecipeSuccess'])
    def testGetRecipe(self):
        global recipeId
        
        response = session.get(
            BASE_URL + "/recipe/" + recipeId
        )

        body = response.json()
        response = session.get(
            BASE_URL + "/recipe/" + recipeId
        )

        recipeCopy = recipe
        recipeCopy["id"] = recipeId
        recipeCopy["creators"] = [
            "Testuser"
        ]

        assert response.status_code == 200
        assert body == recipeCopy
