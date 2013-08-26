package main

import (
	"./archive/zip"
	"fmt"
	"github.com/Bitnick2002/goa/log"
)

func check(err error) {
	if err != nil {
		log.Error(fmt.Sprint(err))
	}
}

func main() {
	err := zip.Unzip("aquincum.zip")
	check(err)
}
