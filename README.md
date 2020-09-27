# Vault

A password manager using 4096-bit RSA-OAEP.

#### Store passwords
Store an existing password or create a new one. You may add meta data to the
password e.g. `email=foo@bar.com "dob=July 1 1990"`.
```
vault put <name> "<password>" --foo=bar --fizz=buzz
vault new <name> --foo=bar --fizz=buzz
```

#### Show passwords
```
vault show <name>
vault show <name> -r
vault ls
```

#### Delete passwords
```
vault drop <name>
```

#### List passwords
```
vault ls
```

#### Encryption Keys:
Keys are stored in `/var/lib/vault/keys` and can be displayed with the following
commands.
```
vault key public
vault key private 
```

#### Backup password database to remote server:
This will prompt for a an ip and ssh key path to perform scp and will save the config in `/var/lib/vault/settings.json`.
Databases are copied to `/srv/pwd/`.
```
vault push
```

## Install:
cd to the repository directory. Vault requires permissions to access data
in /var/lib.
```
go build
cp Vault /usr/local/bin/vault
chmod +X /usr/local/bin/vault
```

## About:
Data is stored in `/var/lib/vault`.
