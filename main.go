package main

import "github.com/sudo-l3wis/vault/commands"

func main() {
	r := ArgumentReader{}
	w := ConsoleWriter{}

	cmd := commands.MakeCommand()
	cmd.Action(r, w)
}
