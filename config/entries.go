package config

import (
	"encoding/json"

	"github.com/passctl/client/lib"
)

var Entries []Entry
type Entry struct {
  Name string  `json:"name"`
  User string  `json:"user"`
  Pwd  string  `json:"pwd"`
}

func GetEntry(name string) *Entry {
  for i, e := range Entries {
    if e.Name == name {
      return &Entries[i]
    }
  }
  return nil
}

func SaveEntries(master string) error {
  tmp := struct {
    List []Entry `json:"entries"`
  }{}

  tmp.List = Entries
  raw, err := json.Marshal(tmp)
  if err != nil {
    return err
  }

  Cfg.Data, err = lib.Encrypt(master, string(raw))
  return err
}

func LoadEntries(master string) error {
  if Cfg.Data == "" {
    Entries = []Entry{}
    return nil
  }

  tmp := struct {
    List []Entry `json:"entries"`
  }{}

  dec, err := lib.Decrypt(master, Cfg.Data)
  if err != nil {
    return err
  }

  err = json.Unmarshal([]byte(dec), &tmp)
  if err != nil {
    return err
  }

  Entries = tmp.List
  return nil 
}
