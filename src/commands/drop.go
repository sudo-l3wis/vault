package commands

type DropCommand struct {}

func (dc DropCommand) Action(ctx *Context) {
	name, ok := ctx.Reader.Value(0)
	if !ok {
		ctx.Writer.Write("Incorrect number of arguments.")
		ctx.Writer.Write("Usage: vault drop <name>")
		return
	}
	ctx.Store.Drop(name)
}
