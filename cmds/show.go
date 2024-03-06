package cmds

import (
	"fmt"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Show() bool {
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

  if len(config.Entries) == 0 {
    lib.Info("There are no entries")
  } else if len(config.Entries) == 1 {
    lib.Info("Listing 1 entry:")
  }else {
    lib.Info("Listing %d entries:", len(config.Entries))
  }

  for _, e := range config.Entries {
    fmt.Printf("    %s%s%s\n", 
      lib.COLOR_BOLD, e.Name, lib.COLOR_RESET)
  }
  return true
}
