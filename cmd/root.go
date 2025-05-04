package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mangostat",
	Version: "0.1.0",
	Short:   "Monitors VM stats",
	Long: `-- MangoStat Version 1.0 --
Monitors VM stats like CPU and RAM`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("---- 🔍 STARTING SYSTEM MONITORING AT REAL TIME ----")
	},	
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("---- ✅ REVIEW ABOVE METRICS TO ASSESS SYSTEM HEALTH ----")
	},
}

func Execute() {
	// fmt.Println("--ALL GOOD ROOT --")
	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println("--ALL BAD ROOT --")
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	cobra.CheckErr(rootCmd.Execute())
}

/*
func init() {
	fmt.Println("THIS IS FROM ROOT INIT")
	// fmt.Println(userLicense)
}
*/
