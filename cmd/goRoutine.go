package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var goRoutineCmd = &cobra.Command{
	Use:   "goroutine",       // command name
	Short: "Collects numGoRoutine", // short description
	Long:  "Collects numGoRoutine Running",  // long description
	Run: func(cmd *cobra.Command, args []string) { // specifes what the command does
		numGoRoutine := countRunningGoRoutines()
		fmt.Println(numGoRoutine,"Running GoRoutine")
	},
}

func init() {
	// fmt.Println("THIS IS FROM NUM numGoRoutine INIT")
	rootCmd.AddCommand(goRoutineCmd)
}

