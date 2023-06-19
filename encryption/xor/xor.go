package xor

import (
	"crypto/rand"
	"io"
)

func NEW_XOR_KEY() ([]byte, error) {
	pre := []byte{82, 68, 80, 65}
	xor_key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, xor_key); err != nil {
		return nil, err
	}
	return append(pre, xor_key...), nil
}

func EncryptDecrypt(intput, key []byte) []byte {
	out := []byte{}
	kl := len(key)
	for i, v := range intput {
		out = append(out, v^key[i%kl])
	}
	return out
}
