from cmds import Command
from api import API
from encryption import Encryption
from log import error, inp, success 
import string
import pyperclip
from random import choice
from getpass import getpass
import readline

class Add(Command):
    def __init__(self):
        super().__init__(
            "add", 
            "Add an entry to the vault",
        )

    def run(self, cfg, master, args):
        if not "server" in cfg.keys():
            error("Server is not set")

        self.cfg = cfg
        self.args = args
        self.enc = Encryption(master)
        self.api = API(self.cfg["server"], self.enc)

        if len(self.args)<1:
            error("Entry name not specified")

        entry = self.args[0]
        vault = self.api.get()

        if entry in vault.keys():
            error("Entry already exists")
       
        try:
            success("Creating the entry, leave the password empty for auto generation")
            inp("Enter username:", end="")
            usr = input(" ")
            inp("Enter password:", end="")
            pwd = getpass(" ")
        except KeyboardInterrupt:
            error("Keyboard interrupt, quitted", start="\n")
            return 

        if pwd == "":
            chars = string.ascii_letters+string.digits+"!?=%.:$+-"
            for _ in range(20):
                pwd += choice(chars)
            success("Auto generated password copied to the clipboard")
            pyperclip.copy(pwd)
        
        vault[entry] = {"user": usr, "pass":pwd}
        res = self.api.set(vault)
        
        if res["error"] == 3:
            error("Can't add the entry, vault limit exceeded")
        if res["error"] == 4:
            error("Invalid server password")

        success(f"Added {entry}")
        return self.cfg
