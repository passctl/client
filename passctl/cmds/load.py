from passctl.log import error, success
from passctl.cmds import Command
from passctl.api import API
from passctl.encryption import Encryption
import json

class Load(Command):
    def __init__(self):
        super().__init__(
            "load", 
            "Load a localy saved vault",
        )
    def run(self, cfg, master, args):
        if not "server" in cfg.keys():
            error("Server is not set")
        
        self.cfg = cfg
        self.args = args
        self.enc = Encryption(master)
        self.api = API(self.cfg["server"], self.enc)

        if len(self.args)<1:
            error("Filename not specified")
        
        fn = self.args[0]

        try:
            f = open(fn, "r")
            vault = json.loads(self.enc.decrypt(f.read()))
            f.close()
        except:
            error(f"Can't open '{fn}' for reading")
            return

        res = self.api.set(vault)
        if res["error"] == 3:
            error("Can't add the entry, vault limit exceeded")
        if res["error"] == 4:
            error("Server password invalid")
        
        success(f"Loaded the vault from '{fn}'")
        return self.cfg
