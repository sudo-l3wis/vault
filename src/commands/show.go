package commands

import (
	"fmt"

	"github.com/sudo-l3wis/vault/crypt"
)

type ShowCommand struct {}

func (sc ShowCommand) Action(r Reader, w Writer) {
	name, ok := r.Value(0)
	if !ok {
		w.Write("Invalid number of arguments.")
		w.Write("usage: vault show <name>")
		return
	}

	password, meta := data.Storage.Show(name)
	if password.ID == 0 {
		w.Write(fmt.Sprintf("No password with name %s", name))
		return
	}

	if r.Option("r") {
		w.Write(password.Body)
		return
	}

	body := crypt.Keys.PemToMsg(password.Body)
	decrypted := crypt.Keys.Decrypt(body)

	w.Write(fmt.Sprintf("password: %s", decrypted))

	for _, m := range meta {
		b := crypt.Keys.PemToMsg(m.Value)
		d := crypt.Keys.Decrypt(b)
		w.Write(fmt.Sprintf("%s: %s", m.Name, d))
	}
}
