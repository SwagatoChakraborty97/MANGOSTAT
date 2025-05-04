package cmd

import (
	"runtime"
)

// returns number of CPU cores present
func cpuCores () int{
	return runtime.NumCPU()
}