package cmd

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@£$%^&*()-_=+[]{};:,<.>?"

type NewCommand struct {}

func (nc NewCommand) Action(ctx *Context) {
	name, ok := ctx.Reader.Value(0)
	meta := ctx.Reader.Arguments()
	if !ok {
		ctx.Writer.Write("Incorrect number of arguments")
		ctx.Writer.Write("Usage: vault new <name> --meta=value")
		return
	}

	password := newPassword(32)
	cipher := ctx.Crypt.Encrypt([]byte(password))

	for name, value := range meta {
		meta[name] = ctx.Crypt.Encrypt([]byte(value))
	}

	ctx.Store.Put(name, cipher, meta)

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
