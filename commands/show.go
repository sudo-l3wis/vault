package commands

import (
	"fmt"

	"github.com/sudo-l3wis/vault/types"
)

type ShowCommand struct {
	command
}

func (c ShowCommand) Action(r types.Reader, w types.Writer) {
	name, ok := r.Value(0)
	if !ok {
		w.Write("Invalid number of arguments.")
		w.Write("usage: vault show <name>")
		return
	}

	password, meta := c.storage.Show(name)
	if password.ID == 0 {
		w.Write(fmt.Sprintf("No password with name %s", name))
		return
	}

	if r.Option("r") {
		w.Write(password.Body)
		return
	}

	body := c.cipher.PemToMsg(password.Body)
	decrypted := c.cipher.Decrypt(body)

	w.Write(fmt.Sprintf("password: %s", decrypted))

	for _, m := range meta {
		b := c.cipher.PemToMsg(m.Value)
		d := c.cipher.Decrypt(b)
		w.Write(fmt.Sprintf("%s: %s", m.Name, d))
	}
}
