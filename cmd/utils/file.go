package utils

import (
	"fmt"
	"os"
)

func IsDir(path string) (bool, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if !fileinfo.IsDir() {
		return false, fmt.Errorf("%v is not a directory", path)
	}

	mode := fileinfo.Mode()
	if mode&0o400 == 0 {
		return false, fmt.Errorf("directory %v is not readable", path)
	}
	if mode&0o200 == 0 {
		return false, fmt.Errorf("directory %v is not writable", path)
	}

	return true, nil
}
