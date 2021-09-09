package integration

import (
	"io/ioutil"
	"os"

	"github.com/lsnow99/conductorr/settings"
)

func CheckPath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	f, err := ioutil.TempFile(path, "access_check")
	f.Close()
	os.Remove(f.Name())
	return err
}

func TempDir() (string, error) {
	var tmpDir string
	if settings.TempDir != "" {
		tmpDir = settings.TempDir
	} else {
		tmpDir = os.TempDir()
	}
	err := CheckPath(tmpDir)
	if err != nil {
		err = os.MkdirAll(tmpDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return tmpDir, nil
}
