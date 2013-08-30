package main

import (
	// "./archive/zip"
	"crypto/md5"
	// "encoding/hex"
	"fmt"
	"github.com/Bitnick2002/goa/log"
	"io"
)

func check(err error) {
	if err != nil {
		log.Error(fmt.Sprint(err))
	}
}

func main() {
	fun()
	fun2()
	hash := md5.New()
	io.WriteString(hash, "Message")
	fmt.Println(hash.Sum(nil)) //hash.Sum(nil)

}

func fun() {
	hash := md5.New()
	hash.Write([]byte("Message"))
	result := hash.Sum(nil)
	//hash.Write(result)
	//result = hash.Sum(nil)
	fmt.Println(result) //hash.Sum(nil)
}
func fun2() {
	hash := md5.New()
	hash.Write([]byte("Message"))
	result := hash.Sum(nil)
	hash.Reset()
	hash.Write([]byte("Message"))
	result = hash.Sum(nil)
	fmt.Println(result) //hash.Sum(nil)
}
