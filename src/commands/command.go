package commands

type Command interface {
	Action(ctx *Context)
}
