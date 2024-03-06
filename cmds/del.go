package cmds

import (
	"os"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Del() bool {
  if len(os.Args) != 3 {
    lib.Error("Specify a name for the entry")
    return false
  }

	if config.Cfg.Token == "" {
    lib.Error("You should first run the %ssetup%s command!",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    return false
  } 

  name := os.Args[2]
  master := GetMaster()

  err := config.LoadEntries(master)
  if err != nil {
    lib.ErrorMsg("Failed to load entries", err)
    return false
  }

  found := false
  for i, e := range config.Entries {
    if e.Name == name {
      config.Entries = append(config.Entries[:i], config.Entries[i+1:]...)
      found = true
      break
    }
  }

  if !found {
    lib.Error("Entry not found") 
    return false
  }

  err = config.SaveEntries(master)
  if err != nil {
    lib.ErrorMsg("Failed to save entries", err)
    return false
  }

  err = config.Save()
  if err != nil {
    lib.ErrorMsg("Failed to save configuration", err)
    return false
  }

  lib.Success("Entry has been deleted")
  return true
}
