package cmd

type UsageCommand struct {}

func (uc UsageCommand) Action(ctx *Context) {
	ctx.Writer.Write("aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQo=")
}
