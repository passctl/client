package cmds

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Pull() bool {
	if config.Cfg.Token == "" {
    lib.Error("You should first run the %ssetup%s command!",
      lib.COLOR_BOLD, lib.COLOR_RESET)
    return false
  } 

  if config.Cfg.Url == "" {
    lib.Error("No server set, first run the %sserver%s command",
      lib.COLOR_BOLD, lib.COLOR_RESET)
  }

  master := GetMaster()
  token, err := lib.Decrypt(master, config.Cfg.Token)
  if err != nil {
    lib.ErrorMsg("Failed to decrypt the token", err)
    return false
  }

  url, err := lib.Decrypt(master, config.Cfg.Url)
  if err != nil {
    lib.ErrorMsg("Failed to decrypt the URL", err)
    return false
  }

  geturl := lib.JoinURL(url, "get?token="+token)
  lib.Info("Sending the GET request")
  res, err := http.Get(geturl)
  if err != nil {
    lib.Error("Request failed, server may be down")
    return false
  }

  defer res.Body.Close()

  if res.StatusCode != 200 {
    lib.Error("Server returned an invalid status code: %d", res.StatusCode)
    return false
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
    lib.ErrorMsg("Failed to read body", err)
    return false
  }

  resjson := struct{
    Error string `json:"error"`
    Data  string `json:"data"`
  }{}

  err = json.Unmarshal(body, &resjson)
  if err != nil {
    lib.ErrorMsg("Failed to parse body", err)
    return false
  }

  if resjson.Error != "" {
    lib.Error("Server returned an error: %s%s", 
      lib.COLOR_BOLD, resjson.Error)
    return false
  }

  config.Cfg.Data = resjson.Data
  err = config.Save()
  if err != nil {
    lib.ErrorMsg("Failed to save the configuration", err)
    return false
  }

  lib.Success("Data has been pulled from the server vault")
  return true
}
