package cmd

import "fmt"

type Writer interface {
	Write(value string)
}

type ConsoleWriter struct {}

func (c ConsoleWriter) Write(value string) {
	fmt.Println(value)
}
