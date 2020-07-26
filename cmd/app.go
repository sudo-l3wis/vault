package cmd

type App struct {}

func (a App) Run(ctx *Context, command Command) {
	command.Action(ctx)
}
