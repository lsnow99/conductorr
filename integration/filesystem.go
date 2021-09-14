package integration

import (
	"io/ioutil"
	"os"

	"github.com/lsnow99/conductorr/settings"
)

// CheckPath checks a path string and if it is a directory, it attempts
// to create a test file in the directory. If it is a file, it checks
// that the file exists. Returns a non-nil error if an error is encountered
func CheckPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		f, err := ioutil.TempFile(path, "access_check")
		if f != nil {
			f.Close()
			os.Remove(f.Name())
		}
		return err
	}
	return nil
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
