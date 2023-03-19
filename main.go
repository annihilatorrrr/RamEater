package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// Get total memory in bytes
	var sysinfo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&sysinfo); err != nil {
		fmt.Println("Failed to get system info:", err)
		os.Exit(1)
	}
	total := sysinfo.Totalram * uint64(sysinfo.Unit)
	// Calculate the amount of memory to use (10% of total)
	use := total / 10
	fmt.Println(use/(1024 * 1024 * 1024))
	// Allocate memory
	if memory := make([]byte, use); memory == nil {
		fmt.Println("Failed to allocate memory!")
	} else {
		fmt.Println("Done!")
		for {
			time.Sleep(1000)
		}
	}
}
