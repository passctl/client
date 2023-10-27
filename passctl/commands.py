from passctl.cmds.add import Add
from passctl.cmds.get import Get
from passctl.cmds.list import List
from passctl.cmds.server import Server
from passctl.cmds.delete import Del 
from passctl.cmds.pull import Pull 
from passctl.cmds.push import Push 
from passctl.config import Config
from passctl.storage import Storage 
from passctl.log import error, banner, info, inp 
from getpass import getpass
from rich.table import Table
from rich.console import Console
import hashlib


class Commands:
    def __init__(self):
        self.list = [
            Add(), 
            Get(), 
            Del(), 
            List(),
            Server(),
            Pull(),
            Push()
        ]

    def run(self, name: str, args: list):
        if name == "help":
            self.help()
            return

        cmd = None
        for c in self.list:
            if c.name != name:
                continue
            cmd = c

        if not cmd:
            error(f"Command [dim]{name}[/] not found")
            info("Run [dim]help[/] to get full list of the commands")
            return

        try:
            master = inp("Master password", pwd=True)
        except KeyboardInterrupt:
            error("Aborted", start="\n")
            return

        cfg = Config()
        if not cfg.read():
            error("Cannot load the config")
            return 

        hash = hashlib.sha512()
        hash.update(bytes(master, "utf-8"))
        
        if not "master" in cfg.data.keys():
            cfg.data["master"] = hash.hexdigest()
            cfg.write()
        elif cfg.data["master"] != hash.hexdigest():
            error("Invalid master password")
    
        stg = Storage(master)
        if not stg.read():
            error("Cannot load the storage")

        cmd.exec(cfg, stg, args)

    def help(self):
        banner()

        console = Console()
        table = Table()        

        table.add_column("Command")
        table.add_column("Description")

        for c in self.list:
            table.add_row(c.name, c.desc)

        console.print(table)
