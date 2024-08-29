package utils

import (
	"fmt"
	"io/fs"
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

	err = CheckRW(fileinfo.Mode())
	if err != nil {
		return false, fmt.Errorf("file %v is %v", fileinfo.Name(), err)
	}

	return true, nil
}

func IsExecutableFile(path string) (bool, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if fileinfo.IsDir() {
		return false, fmt.Errorf("%v is not a file", path)
	}

	err = CheckRW(fileinfo.Mode())
	if err != nil {
		return false, fmt.Errorf("file %v is %v", fileinfo.Name(), err)
	}

	return true, nil
}

// Check for rw------- perms
func CheckRW(mode fs.FileMode) error {
	if mode&0o400 == 0 {
		return fmt.Errorf("not readable")
	}
	if mode&0o200 == 0 {
		return fmt.Errorf("not writable")
	}
	return nil
}

// Check for rwx------ perms
func CheckRWX(mode fs.FileMode) error {
	if err := CheckRW(mode); err != nil {
		return nil
	}

	if mode&0o100 == 0 {
		return fmt.Errorf("not executable")
	}

	return nil
}
