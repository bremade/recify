import requests, os, pymongo, hashlib, pytest, json

BASE_URL = os.getenv("BASE_URL") + "/api/v1"

DB_URI = os.getenv("DATABASE_URI")

user = {
    "username": "Sandalenandi",
    "password": "12213443",
}

badUser = {
    "username": "x",
    "password": "x",
    }

session = requests.Session()

class TestAuth():

    def teardown_class(self):
        client = pymongo.MongoClient(DB_URI)
        db = client["recify"]
        col = db["User"]
        query = {
            "name": user["username"],
            "passwordhash": hashlib.sha256(user["password"].encode('utf-8')).hexdigest()
        }

        col.delete_one(query)

        session.cookies.clear()

    def testRegisterUserSuccess(self):
        response = session.post(
            BASE_URL + "/auth/register",
            headers={"Content-Type": "application/json"},
            data=json.dumps(user)
        )
        assert response.status_code == 200
        assert response.text == "Registration OK"

    @pytest.mark.depends(on=['testRegisterUserSuccess'])
    def testRegisterUserFail(self):
        response = session.post(
            BASE_URL + "/auth/register",
            headers={"Content-Type": "application/json"},
            data=json.dumps(user)
        )
        assert response.status_code == 409
        assert response.text == "Username already taken"

    @pytest.mark.depends(on=['testRegisterUserSuccess'])
    def testLoginSuccess(self):
        response = session.post(
            BASE_URL + "/auth/login",
            headers={"Content-Type": "application/json"},
            data=json.dumps(user)
        )
        assert response.status_code == 200
        assert response.text == "OK"

    def testLoginFail(self):
        response = session.post(
            BASE_URL + "/auth/login",
            headers={"Content-Type": "application/json"},
            data=json.dumps(badUser)
        )
        assert response.status_code == 403
        assert response.text == "Login failed"


    @pytest.mark.depends(on=['testStatusSuccess'])
    def testLogoutSuccess(self):
        response = session.get(BASE_URL + "/auth/status")

        respBody = response.json()

        assert response.status_code == 200

    @pytest.mark.depends(on=['testLoginSuccess'])
    def testStatusSuccess(self):
        response = session.get(BASE_URL + "/auth/status")

        respBody = response.json()

        assert response.status_code == 200
        assert respBody["logged_in"]
        assert respBody["username"] == user["username"]

    @pytest.mark.depends(on=['testStatusSuccess'])
    def testStatusFail(self):
        session.cookies.clear()
        response = session.get(BASE_URL + "/auth/status")

        respBody = response.json()

        assert response.status_code == 200
        assert respBody["logged_in"] == False
        assert respBody["username"] == ""
        
