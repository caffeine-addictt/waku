package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDefaultCacheDir() (string, error) {
	dirPath := os.Getenv("XDG_CACHE_HOME")

	if dirPath == "" {
		os.Getenv("LOCALAPPDATA")
	}

	if dirPath == "" {
		dirPath = filepath.Join(os.Getenv("HOME"), ".cache")
	}

	if dirPath == "" {
		dirPath = "/tmp"
	}

	dirPath = filepath.Join(dirPath, "template")
	ok, err := IsDir(dirPath)
	if err != nil {
		return "", err
	}

	// Create the directory if it doesn't exist
	if !ok {
		err := os.Mkdir(dirPath, 0o600) // rw-rw---
		if err != nil {
			return "", fmt.Errorf("failed to create directory %s", err)
		}
	}

	return dirPath, nil
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

	mode := fileinfo.Mode()
	if mode&0o400 == 0 {
		return false, fmt.Errorf("directory %v is not readable", path)
	}
	if mode&0o200 == 0 {
		return false, fmt.Errorf("directory %v is not writable", path)
	}

	return true, nil
}
