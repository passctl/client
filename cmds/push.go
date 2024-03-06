package cmds

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/passctl/client/config"
	"github.com/passctl/client/lib"
)

func Push() bool {
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

  seturl := lib.JoinURL(url, "set?token="+token)
  reader := bytes.NewReader([]byte(config.Cfg.Data))

  lib.Info("Sending the POST request")
  res, err := http.Post(seturl, "text/plain", reader)
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

  lib.Success("Data has been pushed to the server vault")
  return true
}
