package encryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/square/go-jose/v3"
)

var privateKey *rsa.PrivateKey
var publicKey rsa.PublicKey

func LoadJweKeys() {
	data, err := ioutil.ReadFile("./privateKey.pem")
	if err != nil {
		panic("File reading private error")
	}
	block, _ := pem.Decode(data)
	if block == nil {
		panic("failed to parse PEM block containing the key")
	}

	privRead, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic("Cannot load private rsa key!")
	}
	privateKey = privRead.(*rsa.PrivateKey)
	publicKey = privateKey.PublicKey
}

func CreateToken(payload string) (string, error) {
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: &publicKey}, nil)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	encrypted, err := encrypter.Encrypt([]byte(payload))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return encrypted.FullSerialize(), nil
}

func DecryptToken(token string) (string, error) {
	object, err := jose.ParseEncrypted(token)
	if err != nil {
		return "", err
	}
	decrypted, err := object.Decrypt(privateKey)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
