from cmds import Command
from rich.console import Console
from rich.table import Table
from log import info 
import readline

class List(Command):
    def __init__(self) -> None:
        super().__init__(
            "list", 
            "List the entries in the vault",
        )

    def run(self) -> None:
        vault = self.stg.data
        info(f"Listing total of {len(vault.items())} entries")
        
        console = Console()
        table = Table()                 

        table.add_column("Name")
        table.add_column("Username")
        table.add_column("Password")

        for k, v in vault.items():
            table.add_row(k, v["user"], "*"*len(v["pass"]))
       
        console.print(table)
