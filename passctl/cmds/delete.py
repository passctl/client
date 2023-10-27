from passctl.cmds import Command
from passctl.log import error, success 

class Del(Command):
    def __init__(self) -> None:
        super().__init__(
            "del",
            "Remove an entry from the vault",
        )

    def run(self) -> None:
        if len(self.args)<1:
            error("No entry specified")

        entry = self.args[0]
        vault = self.stg.data

        if not entry in vault.keys():
            error("Entry not found")

        nvault = {}
        for k,v in vault.items():
            if k == entry:
                continue
            nvault[k] = v

        self.stg.data = vault
        if not self.stg.write():
            error("Cannot save the vault")
        
        success(f"Removed {entry}")
