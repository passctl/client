from cmds import Command
from api import API
from log import error, success 

class Push(Command):
    def __init__(self) -> None:
        super().__init__(
            "push", 
            "Push local vault to the server",
        )

    def run(self) -> None:
        cfg = self.cfg.data
        self.api = API(cfg["server"])

        success(f"Pinging server - {self.api.url}")
        res = self.api.ping()
        if res != "":
            return error(res)
        
        success("Uploading the encrypted vault")
        res = self.api.set(self.stg.encrypt_data())

        if res["error"] != "":
            return error(f"Cannot upload the vault: {res['error']}")

        success("Pushed vault to the server")
