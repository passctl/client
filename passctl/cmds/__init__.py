from passctl.config import Config 
from passctl.storage import Storage 

class Command:
    def __init__(self, name: str, desc: str) -> None:
        self.name = name
        self.desc = desc

    def exec(self, cfg: Config, stg: Storage, args: list) -> None:
        self.cfg = cfg
        self.stg = stg
        self.args = args
        return self.run()

    def run(self) -> None:
        return
