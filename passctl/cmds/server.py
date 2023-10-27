from passctl.cmds import Command
from passctl.api import API
from passctl.log import error, inp, success, info
from getpass import getpass
import readline

class Server(Command):
    def __init__(self) -> None:
        super().__init__(
            "server", 
            "Set a passctl server",
        )

    def run(self) -> None:
        cfg = self.cfg.data

        if len(self.args)<1:
            if not "server" in cfg.keys():
                error("URL not specified")
                return

            success("Current server config:")
            success(f"URL:   {cfg['server']['url']}")
            success(f"Key:   {cfg['server']['key']}")
            info("Specify a URL to change the server config")
            return

        url = self.args[0]
        if not (url.startswith("http://") or url.startswith("https://")):
            error("Invalid URL")
        
        try:
            key = inp("Enter vault key", pwd=True)
        except KeyboardInterrupt:
            error("Aborted", start="\n")
            return

        cfg["server"] = {
            "url": url,
            "key": key
        }
        
        self.api = API(cfg["server"])
        res = self.api.ping()
        if res != "":
            error(res)

        self.cfg.data = cfg
        self.cfg.write()
        success("Server saved to the config")
