package cmd

import (
	"runtime"
)

// returns the count of running go routines
func countRunningGoRoutines() int {
	return runtime.NumGoroutine()
}


