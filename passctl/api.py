from passctl.log import error
import requests
import json

class API:
    def __init__(self, server, enc):
        self.url = server["url"]
        self.pw = server["pass"]
        self.key = server["key"]
        self.enc = enc

    def ping(self):
        res = requests.get(self.url+"/api/ping")
        try:
            res = json.loads(res.text)
            if res["error"]!=0:
                return False
            return True
        except:
            return False

    def get(self):
        res = requests.get(self.url+"/api/get/"+self.key)
        res = json.loads(res.text)
        if res["error"] != 0:
            error("Vault does not exists, this means your vault key has been deleted, update your server config")

        return json.loads(
            self.enc.decrypt(res["vault"])
        )

    def gen(self):
        res = requests.post(self.url+"/api/gen", json={"pass":self.pw})
        return json.loads(res.text) 

    def set(self, data):
        res = requests.post(self.url+"/api/set/"+self.key, json={
            "pass": self.pw, 
            "vault": self.enc.encrypt(json.dumps(data))
        })
        return json.loads(res.text)
