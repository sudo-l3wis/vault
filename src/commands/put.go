package commands

import "fmt"

type PutCommand struct{}

func (sc PutCommand) Action(ctx *Context) {
	name, nok := ctx.Reader.Value(0)
	password, pok := ctx.Reader.Value(1)
	meta := ctx.Reader.Arguments()

	if !nok || !pok {
		ctx.Writer.Write("Incorrect number of arguments.")
		ctx.Writer.Write("usage: vault put <name> \"<password>\" --meta=value")
		return
	}

	for name, value := range meta {
		meta[name] = ctx.Crypt.Encrypt([]byte(value))
	}

	pem := ctx.Crypt.Encrypt([]byte(password))
	ctx.Store.Put(name, pem, meta)

	ctx.Writer.Write(fmt.Sprintf("Set password for %s", name))
}
