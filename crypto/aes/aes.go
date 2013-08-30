package aes

import (
	"crypto/aes"
)

func CFBEncrypt(dst, src, key, iv []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(dst, src)
	return nil
}
