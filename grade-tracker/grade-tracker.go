package main

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strconv"
	"strings"
)

func usage(msg string) {
	fmt.Println(msg)
	fmt.Println("Usage: ")
	fmt.Println("  ", os.Args[0], " <file_name>")
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func twoDP(num float64) string {
	//return strconv.FormatFloat(num, 'f', 2, 64)
	return fmt.Sprintf("%.2f", num)
}

func main() {
	if len(os.Args) != 2 {
		usage("Invalid number of arguments")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	checkError(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	printer := table.NewWriter()
	printer.SetOutputMirror(os.Stdout)
	printer.SetAutoIndex(true)
	printer.SetStyle(table.StyleLight)
	printer.Style().Options.SeparateRows = true
	printer.Style().Title.Align = text.AlignCenter
	printer.SetTitle("Report for " + os.Args[1])
	printer.AppendHeader(table.Row{"Name", "Achieved", "Possible", "% of final grade", "Actual Achieved"})
	var (
		totalAchieved, totalPossible, totalFinal, totalActualAchieved float64
		lineNum                                                       int8
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Split(line, ",")

		if len(fields) != 4 {
			fmt.Println("Invalid # of values on line", lineNum)
			os.Exit(1)
		}

		name := strings.TrimSpace(fields[0])

		achieved, err := strconv.ParseFloat(strings.TrimSpace(fields[1]), 32)
		checkError(err)

		possible, err := strconv.ParseFloat(strings.TrimSpace(fields[2]), 32)
		checkError(err)

		final, err := strconv.ParseFloat(strings.TrimSpace(fields[3]), 32)
		checkError(err)

		var actualAchieved = (achieved / possible) * final

		if achieved > possible {
			fmt.Println("Achieved is greater than possible for", name, "on line", lineNum)
			os.Exit(1)
		}

		printer.AppendRow(table.Row{name, twoDP(achieved), twoDP(possible), twoDP(final), twoDP(actualAchieved)})

		totalAchieved += achieved
		totalPossible += possible
		totalFinal += final
		totalActualAchieved += actualAchieved
	}

	printer.AppendFooter(table.Row{"Total", twoDP(totalAchieved), twoDP(totalPossible), twoDP(totalFinal), twoDP(totalActualAchieved)})
	printer.Render()

	if totalFinal > 100 {
		fmt.Println("WARNING: Total % final grade is greater than 100%")
		fmt.Println("----------------------------------------------------")
	}

	fmt.Println("You currently have", twoDP(totalActualAchieved), "out of", twoDP(totalFinal))
}
