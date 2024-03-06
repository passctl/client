package cmds

import (
	"fmt"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Setup() bool {
	if config.Cfg.Token != "" {
    lib.Error("Seems like you already went through the setup")
    return false
  } 
    
  lib.Info("Running setup, generating a token")
  config.Cfg.Token = lib.GetSHA256(lib.MkRandom(32)) 
 
  fmt.Println()
  lib.Info("Master password will be used to encrypt all of your ohter")
  lib.Info("passwords, if you forget it, there is no way to recover your passwords")

MASTER:
  master := lib.Passwd("Enter a master password")
  if len(master) < 8 {
    lib.Error("Please enter a password that is at least %s8%s characters long!",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    goto MASTER
  }

  config.Cfg.Hash = lib.GetSHA256(master)
  fmt.Println()

  var err error
  config.Cfg.Token, err = lib.Encrypt(
    lib.GetMD5(master), config.Cfg.Token)  

  if err != nil {
    lib.ErrorMsg("Failed to encrypt token", err)
    return false
  }

  err = config.Save()
  if err != nil {
    lib.ErrorMsg("Failed to save configuration", err)
    return false
  }

  lib.Success("Setup completed!")
  lib.Info("Use the %sserver%s command to set a server",
    lib.COLOR_BOLD, lib.COLOR_RESET)

	return true
}
