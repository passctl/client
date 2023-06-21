from passctl.cmds import Command
from passctl.api import API
from passctl.encryption import Encryption
from passctl.log import error, success 

class Del(Command):
    def __init__(self):
        super().__init__(
            "del",
            "Remove an entry from the vault",
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

        nvault = {}
        for k,v in vault.items():
            if k == entry:
                continue
            nvault[k] = v

        res = self.api.set(nvault) 
        
        if res["error"] == 3:
            error("Can't remove the entry, vault limit exceeded")
        if res["error"] == 4:
            error("Invalid server password")

        success(f"Removed {entry}")
        return self.cfg
