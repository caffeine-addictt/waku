// Waku's template related functions
package template

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/caffeine-addictt/waku/internal/log"
)

// Filenames is a list of all possible template filenames
// excluding extension
var Filenames [12]string = [12]string{
	"waku.yml",
	"template.yml",
	"waku.yaml",
	"template.yaml",
	"waku.json",
	"template.json",
	".waku.yml",
	".template.yml",
	".waku.yaml",
	".template.yaml",
	".waku.json",
	".template.json",
}

func GetWakuConfig(p string) (string, *os.File, error) {
	// if p is a file, treat it as a config file
	if p != "." {
		f, err := openConfigFile(p)
		if err != nil {
			return "", nil, err
		}
		if f != nil {
			return p, f, nil
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(len(Filenames))

	errChan := make(chan error, 1)
	pathChan := make(chan string, 1)
	fileChan := make(chan *os.File, 1)

	for _, name := range Filenames {
		go func(p, name string) {
			defer wg.Done()
			path := filepath.Join(p, name)

			select {
			case <-ctx.Done():
				log.Debugf("skipping reading config file %s: %s\n", path, ctx.Err())
				return
			default:
			}

			f, err := openConfigFile(path)
			if err != nil {
				once.Do(func() {
					errChan <- fmt.Errorf("failed to open config file %s: %s", path, err)
					cancel()
				})
				return
			}
			if f == nil {
				return
			}

			select {
			case <-ctx.Done():
				defer f.Close()
				log.Debugf("closing config file read at %s as context is done\n", path)
			default:
				once.Do(func() {
					log.Debugf("successfully opened config file %s\n", path)
					pathChan <- path
					fileChan <- f
					cancel()
				})
			}
		}(p, name)
	}

	// cleanup
	go func() {
		wg.Wait()
		log.Debugln("closed config file channels")
		close(errChan)
		close(pathChan)
		close(fileChan)
	}()

	select {
	case err := <-errChan:
		// might also trigger due to all the chan closing
		// but no file was found
		if err == nil {
			return "", nil, fmt.Errorf("no config file found, check that you have a file named: %v", Filenames)
		}

		log.Debugf("failed to open config file: %s\n", err)
		return "", nil, err

	case path := <-pathChan:
		log.Debugf("resolved config file at %s\n", path)
		return path, <-fileChan, nil

	case <-ctx.Done():
		log.Debugln("config file read timed out")
		return "", nil, ctx.Err()
	}
}

// *os.File can be nil for recoverable errors
func openConfigFile(path string) (*os.File, error) {
	path = filepath.Clean(path)

	fi, err := os.Stat(path)
	if (fi != nil && fi.IsDir()) || (err != nil && errors.Is(err, os.ErrNotExist)) {
		log.Debugf("config file %s does not exist or is a directory\n", path)
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	log.Debugf("trying to open config file %s\n", path)
	return os.Open(path)
}
