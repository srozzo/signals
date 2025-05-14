package signals

import "sync"

var (
	debugMu sync.RWMutex
	debug   bool
)

// SetDebug enables or disables debug logging globally in the signals package.
// When enabled, detailed information about signal handling will be logged.
func SetDebug(enabled bool) {
	debugMu.Lock()
	defer debugMu.Unlock()
	debug = enabled

	logf("[signals] debug mode enabled: %t", enabled)
}

// isDebug returns whether debug logging is currently enabled.
// Used internally to control verbose output during signal handling.
func isDebug() bool {
	debugMu.RLock()
	defer debugMu.RUnlock()
	return debug
}
