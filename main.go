package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/mem"
)

func ramToConsume() []byte {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error retrieving virtual memory info:", err)
		return nil
	}
	takeEnv := os.Getenv("TAKE")
	take := 15
	if takeEnv != "" {
		take, err = strconv.Atoi(takeEnv)
		if err != nil {
			fmt.Println("Error parsing TAKE environment variable:", err)
			take = 15
		}
	}
	memoryToAllocate := float64(v.Total)*float64(take)/100.0 - float64(v.Used)
	if memoryToAllocate > 0 {
		allocatedMem := bytes.Repeat([]byte{0}, int(memoryToAllocate))
		fmt.Printf("Allocated %d bytes to reach target memory usage.\n", len(allocatedMem))
		return allocatedMem
	} else {
		fmt.Println("No need to allocate memory.")
		return nil
	}
}

func main() {
	_ = ramToConsume()
	fmt.Println("Done!")
	for {
		time.Sleep(24 * time.Hour)
		if os.Getenv("NOCPUB") == "" {
			result := 1
			for i := 1; i < 1000000; i++ {
				result *= i
			}
		}
		args := []string{"main.go"}
		if _, err := os.Stat("main.go"); os.IsNotExist(err) {
			args[0] = "main"
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println("Error re-executing the program:", err)
		}
	}
}
