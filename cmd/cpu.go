package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",       // command name
	Short: "Collects numCPUCore Data", // short description
	Long:  "Collects CPU Data Like number of CPU cores",  // long description
	Run: func(cmd *cobra.Command, args []string) { // specifes what the command does
		numCPUCores := cpuCores()
		fmt.Println(numCPUCores,"CPU Cores Present")
	},
}

func init() {
	// fmt.Println("THIS IS FROM NUM CPU CORES INIT")
	rootCmd.AddCommand(cpuCmd)
}

