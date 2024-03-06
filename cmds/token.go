package cmds

import (
	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Token() bool {
	if config.Cfg.Token == "" {
    lib.Error("You should first run the %ssetup%s command!",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    return false
  } 

  master := GetMaster()
  token, err := lib.Decrypt(master, config.Cfg.Token)
  if err != nil {
    lib.ErrorMsg("Failed to decrypt the token", err)
    return false
  }

  lib.InfoVar("Token", token)
  return true
}
