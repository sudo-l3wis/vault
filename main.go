package main

import (
	"github.com/sudo-l3wis/vault/commands"
	"github.com/sudo-l3wis/vault/io"
)

func main() {
	r := io.ArgumentReader{}
	w := io.ConsoleWriter{}

	cmd := commands.New()
	cmd.Action(r, w)
}
