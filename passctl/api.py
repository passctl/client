import requests
import json

class API:
    def __init__(self, server: dict) -> None:
        self.url = server["url"]
        self.key = server["key"]

    def ping(self) -> str:
        try:
            res = requests.get(self.url+"/api/ping/"+self.key)
        except:
            return "Cannot reach the server"

        try:
            res = json.loads(res.text)
            if res["error"]=="":
                return ""
            return "Invalid server vault key"
        except:
            return "Bad response from the server" 

    def get(self) -> dict:
        res = requests.get(self.url+"/api/get/"+self.key)
        return res.json() 

    def set(self, data: str) -> dict:
        res = requests.post(self.url+"/api/set/"+self.key, json={
            "vault": data
        })
        return res.json() 
