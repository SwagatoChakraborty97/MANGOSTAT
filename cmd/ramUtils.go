package cmd

import (
    "github.com/shirou/gopsutil/v3/mem"
)

// "all" - Return all the stats below 

// "total" - Returns total allocated RAM
func totalAllocatedRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Total
}

// "used" - Returns percentage of RAM used (excluding cache)
func ramUnderUse() float64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.UsedPercent
}

// "free" - Returns amount of free RAM
func freeRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Free
}

// "available" - Returns amount of available RAM
func availableRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Available
}

// "used_raw" - Returns used RAM in bytes
func usedRamBytes() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Used
}

// "cached" - Returns cached memory in bytes
func cachedRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Cached
}

// "buffered" - Returns buffered memory in bytes
func bufferedRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Buffers
}

// "swap_total" - Returns total swap memory in bytes
func totalSwap() uint64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.Total
}

// "swap_used" - Returns used swap memory in bytes
func usedSwap() uint64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.Used
}

// "swap_free" - Returns free swap memory in bytes
func freeSwap() uint64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.Free
}

// "swap_used_percent" - Returns percentage of swap used
func swapUsedPercent() float64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.UsedPercent
}

// "swap_in" - Returns number of swap pages read into memory
/*func swapIn() uint64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.swapIn
}*/

// "swap_out" - Returns number of swap pages written out to disk
/*func swapOut() uint64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.Swapout
}*/

// "swap_cached" - Returns swap cached memory in bytes
/* func swapCachedMemory() uint64 {
	swapStat, _ := mem.SwapMemory()
	return swapStat.Cached
} */

// "active" - Returns amount of active RAM in bytes
func activeRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Active
}

// "inactive" - Returns amount of inactive RAM in bytes
func inactiveRam() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Inactive
}

// "slab" - Returns slab memory used by the kernel in bytes
func slabMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Slab
}

// "wired" - Returns wired memory in bytes
func wiredMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Wired
} 

// "filecache" - Returns file cache memory in bytes
/* func fileCacheMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Filecache
} */

// "anonpages" - Returns anonymous pages memory in bytes
/* func anonPagesMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Anonpages
} */

// "user_time" - Returns user space memory in bytes
/* func userSpaceMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.UserTime
} */

// "system_time" - Returns system space memory in bytes
/* func systemSpaceMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.SystemTime
} */

// "kmem" - Returns kernel memory in bytes
/* func kernelMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Kmem
} */

// "pgpgin" - Returns number of pages read from disk
/* func pgpgIn() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Pgpgin
} */

// "pgpgout" - Returns number of pages written to disk
/* func pgpgOut() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Pgpgout
} */

// "file_backed" - Returns file-backed pages memory in bytes
/* func fileBackedMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Filebacked
} */

// "pagecache" - Returns page cache memory used in bytes
/* func pageCache() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Pagecache
}  */

// "dirty" - Returns dirty memory in bytes
func dirtyMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Dirty
}

// "dirty_ratio" - Returns the ratio of dirty pages in memory
/* func dirtyPagesRatio() float64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.DirtyRatio
} */

// "high_total" - Returns high memory total in bytes
func highMemoryTotal() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.HighTotal
}

// "high_free" - Returns free high memory in bytes
func highMemoryFree() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.HighFree
}

// "hugepages" - Returns memory used by hugepages in bytes
/* func hugePagesMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Hugepages
}*/

// "total_percent" - Returns percentage of total RAM used
func totalRamUsagePercent() float64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.UsedPercent
}

// "swap_pressure" - Returns swap pressure as a percentage
/* func swapPressure() float64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Swapin
}*/

// "committed" - Returns committed memory in bytes
/* func committedMemory() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Committed
}*/

// "wired_memory" - Returns wired memory usage in bytes
func wiredMemoryUsage() uint64 {
	vmStat, _ := mem.VirtualMemory()
	return vmStat.Wired
}
