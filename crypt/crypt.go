package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

const (
	bits = 4096
	path = "/var/lib/vault/keys"
)

type Crypt struct {
	public *rsa.PublicKey
	private *rsa.PrivateKey
}

func (c* Crypt) Innit() {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		c.GenKeys()
		c.WriteKeys()
	}
	c.ReadPrivateKey()
	c.ReadPublicKey()
}

func (c* Crypt) PublicKeyToPem() string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(c.public),
	}))
}

func (c* Crypt) PrivateKeyToPem() string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.private),
	}))
}

func (c* Crypt) MsgToPem(msg []byte) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "PASSWORD",
		Bytes: msg,
	}))
}

func (c* Crypt) PemToMsg(msg string) []byte {
	p, _ := pem.Decode([]byte(msg))
	return p.Bytes
}

func (c* Crypt) GenKeys() {
	key, _ := rsa.GenerateKey(rand.Reader, bits)
	c.private = key
	c.public = &key.PublicKey
}

func (c* Crypt) WritePrivateKey() {
	private := c.PrivateKeyToPem()
	err := ioutil.WriteFile(path + "/private.pem", []byte(private), 0644)
	if err != nil {
		panic(err)
	}
}

func (c* Crypt) WritePublicKey() {
	public := c.PublicKeyToPem()
	err := ioutil.WriteFile(path + "/public.pem", []byte(public), 0644)
	if err != nil {
		panic(err)
	}
}

func (c* Crypt) ReadPrivateKey() {
	file, err := ioutil.ReadFile(path + "/private.pem")
	if err != nil {
		log.Fatal(err)
	}

	decoded, _ := pem.Decode(file)

	var parsedKey interface{}
	parsedKey, err = x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	var privateKey *rsa.PrivateKey
	privateKey, _ = parsedKey.(*rsa.PrivateKey)
	c.private = privateKey
}

func (c* Crypt) ReadPublicKey() {
	file, err := ioutil.ReadFile(path + "/public.pem")
	if err != nil {
		log.Fatal(err)
	}

	decoded, _ := pem.Decode(file)

	var parsedKey interface{}
	parsedKey, err = x509.ParsePKCS1PublicKey(decoded.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	var pubKey *rsa.PublicKey
	pubKey, _ = parsedKey.(*rsa.PublicKey)
	c.public = pubKey
}

func (c* Crypt) WriteKeys() {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	c.WritePublicKey()
	c.WritePrivateKey()
}

func (c* Crypt) Encrypt(msg []byte) string {
	text, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, c.public, msg, []byte(""))
	return c.MsgToPem(text)
}

func (c* Crypt) Decrypt(msg []byte) []byte {
	text, _:= rsa.DecryptOAEP(sha256.New(), rand.Reader, c.private, msg, []byte(""))
	return text
}

func (c* Crypt) GetKeys() (string, string) {
	return c.PrivateKeyToPem(), c.PublicKeyToPem()
}
