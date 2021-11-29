package commands

import "github.com/sudo-l3wis/vault/types"

type LsCommand struct {
	command
}

func (c LsCommand) Action(r Reader, w Writer) {
	passwords := c.storage.List()
	for _, p := range passwords {
		w.Write(p.Name)
	}
}
