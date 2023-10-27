from passctl.util import passdir
from os import path
import json

class Config:
    def __init__(self) -> None:
        self.file = path.join(passdir(), "config")
        self.data = {}

    def check(self) -> None:
        if path.exists(self.file) and path.isfile(self.file):
            return
        
        self.write()
    
    def read(self) -> bool:
        try:
            self.check()
            f = open(self.file)
            self.data = json.loads(f.read())
            f.close()
            return True
        except:
            return False

    def write(self) -> bool:
        try:
            f = open(self.file, "w")
            f.write(json.dumps(self.data))
            f.close()
            return True
        except:
            return False
