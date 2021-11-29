package commands

import (
	"github.com/sudo-l3wis/vault/ciphers"
	"github.com/sudo-l3wis/vault/storage"
)

type command struct {
	cipher  Cipher
	storage Storage
}

func (c *command) SetCipher(cipher Cipher) {
	c.cipher = cipher
}

func (c *command) SetStorage(storage Storage) {
	c.storage = storage
}

func New(name string) Command {
	var cmd Command
	switch name {
	case "show":
		cmd = ShowCommand{}
	case "put":
		cmd = PutCommand{}
	case "new":
		cmd = NewCommand{}
	case "drop":
		cmd = DropCommand{}
	case "ls":
		cmd = LsCommand{}
	case "key":
		cmd = KeyCommand{}
	case "push":
		cmd = PushCommand{}
	default:
		cmd = UsageCommand{}
	}

	cmd.SetCipher(ciphers.X509{})
	cmd.SetStorage(storage.Sqlite{})

	return cmd
}
