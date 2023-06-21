from passctl.commands import Commands
from passctl.log import error, inp
from getpass import getpass
from passctl.config import Config
from sys import argv
import readline

def main():
    cfg = Config()
    cmds = Commands()

    try:
        cfg.read()
    except:
        error("Can't load the config")

    inp("Master password:", end="")
    try:
        master = getpass(" ")
    except KeyboardInterrupt:
        error("Keyboard interrupt, quitted", start="\n")
        return

    if len(argv) < 2:
        cmds.help()
        exit()
    
    cmd = argv[1]
    newcfg = cmds.run(cmd, cfg, master, argv[2:])
    if newcfg != None:
        cfg.write()

if __name__ == "__main__":
    main()
