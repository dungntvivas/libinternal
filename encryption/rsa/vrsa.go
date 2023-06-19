package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type Vras struct {
	key *rsa.PrivateKey
}

func VRSA_NEW() (*Vras, error) {
	p := &Vras{}
	var err error
	p.key, err = rsa.GenerateKey(rand.Reader, 1024)
	return p, err
}

func RSA_OAEP_Encrypt(secretMessage []byte, pKey *rsa.PublicKey) ([]byte, error) {

	label := []byte("RDPA-VIVAS")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, pKey, secretMessage, label)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
func RSA_PKCS1_Encrypt(secretMessage []byte, pKey *rsa.PublicKey) ([]byte, error) {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptPKCS1v15(rng, pKey, secretMessage)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func (p *Vras) RSA_OAEP_Decrypt(cipherText []byte) ([]byte, error) {

	label := []byte("RDPA-VIVAS")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, p.key, cipherText, label)
	return plaintext, err
}

func (p *Vras) RSA_PKCS1_Decrypt(cipherText []byte) ([]byte, error) {
	rng := rand.Reader
	plaintext, err := rsa.DecryptPKCS1v15(rng, p.key, cipherText)
	return plaintext, err
}
func (p *Vras) PrivateKey() *rsa.PrivateKey {
	return p.key
}
func (p *Vras) GetPrivateKey() []byte {
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(p.key),
		},
	)
	return pubPEM
}
func (p *Vras) PublicKey() *rsa.PublicKey {
	return &p.key.PublicKey
}
func (p *Vras) GetPublicKey() []byte {
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(p.PublicKey()),
		},
	)
	return pubPEM
}
func (p *Vras) BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	return key
}
func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PublicKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}
