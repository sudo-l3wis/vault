package commands

import (
	"github.com/sudo-l3wis/vault/types"
)

type command struct {
	cipher  types.Cipher
	storage types.Storage
}

func (c *command) SetCipher(cipher types.Cipher) {
	c.cipher = cipher
}

func (c *command) SetStorage(storage types.Storage) {
	c.storage = storage
}

func New(name string) types.Command {
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
