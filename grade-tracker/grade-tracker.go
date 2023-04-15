package main

import (
	"bufio"
	"fmt"
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

	fmt.Println("Starting report for", os.Args[1])
	fmt.Println("========================================================")
	fmt.Printf("%-15s %10s %10s %19s", "Name", "|achieved", "|possible", "|% of final grade\n")
	fmt.Println("========================================================")
	var (
		totalAchieved, totalPossible, totalFinal float32
		lineNum                                  int8
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasPrefix(line, "name") {
			fmt.Println("Report name:", strings.SplitN(line, "=", 2)[1])
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

		if achieved > possible {
			fmt.Println("Achieved is greater than possible for", name)
			os.Exit(1)
		}
		//fmt.Printf("%-15s %10s %10s %19s", "Name", "|achieved", "|possible", "|% of final grade\n")
		fmt.Printf("%-15s  |%6.2f    |%7.2f   |%8.2f       |\n", name, achieved, possible, final)
		fmt.Println("--------------------------------------------------------")

		totalAchieved += float32(achieved)
		totalPossible += float32(possible)
		totalFinal += float32(final)
	}

	if totalFinal > 100 {
		fmt.Println("WARNING: Total % final grade is greater than 100%")
		fmt.Println("----------------------------------------------------")
	}

	total := (totalAchieved / totalPossible) * totalFinal
	fmt.Printf("You currently have %.2f out of %.2f\n", total, totalFinal)
}
