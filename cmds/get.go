package cmds

import (
	"os"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Get() bool {
  if len(os.Args) != 3 {
    lib.Error("Specify a name for the entry")
    return false
  }

	if config.Cfg.Token == "" {
    lib.Error("You should first run the %ssetup%s command!",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    return false
  } 

  master := GetMaster()
  err := config.LoadEntries(master)
  if err != nil {
    lib.ErrorMsg("Failed to load entries", err)
    return false
  }

  en := config.GetEntry(os.Args[2])
  if en == nil {
    lib.Error("Entry not found") 
    return false
  }

  lib.InfoVar("Username", en.User)
  lib.InfoVar("Password", en.Pwd)
  return true
}
