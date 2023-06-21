from passctl.cmds.add import Add
from passctl.cmds.get import Get
from passctl.cmds.list import List
from passctl.cmds.load import Load 
from passctl.cmds.backup import Backup 
from passctl.cmds.server import Server
from passctl.cmds.delete import Del 
from passctl.log import error, success
import hashlib

class Commands:
    def __init__(self):
        self.list = [
            Add(), 
            Get(), 
            Del(), 
            Load(),
            List(),
            Backup(), 
            Server()
        ]

    def run(self, name, cfg, master, args):
        hash = hashlib.sha512()
        hash.update(bytes(master, "utf-8"))
        
        if not "master" in cfg.json.keys():
            cfg.json["master"] = hash.hexdigest()
            cfg.write()
        elif cfg.json["master"] != hash.hexdigest():
            error("Invalid master password")

        for c in self.list:
            if c.name != name:
                continue           
            return c.run(cfg.json, master, args)

        error(f"Command '{name}' not found")

    def help(self):
        success(f"passctl v1.0 | github.com/passctl/client")
        for c in self.list:
            print(f"{c.name}\t\t{c.desc}")
