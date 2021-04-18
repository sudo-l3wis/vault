package commands

import "github.com/sudo-l3wis/vault/types"

type UsageCommand struct {
	command
}

func (uc UsageCommand) Action(r types.Reader, w types.Writer) {
	w.Write("Command not found. Try `vlt help`.")
}
