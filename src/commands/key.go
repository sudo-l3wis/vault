package commands

import (
	"github.com/sudo-l3wis/vault/crypt"
)

type KeyCommand struct {}

func (kc KeyCommand) Action(r Reader, w Writer) {
	name, ok := r.Value(0)
	if !ok {
		w.Write("Incorrect number of arguments.")
		w.Write("Usage: vault key <public|private>")
		return
	}

	private, public := crypt.Keys.GetKeys()

	if name == "public" {
		w.Write(public)
	} else if name == "private" {
		w.Write(private)
	} else {
		w.Write("Invalid option")
	}
}
