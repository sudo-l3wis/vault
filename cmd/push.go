package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type PushCommand struct {}

type Config struct {
	Ip string `json:"ip"`
	Key string `json:"key"`
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func loadSettings() Config {
	path := "/var/lib/vault/settings.json"
	config := Config{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter private ssh key path: ")
		key, _ := reader.ReadString('\n')
		config.Key = strings.TrimSuffix(key, "\n")

		fmt.Print("Enter backup ip including port: ")
		ip, _ := reader.ReadString('\n')
		config.Ip = strings.TrimSuffix(ip, "\n")

		data, _ := json.Marshal(config)

		ioutil.WriteFile(path, data, 0644)
	} else {
		data, _ := ioutil.ReadFile(path)
		json.Unmarshal(data, &config)
	}

	return config
}

func (pc PushCommand) Action(ctx *Context) {
	settings := loadSettings()

	config := &ssh.ClientConfig {
		User: "root",
		Auth: []ssh.AuthMethod{
			publicKey(settings.Key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", settings.Ip, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	dt := time.Now()
	name := dt.Format("02-01-2006-15-04")
	path := "/srv/pwd/" + name + ".db"

	err = scp.CopyPath("/var/lib/vault/vault.db", path, session)
	if err != nil {
		panic("Failed to push password backup: " + err.Error())
	}

	ctx.Writer.Write("Pushed password backup: " + path)
}
