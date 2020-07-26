package main

import "./cmd"

func main() {
	ctx := cmd.MakeContext()
	command := cmd.MakeCommand()

	app := cmd.App{}
	app.Run(ctx, command)
}
