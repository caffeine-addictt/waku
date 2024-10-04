// For async cleanup in cases of interrupts
package cleanup

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/caffeine-addictt/waku/internal/log"
)

var (
	cleanupStack      []func() error
	cleanupRwLock     sync.RWMutex
	errorCleanupStack []func() error
	errorRwLock       sync.RWMutex
)

// Schedule adds a function to the cleanup stack
//
// Cleanups are ran no matter what, before error cleanup
func Schedule(fn func() error) {
	cleanupRwLock.Lock()
	defer cleanupRwLock.Unlock()
	cleanupStack = append(cleanupStack, fn)
}

// ScheduleError adds a function to the error cleanup stack
//
// Error cleanups are ran ONLY when an error occurs, and
// are ran after all regular cleanups
func ScheduleError(fn func() error) {
	errorRwLock.Lock()
	defer errorRwLock.Unlock()
	errorCleanupStack = append(errorCleanupStack, fn)
}

// Runs all cleanup functions
func Cleanup() {
	cleanupRwLock.Lock()
	defer cleanupRwLock.Unlock()

	log.Debugf("cleaning up %d items...\n", len(cleanupStack))
	for i := len(cleanupStack) - 1; i >= 0; i-- {
		if err := cleanupStack[i](); err != nil {
			log.Errorf("error while cleaning up: %v\n", err)
		}
	}

	cleanupStack = []func() error{}
}

// Runs all error cleanup functions
func CleanupError() {
	errorRwLock.Lock()
	defer errorRwLock.Unlock()

	log.Debugf("cleaning up %d items due to error...\n", len(errorCleanupStack))
	for i := len(errorCleanupStack) - 1; i >= 0; i-- {
		if err := errorCleanupStack[i](); err != nil {
			log.Errorf("error while cleaning up: %v\n", err)
		}
	}

	errorCleanupStack = []func() error{}
}

// Watches for interrupts and runs all cleanup functions.
func On() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("%v received, cleaning up...\n", sig)

		Cleanup()
		CleanupError()
	}()
}
