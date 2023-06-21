import base64
import json
from Crypto import Random
from Crypto.Cipher import AES
import hashlib

class Encryption:
    def __init__(self, master):
        self.bs = AES.block_size
        self.master = hashlib.sha256(master.encode()).digest()

    def _pad(self, s):
        return s + (self.bs - len(s) % self.bs) * chr(self.bs - len(s) % self.bs)

    @staticmethod
    def _unpad(s):
        return s[:-ord(s[len(s)-1:])]

    def encrypt(self, data):
        data = self._pad(data)
        iv = Random.new().read(AES.block_size)
        cipher = AES.new(self.master, AES.MODE_CBC, iv)
        return base64.b64encode(iv+cipher.encrypt(data.encode())).decode("utf-8")

    def decrypt(self, enc):
        if enc == "":
            return "{}"
        enc = base64.b64decode(enc)
        iv = enc[:AES.block_size]
        cipher = AES.new(self.master, AES.MODE_CBC, iv)
        return self._unpad(cipher.decrypt(enc[AES.block_size:])).decode("utf-8")

