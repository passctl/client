package cmds

import (
	"os"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Add() bool {
  if len(os.Args) != 3 {
    lib.Error("Specify a name for the new entry")
    return false
  }

	if config.Cfg.Token == "" {
    lib.Error("You should first run the %ssetup%s command!",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    return false
  } 

  var entry config.Entry
  entry.Name = os.Args[2]
  if len(entry.Name) < 3 {
    lib.Error("Entry name should be at least %s3%s characters long",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    return false
  }

  master := GetMaster()

USER:
  err := config.LoadEntries(master)
  if err != nil {
    lib.ErrorMsg("Failed to load entries", err)
    return false
  }

  if config.GetEntry(entry.Name)!=nil {
    lib.Error("An entry with the same name exists") 
    return false
  }

  entry.User = lib.Input("Enter a username")
  if entry.User == "" {
    lib.Error("Username cannot be empty")
    goto USER
  }

  lib.Info("Leave the password field empty for auto generation")
  entry.Pwd = lib.Passwd("Enter a password")

  if entry.Pwd == "" {
    entry.Pwd = GenPwd(12)
    lib.Success("Generated password: %s%s", 
      lib.COLOR_BOLD, entry.Pwd)
  }

  config.Entries = append(config.Entries, entry)
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

  lib.Success("Configuration has been saved")
  return true
}
