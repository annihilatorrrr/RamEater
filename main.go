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
	fmt.Printf("%T", sysinfo.Unit)
	total := sysinfo.Totalram * sysinfo.Unit
	// Calculate the amount of memory to use (10% of total)
	use := total / 10
	// Allocate memory
	mem, err := syscall.Mmap(-1, 0, use, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE)
	defer syscall.Munmap(mem)
	if err != nil {
		fmt.Println("Failed to allocate more memory! " + err.Error())
	} else {
		// Lock the memory into RAM
		if err = syscall.Mlock(mem); err != nil {
			panic(err)
		}
		fmt.Println("Done!")
		for {
			time.Sleep(1000)
		}
	}
}
