package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var ramCmd = &cobra.Command{
	Use:   "ram",               // command name
	Short: "Collects RAM Data", // short description
	Long:  "Collects RAM Data", // long description
	Run: func(cmd *cobra.Command, args []string) { // specifes what the command does
		// DEBUG LINE
		// fmt.Println("HELLO INSIDE", args)

		// COLLECTING FIRST METRIC - go run . ram -m <use>
		firstMetric, err := cmd.Flags().GetString("metric")
		// ERROR HANDLING
		if err != nil {
			log.Fatal(err)
		}

		// APPENDING OTHER METRICS TO A SLICE CALLED allMetrics - go run . -m <use> <total> <free>
		allMetrics := []string{}
		allMetrics = append(allMetrics, firstMetric)
		allMetrics = append(allMetrics, args...)

		// DEFINE TABLE FOR METRICS
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Metric", "Value"})

		// DEFINING WHAT ACTIONS NEED TO BE MADE FOR WHICH ARGUMENT
		// ACTIONS = {string : anonymous func} // When a string is called it triggers the anonymous func it is mapped to
		actions := map[string]func(){
			"total": func() {
				currTotalRamAlloc := totalAllocatedRam()
				t.AppendRow([]interface{}{"Total Allocated Memory", fmt.Sprintf("%d bytes", currTotalRamAlloc)})
				// fmt.Printf("Total Allocated Memory: %d bytes\n", currTotalRamAlloc)
			},
			"used": func() {
				currRamUse := ramUnderUse()
				t.AppendRow([]interface{}{"Memory Under Use", fmt.Sprintf("%.2f%%", currRamUse)})
				// fmt.Printf("Memory Under Use: %.2f%%\n", currRamUse)
			},
			"free": func() {
				currFreeRam := freeRam()
				fmt.Printf("Free Memory: %d bytes\n", currFreeRam)
			},
			"available": func() {
				currAvailableRam := availableRam()
				fmt.Printf("Available Memory: %d bytes\n", currAvailableRam)
			},
			"used_raw": func() {
				currUsedRam := usedRamBytes()
				fmt.Printf("Used Memory: %d bytes\n", currUsedRam)
			},
			"cached": func() {
				currCachedRam := cachedRam()
				fmt.Printf("Cached Memory: %d bytes\n", currCachedRam)
			},
			"buffered": func() {
				currBufferedRam := bufferedRam()
				fmt.Printf("Buffered Memory: %d bytes\n", currBufferedRam)
			},
			"swap_total": func() {
				currTotalSwap := totalSwap()
				fmt.Printf("Total Swap Memory: %d bytes\n", currTotalSwap)
			},
			"swap_used": func() {
				currUsedSwap := usedSwap()
				fmt.Printf("Used Swap Memory: %d bytes\n", currUsedSwap)
			},
			"swap_free": func() {
				currFreeSwap := freeSwap()
				fmt.Printf("Free Swap Memory: %d bytes\n", currFreeSwap)
			},
			"swap_used_percent": func() {
				currSwapUsedPercent := swapUsedPercent()
				fmt.Printf("Swap Used Percentage: %.2f%%\n", currSwapUsedPercent)
			},
			"active": func() {
				currActiveRam := activeRam()
				fmt.Printf("Active Memory: %d bytes\n", currActiveRam)
			},
			"inactive": func() {
				currInactiveRam := inactiveRam()
				fmt.Printf("Inactive Memory: %d bytes\n", currInactiveRam)
			},
			"slab": func() {
				currSlabMemory := slabMemory()
				fmt.Printf("Slab Memory: %d bytes\n", currSlabMemory)
			},
			"wired": func() {
				currWiredMemory := wiredMemory()
				fmt.Printf("Wired Memory: %d bytes\n", currWiredMemory)
			},
			"dirty": func() {
				currDirtyMemory := dirtyMemory()
				fmt.Printf("Dirty Memory: %d bytes\n", currDirtyMemory)
			},
			"high_total": func() {
				currHighMemoryTotal := highMemoryTotal()
				fmt.Printf("High Memory Total: %d bytes\n", currHighMemoryTotal)
			},
			"high_free": func() {
				currHighMemoryFree := highMemoryFree()
				fmt.Printf("High Memory Free: %d bytes\n", currHighMemoryFree)
			},
			"total_percent": func() {
				currTotalRamUsagePercent := totalRamUsagePercent()
				fmt.Printf("Total RAM Usage Percentage: %.2f%%\n", currTotalRamUsagePercent)
			},
			"wired_memory": func() {
				currWiredMemoryUsage := wiredMemoryUsage()
				fmt.Printf("Wired Memory Usage: %d bytes\n", currWiredMemoryUsage)
			},
		}

		// allMetrics contains all the metrics passed
		for _, metricValue := range allMetrics {
			// fmt.Println(index, metricValue)
			switch metricValue {
			case "all":
				actions["total"]()
				actions["used"]()
				actions["free"]()
				actions["available"]()
				actions["cached"]()
				actions["buffered"]()
				actions["swap_total"]()
				actions["swap_used"]()
				actions["swap_free"]()
				actions["swap_used_percent"]()
				actions["active"]()
				actions["inactive"]()
				actions["slab"]()
				actions["wired"]()
				actions["dirty"]()
				actions["high_total"]()
				actions["high_free"]()
				actions["total_percent"]()
				actions["wired_memory"]()
			case "total", "used", "free", "available", "cached", "buffered",
				"swap_total", "swap_used", "swap_free", "swap_used_percent",
				"active", "inactive", "slab", "wired", "dirty",
				"high_total", "high_free", "total_percent", "total_swap_percent", "wired_memory":
				if action, ok := actions[metricValue]; ok {
					action()
				}
			default:
				fmt.Printf("%q is an invalid argument\n", metricValue)
			}
		}

		t.AppendRow([]interface{}{"Time Now", time.Now()})
		
		// fmt.Println("Time Now:", time.Now())
		// t.AppendRow()
		// PRINT THE TABLE
		t.Render()
	},
}

func init() {
	// fmt.Println("THIS IS FROM RAM INIT")
	ramCmd.Flags().StringP("metric", "m", "all", "Greeting message")
	// helloCmd.MarkFlagRequired("greeting") // Marks flag which is required
	rootCmd.AddCommand(ramCmd)
}
