from cmds import Command
from api import API
from log import error, success 

class Pull(Command):
    def __init__(self) -> None:
        super().__init__(
            "pull", 
            "Pull the vault from the server",
        )

    def run(self) -> None:
        cfg = self.cfg.data
        self.api = API(cfg["server"])

        success(f"Pinging server - {self.api.url}")
        res = self.api.ping()
        if res != "":
            return error(res)
 
        success("Downloading the encrypted vault")
        res = self.api.get()

        if res["error"] != "":
            return error(f"Cannot download the vault: {res['error']}")

        self.stg.data = self.stg.decrypt_data(res["vault"])
        self.stg.write()
        success("Pulled vault from the server")
