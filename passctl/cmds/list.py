from passctl.cmds import Command
from passctl.api import API
from passctl.encryption import Encryption
from passctl.log import error, success 
from random import choice
from getpass import getpass
import readline

class List(Command):
    def __init__(self):
        super().__init__(
            "list", 
            "List the entries in the vault",
        )

    def run(self, cfg, master, args):
        if not "server" in cfg.keys():
            error("Server is not set")

        self.cfg = cfg
        self.args = args
        self.enc = Encryption(master)
        self.api = API(self.cfg["server"], self.enc)

        vault = self.api.get()

        entries = ""
        counter = 0
        for k,v in vault.items():
            counter += 1
            if counter==1:
                entries = k
                continue 
            entries += f", {k}"
        
        success(f"Entries: {entries}")
        return self.cfg
