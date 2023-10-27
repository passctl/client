from passctl.encryption import Encryption
from passctl.util import passdir
from os import path
import json

class Storage:
    def __init__(self, master: str) -> None:
        self.file = path.join(passdir(), "data")
        self.enc = Encryption(master)
        self.data = {}

    def read(self) -> bool:
        try:
            f = open(self.file, "rb")
            self.data = json.loads(self.enc.decrypt(f.read()))
            f.close()
            return True
        except:
            self.data = {}
            return True 

    def write(self) -> bool:
        try:
            f = open(self.file, "w")
            f.write(self.enc.encrypt(json.dumps(self.data)))
            f.close()
            return True
        except:
            return False

    def encrypt_data(self) -> str:
        return self.enc.encrypt(
            json.dumps(self.data)
        )

    def decrypt_data(self, data: bytes) -> str:
        return json.loads(
            self.enc.decrypt(data)
        )
