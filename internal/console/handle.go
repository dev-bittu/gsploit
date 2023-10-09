package console

import (
	"fmt"
	"strings"

	alert "github.com/dev-bittu/goalert"
)

func (c *Console) handle() error {
  var msg string

  fmt.Print(c.prompt)
  fmt.Scanln(&msg)

  msg = strings.ToLower(msg)
  msg = strings.TrimSpace(msg)

  switch msg {
  case "help", "?":
    c.help()
  default:
    alert.Info("No command found: "+msg)
  }

  fmt.Println()

  return nil
}
