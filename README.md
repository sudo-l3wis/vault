# Vault

A Linux based password manager.

## Install:
cd to the repository directory.
```
go build
cp Vault /usr/local/bin/vault
chmod +X /usr/local/bin/vault
```

## Usage:

#### Store passwords
Store an existing password or create a new one. You may add meta data to the
password e.g. `email=foo@bar.com "dob=July 1 1990"` Note that meta data is not
encrypted.
```
vault put <name> "<password>" foo=bar fizz=buzz
vault new <name> foo=bar fizz=buzz
```

#### Show passwords
```
vault show <name>
vault ls
```

#### Delete passwords
```
vault drop <name>
```

#### Encryption Keys:
Keys are store in `/var/lib/vault/keys` and can be displayed with the following
commands.
```
vault key public
vault key private 
```

#### Remote backup
Configure a remote server. You must have ssh access to
this server.  
```
vault register <name> <ip> <path/on/server>
```
Pull password database from the selected vault. Note that this will override your
existing database.
```
vault pull
```
Backup password database to a remote vault.
```
vault backup
```

## About:
Data is stored in `/var/lib/vault`. Make sure to backup the sqlite database and
encryption keys.
