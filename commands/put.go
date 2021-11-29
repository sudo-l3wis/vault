package commands

import (
	"fmt"

	"github.com/sudo-l3wis/vault/crypt"
	"github.com/sudo-l3wis/vault/types"
)

type PutCommand struct {
	command
}

func (c PutCommand) Action(r Reader, w Writer) {
	name, nok := r.Value(0)
	password, pok := r.Value(1)
	meta := r.Arguments()

	if !nok || !pok {
		w.Write("Incorrect number of arguments.")
		w.Write("usage: vault put <name> \"<password>\" --meta=value")
		return
	}

	for name, value := range meta {
		meta[name] = crypt.Keys.Encrypt([]byte(value))
	}

	pem := crypt.Keys.Encrypt([]byte(password))
	c.storage.Put(name, pem, meta)

	w.Write(fmt.Sprintf("Set password for %s", name))
}
