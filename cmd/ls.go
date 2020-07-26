package cmd

type LsCommand struct {}

func (lc LsCommand) Action(ctx *Context) {
	passwords := ctx.Store.List()
	for _, p := range passwords {
		ctx.Writer.Write(p.Name)
	}
}
