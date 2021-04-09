package vault

import "github.com/sudo-l3wis/vault/commands"

func main() {
	ctx := commands.MakeContext()
	command := commands.MakeCommand()

	app := commands.App{}
	app.Run(ctx, command)
}
