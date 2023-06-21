from cmds import Command
from api import API
from encryption import Encryption
from log import error, inp, success 
from getpass import getpass
import readline

class Server(Command):
    def __init__(self):
        super().__init__(
            "server", 
            "Set a passctl server",
        )

    def run(self, cfg, master, args):
        self.cfg = cfg
        self.args = args
        self.enc = Encryption(master)

        if len(self.args)<1:
            if not "server" in self.cfg.keys():
                error("URL not specified")
            success("Current server config:")
            success(f"Pass:     {self.cfg['server']['pass']}")
            success(f"Key:      {self.cfg['server']['key']}")
            success(f"URL:      {self.cfg['server']['url']}")
            success("Specify a URL to change the server config")
            return

        url = self.args[0]
        if not (url.startswith("http://") or url.startswith("https://")):
            error("Invalid URL")
        
        try:
            inp("Enter server password:", end="")
            srvpass = getpass(" ")
            inp("Enter server key (leave empty to generate):", end="")
            key = getpass(" ")
        except KeyboardInterrupt:
            error("Keyboard interrupt, quitted", start="\n")
            return

        self.cfg["server"] = {
            "url": url,
            "pass": srvpass,
            "key": key
        }
        
        self.api = API(self.cfg["server"], self.enc)
        if key == "":
            success("Requesting a key from the server")
            res = self.api.gen()
            if res["error"] == 4:
                error("Invalid server password")
            self.cfg["server"]["key"] = res["key"]
        
        success("Server saved to the config")
        return self.cfg 
