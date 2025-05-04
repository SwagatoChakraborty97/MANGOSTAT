// NOTE: In a Cobra app, the main.go file is very bare. It serves one purpose: initializing Cobra.

package main

import (
	"MANGOSTAT/cmd"
)

func main() {
	cmd.Execute()
}

