package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	processes, err := process.Processes()

	if err != nil {
		log.Fatalf("Error fetching processes: %v", err)
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			log.Printf("Error fetching process name for PID %d: %v", p.Pid, err)
			continue
		}
		memInfo, err := p.MemoryInfo()
		if err != nil {
			log.Printf("Error fetching memory info for PID %d: %v", p.Pid, err)
			continue
		}
		if name == "swapcheck" {
			fmt.Printf("Process ID: %d, Name: %s, RSS: %d mb, Swap: %d bytes\n", p.Pid, name, memInfo.RSS/1024/1024, memInfo.Swap)
			p.Kill()

		}

	}
}
