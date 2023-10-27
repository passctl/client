from rich.console import Console
from rich.table import Table
from passctl.log import error, success
from passctl.cmds import Command

class Get(Command):
    def __init__(self) -> None:
        super().__init__(
            "get", 
            "Read an entry from the vault",
        )

    def run(self) -> None:
        if len(self.args)<1:
            error("No entry specified")
        
        entry = self.args[0]
        vault = self.stg.data

        if not entry in vault.keys():
            error("Entry not found")

        console = Console()
        table = Table()        
        
        table.add_column("Username")
        table.add_column("Password")
        table.add_row(vault[entry]["user"], vault[entry]["pass"])

        console.print(table)
