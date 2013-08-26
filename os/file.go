package os

import (
	"os"
)

func OpenFile(name string, flag int, perm FileMode) (file *os.File, err error) {
	err = os.MkdirAll(filepath.Dir(name), perm)

	file, err = os.OpenFile(name, flag, perm)

	return file, err
}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
