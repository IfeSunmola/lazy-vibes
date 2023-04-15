package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
	"math"
	"os"
	"time"
)

const (
	defaultInterval     = 10 * time.Second
	defaultUsagePercent = 999
	defaultTime         = 3240000
	defaultName         = "-------"
	defaultMaxCPUUsage  = 80.0
	defaultMaxMemUsage  = 10.0
	defaultMaxTime      = 10 * time.Minute
)

func currentTime() string {
	return time.Now().Format("January 2, 2006 at 3:04:05 PM")
}

func getFormattedOwner(owner string) string {
	const maxLen = 7
	if len(owner) < maxLen {
		return owner
	}
	return owner[:maxLen] + "+"
}

func getTimeStr(seconds float64) string {
	raw := int32(seconds)
	h := raw / 3600
	m := (raw % 3600) / 60
	s := raw % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func checkProcesses() error {
	fmt.Printf("Process tracker output (%s)\n", currentTime())
	printer := table.NewWriter()
	printer.SetOutputMirror(os.Stdout)
	printer.AppendHeader(table.Row{"PID", "USER", "NAME", "CPU %", "MEM %", "TIME"})
	printer.SetAutoIndex(true)

	pids, err := process.Pids()
	if err != nil {
		return fmt.Errorf("couldn't get all pids: %w", err)
	}

	for _, pid := range pids {
		proc, err := process.NewProcess(pid)
		if err != nil {
			// couldn't create process instance. Maybe because it's exited or
			//permission issues, or something else
			// just skip it
			continue
		}

		cpuUsage, err := proc.CPUPercent()
		if err != nil {
			cpuUsage = defaultUsagePercent
		}
		cpuUsage = math.Ceil(cpuUsage*100) / 100

		memUsage, err := proc.MemoryPercent()
		if err != nil {
			memUsage = defaultUsagePercent
		}
		memUsage = float32(math.Ceil(float64(memUsage*100)) / 100)

		times, err := proc.Times()
		if err != nil {
			times = &cpu.TimesStat{System: defaultTime}
		}

		totalTime := times.System + times.User
		timeStr := getTimeStr(totalTime)

		if cpuUsage >= defaultMaxCPUUsage || memUsage >= defaultMaxMemUsage ||
			totalTime >= defaultMaxTime.Seconds() {

			procName, err := proc.Name()
			if err != nil {
				procName = defaultName
			}

			procOwner, err := proc.Username()
			if err != nil {
				procOwner = defaultName
			}
			procOwner = getFormattedOwner(procOwner)

			printer.AppendRow(table.Row{pid, procOwner, procName, cpuUsage, memUsage, timeStr})
		}
	}
	printer.Render()
	return nil
}

func main() {
	for range time.Tick(defaultInterval) { // run loop content every defaultInterval seconds
		if err := checkProcesses(); err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}
