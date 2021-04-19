package ciphers

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
	// bits defines the length of the encryption key.
	bits = 4096

	// path is the absolute path to the encryption keys used
	// to encrypt passwords.
	path = "/var/lib/vault/keys"
)

type X509 struct {
	public  *rsa.PublicKey
	private *rsa.PrivateKey
}

func (x *X509) Encrypt(msg string) string {
	text, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, x.public, []byte(msg), []byte(""))
	return x.MsgToPem(text)
}

func (x *X509) Decrypt(msg string) string {
	text, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, x.private, []byte(msg), []byte(""))
	return string(text)
}

func (x *X509) WriteKeys(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	x.WritePublicKey()
	x.WritePrivateKey()
}

func (x *X509) GetKeys() (string, string) {
	return x.PrivateKeyToPem(), x.PublicKeyToPem()
}

func (x *X509) PublicKeyToPem() string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(x.public),
	}))
}

func (x *X509) PrivateKeyToPem() string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(x.private),
	}))
}

func (x *X509) MsgToPem(msg []byte) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "PASSWORD",
		Bytes: msg,
	}))
}

func (x *X509) PemToMsg(msg string) []byte {
	p, _ := pem.Decode([]byte(msg))
	return p.Bytes
}

func (x *X509) GenKeys() {
	key, _ := rsa.GenerateKey(rand.Reader, bits)
	x.private = key
	x.public = &key.PublicKey
}

func (x *X509) WritePrivateKey() {
	private := x.PrivateKeyToPem()
	err := ioutil.WriteFile(path+"/private.pem", []byte(private), 0644)
	if err != nil {
		panic(err)
	}
}

func (x *X509) WritePublicKey() {
	public := x.PublicKeyToPem()
	err := ioutil.WriteFile(path+"/public.pem", []byte(public), 0644)
	if err != nil {
		panic(err)
	}
}

func (x *X509) ReadPrivateKey() {
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
	x.private = privateKey
}

func (x *X509) ReadPublicKey() {
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
	x.public = pubKey
}
