from log import error, success
from cmds import Command
from api import API
from encryption import Encryption
import json

class Backup(Command):
    def __init__(self):
        super().__init__(
            "backup", 
            "Backup and save your vault localy",
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
        vault = self.api.get()

        try:
            f = open(fn, "w")
            f.write(self.enc.encrypt(json.dumps(vault)))
            f.close()
        except:
            error(f"Can't open '{fn}' for writing")
        
        success(f"Vault saved to '{fn}'")
        return self.cfg
