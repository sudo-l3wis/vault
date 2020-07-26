package cmd

import (
	"../crypt"
	"../data"
)

type Context struct {
	Reader Reader
	Writer Writer
	Store data.Store
	Crypt crypt.Crypt
}

func MakeContext() *Context {
	ctx := Context{
		Reader: ArgumentReader{},
		Writer: ConsoleWriter{},
		Store:  data.Store{},
		Crypt:  crypt.Crypt{},
	}
	ctx.Crypt.Innit()
	ctx.Store.Load()

	return &ctx
}
