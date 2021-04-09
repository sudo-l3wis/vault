package commands

import (
	"github.com/sudo-l3wis/vault/data"
)

type LsCommand struct {}

func (lc LsCommand) Action(r Reader, w Writer) {
	passwords := data.Storage.List()
	for _, p := range passwords {
		w.Writer.Write(p.Name)
	}
}
