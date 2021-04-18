package commands

import "github.com/sudo-l3wis/vault/types"

type KeyCommand struct {
	command
}

func (c KeyCommand) Action(r types.Reader, w types.Writer) {
	name, ok := r.Value(0)
	if !ok {
		w.Write("Incorrect number of arguments.")
		w.Write("Usage: vault key <public|private>")
		return
	}

	private, public := c.cipher.GetKeys()

	if name == "public" {
		w.Write(public)
	} else if name == "private" {
		w.Write(private)
	} else {
		w.Write("Invalid option")
	}
}
