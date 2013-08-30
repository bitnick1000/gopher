package aes

import (
	"encoding/hex"
	. "github.com/Bitnick2002/goa/jasmine"
	"testing"
)

func Test_(t *testing.T) {
	Describe("aes ", func() {
		It("should encrypt  ", func() {
			key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
			iv, _ := hex.DecodeString("101112131415161718191a1b1c1d1e1f")
			message := "message"
			result := make([]byte, len(message))
			err := CFBEncrypt(result, []byte(message), key, iv)
			Expect(err == nil).ToBeTruthy()
		})
	})
}
