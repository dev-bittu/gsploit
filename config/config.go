package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var config []byte

var Cfg map[string]interface{}

func init() {
  err := json.Unmarshal(config, &Cfg)
  if err != nil {
    panic("Cant able to load configuration")
  }
}
