package commands

type Command interface {
	Action(r Reader, w Writer)
}
