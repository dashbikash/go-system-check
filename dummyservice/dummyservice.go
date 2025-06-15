package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fasthttp/router"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/valyala/fasthttp"
)

func CheckMemory() {
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
		fmt.Printf("Process ID: %d, Name: %s, RSS: %d mb, Swap: %d bytes\n", p.Pid, name, memInfo.RSS/1024/1024, memInfo.Swap)
		if (p.Pid == 1 || name == "myapp") && memInfo.Swap > 6*1024*1024 { // Check if RSS is greater than or equal to 13 MB
			//fmt.Printf("Warning: Process %s with PID %d using swap memory usage.\n", name, p.Pid)
			// If the process exceeds the memory limit, kill it and log the information
			// if err := p.Kill(); err != nil {
			// 	log.Printf("Error killing process %s with PID %d: %v", name, p.Pid, err)
			// }
			panic(fmt.Sprintf("Warning: Process %s with PID %d using swap memory usage.\n", name, p.Pid))
		}

	}
}
func CheckMemoryMonitor() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		CheckMemory()
	}
}

var (
	persons []Person
)

func main() {
	go CheckMemoryMonitor()
	r := router.New()
	r.POST("/api", func(ctx *fasthttp.RequestCtx) {
		var person Person
		if err := json.Unmarshal(ctx.PostBody(), &person); err != nil {
			ctx.Error("Invalid JSON", fasthttp.StatusBadRequest)
			return
		}
		persons = append(persons, person)
		fmt.Fprintf(ctx, "Received %d persons\n", len(persons))
	})

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))

}

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
