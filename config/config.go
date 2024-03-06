package config 

import (
	"encoding/json"
	"os"
	"path"

	"github.com/passctl/client/lib"
)

var Cfg Config
type Config struct {
  Token  string `json:"token"`
  Data   string `json:"data"`
  Hash   string `json:"hash"`
  Url    string `json:"url"`
}

func Dir() (string, error) {
  home, err := os.UserHomeDir()
  if err != nil {
    return "", err 
  }

  localdir := path.Join(home, ".local")
  sharedir := path.Join(localdir, "share")
  ctldir := path.Join(sharedir, "passctl")
  
  _, err = os.Stat(ctldir)
  if err == nil {
    return ctldir, nil
  }

  err = lib.Mkdirp(localdir)
  if err != nil {
    return "", err
  }

  err = lib.Mkdirp(sharedir)
  if err != nil {
    return "", err
  }

  err = lib.Mkdirp(ctldir)
  if err != nil {
    return "", err
  }

  return ctldir, nil
}

func Read() error {
  pth, err := Dir()
  if err != nil {
    return err
  }  
  
  cfgpth := path.Join(pth, "config")
  raw, err := os.ReadFile(cfgpth)
  if err != nil && !os.IsNotExist(err) {
    return err 
  }

  if err != nil && os.IsNotExist(err) {
    return nil
  }

  return json.Unmarshal(raw, &Cfg)
} 

func Save() error {
  pth, err := Dir()
  if err != nil {
    return err
  }

  cfgpth := path.Join(pth, "config")
  raw, err := json.Marshal(Cfg)
  if err != nil {
    return err
  }

  return os.WriteFile(cfgpth, raw, 0644)
}
