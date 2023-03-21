package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
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
	num, _ := strconv.Atoi(os.Getenv("MULTI"))
	if num == 0 {
		num = 5
	}
	if memory := make([]byte, use*uint64(num)); memory == nil {
		fmt.Println("Failed to allocate memory!")
	} else {
		fmt.Println("Done!")
		/* 	"os/signal" sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		// Loop until SIGINT is received
		<-sigint */
		time.Sleep(time.Second * 10)
		for {
			p := int64(0)
			for i := int64(1); i <= 10000000; i++ {
				p = i
			}
			fmt.Printf("Cpu waste data: %d\n", p)
			time.Sleep(time.Minute * 10)
		}
		
	}
}
