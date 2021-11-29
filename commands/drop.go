package commands

import (
	"github.com/sudo-l3wis/vault/types"
)

type DropCommand struct {
	command
}

func (c DropCommand) Action(r Reader, w Writer) {
	name, ok := r.Value(0)
	if !ok {
		w.Write("Incorrect number of arguments.")
		w.Write("Usage: vault drop <name>")
		return
	}
	c.storage.Drop(name)
}
