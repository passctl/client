package main

import (
	"os"

	"github.com/passctl/client/cmds"
	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func main() {
  cmds.LoadCmds()
  config.Read()

  if len(os.Args) < 2 {
    cmds.Help()
    return
  }

  for _, c := range cmds.List {
    if c.Name == os.Args[1] {
      c.Func()
      return
    }
  }

  lib.Error("Command not found: %s", os.Args[1])
}
