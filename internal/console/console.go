package console

import (
  alert "github.com/dev-bittu/goalert"
)

type Console struct {
  name string
  prompt string
  vars map[string]interface{}
}

func (c Console) Run() {
  for {
    err := c.handle()
    if err != nil {
      alert.Warning(err.Error())
    }
  }
}

func NewConsole(name string) *Console {
  var vars = make(map[string]interface{})
  vars["verbose"] = false
  return &Console{
    name: name,
    prompt: "gs > ",
    vars: vars,
  }
}
