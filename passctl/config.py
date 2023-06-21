from os import path
import json

class Config:
    def __init__(self):
        self.file = path.expanduser("~/.passctl")
        self.json = {}

    def check(self):
        if path.exists(self.file) and path.isfile(self.file):
            return 
        
        self.write()
    
    def read(self):
        self.check()
        f = open(self.file)
        self.json = json.loads(f.read())
        f.close()

    def write(self):
        f = open(self.file, "w")
        f.write(json.dumps(self.json))
        f.close()
