package main

import (
	"fmt"
	"os"
	"os/signal"
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
	// Allocate memory
	if memory := make([]byte, use*5); memory == nil {
		fmt.Println("Failed to allocate memory!")
	} else {
		fmt.Println("Done!")
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		// Loop until SIGINT is received
		<-sigint
	}
	fmt.Println("Bye!")
}
