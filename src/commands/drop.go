package commands

import (
	"github.com/sudo-l3wis/vault/data"
)

type DropCommand struct {}

func (dc DropCommand) Action(r Reader, w Writer) {
	name, ok := r.Value(0)
	if !ok {
		w.Write("Incorrect number of arguments.")
		W.Write("Usage: vault drop <name>")
		return
	}
	data.Storage.Drop(name)
}
