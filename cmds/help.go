package cmds

import (
	"fmt"
	"os"

	"github.com/passctl/client/lib"
)

func Help() bool {
  lib.Info("passctl version %s%s", 
    lib.COLOR_BOLD, lib.GetVersion())
  lib.Info("Usage: %s%s <command>", lib.COLOR_BOLD, os.Args[0])
  lib.Info("Avaliable commands:")

  longest := 0
  for _, c := range List {
    if len(c.Name) > longest {
      longest = len(c.Name)
    }
  }

  for _, c := range List {
    space := ""
    for i := 0; i < longest-len(c.Name); i++ {
      space += " "
    }
    fmt.Printf("%s%s%s%s => %s\n", 
      lib.COLOR_BOLD, c.Name, space, lib.COLOR_RESET, c.Desc)
  }
  return true
}
