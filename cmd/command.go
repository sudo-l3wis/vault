package cmd

type Command interface {
	Action(ctx *Context)
}
