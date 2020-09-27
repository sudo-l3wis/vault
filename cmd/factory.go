package cmd

import "os"

func MakeCommand() Command {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = ""
	}

	switch name {
	case "show":
		return ShowCommand{}
	case "put":
		return PutCommand{}
	case "new":
		return NewCommand{}
	case "drop":
		return DropCommand{}
	case "ls":
		return LsCommand{}
	case "key":
		return KeyCommand{}
	case "push":
		return PushCommand{}
	default:
		return UsageCommand{}
	}
}
