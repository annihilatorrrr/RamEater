package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

func main() {
	// Get total memory in bytes
	var sysinfo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&sysinfo); err != nil {
		fmt.Println("Failed to get system info:", err.Error())
		os.Exit(1)
	}
	total := sysinfo.Totalram * uint64(sysinfo.Unit)
	// Calculate the amount of memory to use (10% of total)
	use := total / 10
	// Allocate memory
	/* num, _ := strconv.Atoi(os.Getenv("MULTI")) "strconv"
	if num == 0 {
		num = 6
	}
	*uint64(num) */
	memory := make([]byte, use)
	if memory == nil {
		fmt.Println("Failed to allocate memory!")
	} else {
		unsafe.Pointer(&memory)
		fmt.Println("Done!")
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		// Loop until SIGINT is received
		<-sigint
	}
}
