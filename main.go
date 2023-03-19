package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// Get total memory in bytes
	var sysinfo syscall.Sysinfo
	if err := syscall.Sysinfo(&sysinfo); err != nil {
		fmt.Println("Failed to get system info:", err)
		os.Exit(1)
	}
	total := sysinfo.Totalram * uint64(sysinfo.Unit)
	// Calculate the amount of memory to use (4% of total)
	use := total / 4
	// Allocate memory
	if memory := make([]byte, use); memory == nil {
		fmt.Println("Failed to allocate more memory!")
	} else {
		fmt.Println("Done!")
		for {
			time.Sleep(1000)
		}
	}
}
