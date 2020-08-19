import requests, os

BASE_URL = os.getenv("BASE_URL") + "/api/v1"

def testStatus():
    response = requests.get(BASE_URL + "/status")
    respBody = response.json()
    assert response.status_code == 200
    assert respBody["status"] == "ok"
