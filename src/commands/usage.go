package commands

type UsageCommand struct {}

func (uc UsageCommand) Action(ctx *Context) {
	ctx.Writer.Write("Command not found. Try `vlt help`.")
}
