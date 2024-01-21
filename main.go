package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
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
	fmt.Println("Total memory:", total)
	use := total / 20
	fmt.Println("Memory to use:", use)
	num, _ := strconv.Atoi(os.Getenv("MULTI"))
	if num == 0 {
		num = 5
	}
	if memory := make([]byte, use*uint64(num)); memory == nil {
		fmt.Println("Failed to allocate memory!")
	}
	if num > 9 && total > 12520169472 {
		if memory1 := make([]byte, use*uint64(num)); memory1 == nil {
			fmt.Println("Failed to allocate memory 1!")
		}
		if memory2 := make([]byte, use*uint64(num)); memory2 == nil {
			fmt.Println("Failed to allocate memory 1!")
		}
	}
	fmt.Println("Done!")
	if os.Getenv("NOBURN") == "1" {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
	} else {
		time.Sleep(time.Hour * 12)
		for {
			done := make(chan int)
			for i := 0; i < runtime.NumCPU()+2; i++ {
				go func() {
					for {
						select {
						case <-done:
							return
						default:
						}
					}
				}()
			}
			time.Sleep(time.Second * 10)
			close(done)
			time.Sleep(time.Hour * 12)
		}
	}
}
