from passctl.commands import Commands
from sys import argv

def main():
    cmds = Commands()
        
    if len(argv) == 1:
        cmds.help()
        exit()
    
    cmd = argv[1]
    cmds.run(cmd, argv[2:])

if __name__ == "__main__":
    main()
