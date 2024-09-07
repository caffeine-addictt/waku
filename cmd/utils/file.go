package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	FilePerms = 0o644 // -rw-r--r--
	DirPerms  = 0o755 // drwxr-xr-x
)

func WalkDirRecursive(root string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			paths = append(paths, filepath.ToSlash(relPath))
		}

		return nil
	})

	return paths, err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

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

func IsFile(path string) (bool, error) {
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

	if err = CheckRW(fileinfo.Mode()); err != nil {
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

	err = CheckRWX(fileinfo.Mode())
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
