package md5

import (
	"crypto/md5"
)

func Iteraion(message []byte, times int) []byte {
	hash := md5.New()
	hash.Write(message)
	ret := hash.Sum(nil)
	for i := 1; i < times; i++ {
		hash.Reset()
		hash.Write(ret)
		ret = hash.Sum(nil)
	}
	return ret
}
