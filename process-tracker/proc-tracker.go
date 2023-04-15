package main

import (
	"bytes"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	defaultInterval     = 10 * time.Second
	defaultUsagePercent = 999
	defaultTime         = 3240000
	defaultName         = "-------"
	defaultMaxCPUUsage  = 2.9
	defaultMaxMemUsage  = 50.0
	defaultMaxTime      = 6 * time.Minute
)

const (
	// column names
	cPid  = "PID"
	cUser = "USER"
	cName = "NAME"
	cCPU  = "CPU %"
	cMem  = "MEM %"
	cTime = "TIME"
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

func checkProcesses() (string, error) {
	printer := table.NewWriter()
	output := bytes.NewBufferString("")
	printer.SetOutputMirror(output) // save in variable, not stdout
	printer.AppendHeader(table.Row{cPid, cUser, cName, cCPU, cMem, cTime})
	printer.SetAutoIndex(true)
	printer.SortBy([]table.SortBy{{Name: cCPU, Mode: table.Dsc}})

	pids, err := process.Pids()
	if err != nil {
		return "", fmt.Errorf("couldn't get all pids: %w", err)
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
			output.Reset()   // reset or else it will append to previous output
			printer.Render() // save in "output" variable
		}
	}
	// at this point, if there are any "bad" process, output will contain the rendered table
	// if not, it will be empty
	return output.String(), nil
}

func main() {
	fmt.Println("Process Tracker started")
	for range time.Tick(defaultInterval) { // run loop content every defaultInterval seconds
		output, err := checkProcesses()
		if err != nil {
			fmt.Printf("Error while checking proccesses : %v", err)
			fmt.Println("Program will exit now")
			os.Exit(1)
		}

		output = strings.TrimSpace(output)

		if output != "" { // there are some "bad" processes, send notification
			// send GUI notification
			_, err := exec.Command("notify-send", "-u", "critical", "Some processes are running wild!").Output()
			if err != nil {
				fmt.Println("Couldn't send notification:", err)
				fmt.Println("Do you have notify-send installed?")
			}

			// send broadcast message across logged in shells
			title := fmt.Sprintf("Process Tracker output at %s\n", currentTime())
			_, err = exec.Command("wall", title, output).Output()
			if err != nil {
				fmt.Println("Couldn't send broadcast message:", err)
				fmt.Println("Do you have wall installed?")
			}
		}
	}
}
