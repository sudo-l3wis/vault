package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sudo-l3wis/vault/crypt"
	"github.com/sudo-l3wis/vault/data"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@£$%^&*()-_=+[]{};:,<.>?"

type NewCommand struct{}

func (nc NewCommand) Action(r Reader, w Writer) {
	name, ok := r.Value(0)
	meta := r.Arguments()
	if !ok {
		w.Write("Incorrect number of arguments")
		w.Write("Usage: vault new <name> --meta=value")
		return
	}

	password := newPassword(32)
	cipher := crypt.Keys.Encrypt([]byte(password))

	for name, value := range meta {
		meta[name] = crypt.Keys.Encrypt([]byte(value))
	}

	data.Storage.Put(name, cipher, meta)

	fmt.Printf("Set password for %s\n", name)
	fmt.Println(password)
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
