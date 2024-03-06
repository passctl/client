package cmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Search() bool {
  if len(os.Args) != 3 {
    lib.Error("Specify a name to search for")
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

  var results []config.Entry
  for _, e := range config.Entries {
    if strings.Contains(e.Name, os.Args[2]) {
      results = append(results, e)
    }
  }

  if len(results) == 0 {
    lib.Error("Got no matches")
    return false
  }

  if len(results) == 1 {
    lib.Success("Got exactly one match, showing entry")
    lib.InfoVar("Name    ", results[0].Name)
    lib.InfoVar("Username", results[0].User)
    lib.InfoVar("Password", results[0].Pwd)
    return true
  }

  lib.Info("Listing %d entries:", len(results))
  for _, e := range results {
    fmt.Printf("    %s%s%s\n", 
      lib.COLOR_BOLD, e.Name, lib.COLOR_RESET)
  }
 
  return true
}
