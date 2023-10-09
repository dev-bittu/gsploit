package console

import (
	"fmt"

	"github.com/dev-bittu/gsploit/config"
)

func (c *Console) help() {
  data, _ := config.Cfg["desc"]
  fmt.Println(data, "\n\n")
  for i, j := range config.Cfg {
    if i != "title" || i != "desc" {
      fmt.Printf("%s : %v", i, j)
    }
  }
}

func (c *Console) getVar(key string) interface{} {
  value, exists := c.vars[key]
  if !exists {
    return nil
  }
  return value
}

func (c *Console) setVar(key string, value interface{}) {
  c.vars[key] = value
}
