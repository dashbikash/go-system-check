package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
)

func swapChecker() {
	for {
		v, err := mem.SwapMemory()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Total Swap Memory: %v mb\n", v.Total/1024/1024)
			if v.Total > 1 {
				//panic("Swap memory is enabled, which is not allowed in this environment.")
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func main() {
	go swapChecker()

	select {}
}
