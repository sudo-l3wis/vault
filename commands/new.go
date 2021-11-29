package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sudo-l3wis/vault/types"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@£$^&*()-_=+[]{};:,<.>?"

type NewCommand struct {
	command
}

func (c NewCommand) Action(r Reader, w Writer) {
	name, ok := r.Value(0)
	meta := r.Arguments()
	if !ok {
		w.Write("Incorrect number of arguments")
		w.Write("Usage: vault new <name> --meta=value")
		return
	}

	password := newPassword(32)
	cipher := c.cipher.Encrypt([]byte(password))

	for name, value := range meta {
		meta[name] = c.cipher.Encrypt([]byte(value))
	}

	c.storage.Put(name, cipher, meta)

	w.Write(fmt.Sprintf("Set password for %s\n", name))
	w.Write(password)
}

func newPassword(n int) string {
	chars := []rune(charset)
	rand.Seed(time.Now().UnixNano())

	password := make([]rune, n)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}
