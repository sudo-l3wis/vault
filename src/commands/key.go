package commands

type KeyCommand struct {}

func (kc KeyCommand) Action(ctx *Context) {
	name, ok := ctx.Reader.Value(0)
	if !ok {
		ctx.Writer.Write("Incorrect number of arguments.")
		ctx.Writer.Write("Usage: vault key <public|private>")
		return
	}

	private, public := ctx.Crypt.GetKeys()

	if name == "public" {
		ctx.Writer.Write(public)
	} else if name == "private" {
		ctx.Writer.Write(private)
	} else {
		ctx.Writer.Write("Invalid option")
	}
}
