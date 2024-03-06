package cmds

import (
	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

var List []Cmd
type CmdFunc func() bool 
type Cmd struct {
  Name string
  Desc string
  Func CmdFunc 
}

func LoadCmds() {
  List = []Cmd{
    Cmd{
      Name: "help",
      Desc: "Show this information",
      Func: Help,
    },
    Cmd{
      Name: "setup",
      Desc: "Setup the client",
      Func: Setup,
    },
    Cmd{
      Name: "token",
      Desc: "Print your client token",
      Func: Token,
    },
    Cmd{
      Name: "server",
      Desc: "Use a server for the vault",
      Func: Server,
    },
    Cmd{
      Name: "push",
      Desc: "Push local vault to server",
      Func: Push,
    },
    Cmd{
      Name: "pull",
      Desc: "Pull server vault to local",
      Func: Pull,
    },
    Cmd{
      Name: "gen",
      Desc: "Generate a random password",
      Func: Gen,
    },
    Cmd{
      Name: "add",
      Desc: "Add an entry to the vault",
      Func: Add,
    },
    Cmd{
      Name: "get",
      Desc: "Get data of an entry",
      Func: Get,
    },
    Cmd{
      Name: "del",
      Desc: "Delete an entry",
      Func: Del,
    },
    Cmd{
      Name: "list",
      Desc: "List all entries",
      Func: Show,
    },
    Cmd{
      Name: "search",
      Desc: "Search for an entries by name",
      Func: Search,
    },
  }
}

func GetMaster() string {
START:
  master := lib.Passwd("Enter the master password")
  if len(master) < 8 {
    lib.Error("Incorrect master password, try again")
    goto START;
  }

  if config.Cfg.Hash == lib.GetSHA256(master) {
    lib.Success("Master password is correct")
    return lib.GetMD5(master)
  }
    
  lib.Error("Incorrect master password, try again")
  goto START
}

