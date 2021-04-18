package io

import "fmt"

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(value string) {
	fmt.Println(value)
}
