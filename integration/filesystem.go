package integration

import (
	"io/ioutil"
	"os"
)

func CheckPath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	f, err := ioutil.TempFile(path, "access_check")
	f.Close()
	return err
}
