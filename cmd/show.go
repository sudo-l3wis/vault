package cmd

import "fmt"

type ShowCommand struct {}

func (sc ShowCommand) Action(ctx *Context) {
	name, ok := ctx.Reader.Value(0)
	if !ok {
		ctx.Writer.Write("Invalid number of arguments.")
		ctx.Writer.Write("usage: vault show <name>")
		return
	}

	password, meta := ctx.Store.Show(name)
	if password.ID == 0 {
		ctx.Writer.Write(fmt.Sprintf("No password with name %s", name))
		return
	}

	if ctx.Reader.Option("r") {
		ctx.Writer.Write(password.Body)
		return
	}

	body := ctx.Crypt.PemToMsg(password.Body)
	decrypted := ctx.Crypt.Decrypt(body)

	ctx.Writer.Write(fmt.Sprintf("password: %s", decrypted))

	for _, m := range meta {
		b := ctx.Crypt.PemToMsg(m.Value)
		d := ctx.Crypt.Decrypt(b)
		ctx.Writer.Write(fmt.Sprintf("%s: %s", m.Name, d))
	}
}
