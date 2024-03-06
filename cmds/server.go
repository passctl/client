package cmds

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Server() bool {
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

  fmt.Println()
  lib.Info("Use the following token to complete the vault creation process")
  lib.Info("on the server interface, then paste in the vault key")
  lib.InfoVar("Token", token)
  fmt.Println()

  vkey := lib.Passwd("Enter the vault key")
  if vkey == "" {
    lib.Error("Bad vault key")
    return false
  }

  dec, err := lib.Decrypt(lib.GetMD5(token), vkey)
  if err != nil {
    lib.Error("Failed to decrypt the vault key")
    return false
  }

  var srvcfg config.Config
  err = json.Unmarshal([]byte(dec), &srvcfg)
  if err != nil {
    lib.Error("Failed to parse the vault key")
    return false
  }

  if srvcfg.Token != lib.GetMD5(token) {
    lib.Error("Server and client token mismatch")
    return false
  }

  if !strings.HasPrefix(srvcfg.Url, "http://") &&
     !strings.HasPrefix(srvcfg.Url, "https://") {
    lib.Error("Vault key contains a bad URL, probably the server is configured incorrectly")
    return false
  }

  if strings.HasPrefix(srvcfg.Url, "http://") {
    lib.Info("This server is using HTTP, which is insecure, attackers listening the network")
    lib.Info("May compromise your token and get access to the encrypted vault")
  }

  if config.Cfg.Data != "" {
    lib.Info("You may want to run %spush%s to sync your local vault with the server",
      lib.COLOR_BOLD, lib.COLOR_RESET)
  }

  config.Cfg.Url, err = lib.Encrypt(master, srvcfg.Url)
  if err != nil {
    lib.ErrorMsg("Failed to encrypt the configuration", err)
    return false
  }

  err = config.Save()
  if err != nil {
    lib.ErrorMsg("Failed to save the configuration", err)
    return false
  }

  lib.Success("Configuration has been saved")
  return true
}
