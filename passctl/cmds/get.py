from log import error, success
from cmds import Command
from api import API
from encryption import Encryption
import pyperclip

class Get(Command):
    def __init__(self):
        super().__init__(
            "get", 
            "Read an entry from the vault",
        )

    def run(self, cfg, master, args):
        if not "server" in cfg.keys():
            error("Server is not set")
        
        self.cfg = cfg
        self.args = args
        self.enc = Encryption(master)
        self.api = API(self.cfg["server"], self.enc)

        if len(self.args)<1:
            error("No entry specified")
        
        entry = self.args[0]
        vault = self.api.get()

        if not entry in vault.keys():
            error("Entry not found")

        pyperclip.copy(vault[entry]["pass"])
        success(f"Username: {vault[entry]['user']}")
        success(f"Password: {vault[entry]['pass']}")
        success("Password copied to the clipboard")
        return self.cfg
