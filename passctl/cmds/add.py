from cmds import Command
from log import error, inp, success, info 
from random import choice
from getpass import getpass
import readline
import string

class Add(Command):
    def __init__(self) -> None:
        super().__init__(
            "add", 
            "Add an entry to the vault",
        )

    def run(self) -> None:
        if len(self.args)<1:
            error("Entry name not specified")

        entry = self.args[0]
        vault = self.stg.data

        if entry in vault.keys():
            error("Entry already exists")
       
        try:
            info("Leave the password empty for auto generation")
            usr = inp("Username")
            pwd = inp("Password", pwd=True)
        except KeyboardInterrupt:
            error("Aborted", start="\n")
            return 

        if pwd == "":
            chars = string.ascii_letters+string.digits+"!?=%.:$+-"
            for _ in range(20):
                pwd += choice(chars)
            success(f"Auto generated password: {pwd}")
        
        vault[entry] = {
            "user": usr, 
            "pass":pwd
        }
        success(f"Added {entry}")

        self.stg.data = vault
        if not self.stg.write():
            error("Cannot save the vault")
