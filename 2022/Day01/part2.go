package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func filepathToString(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
		panic(err)
    }
    return string(content)
}

func filepathToStringArray(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	var caloriesList []int
	calories := 0

	for _, line := range file {
		if line == "" {
			caloriesList = append(caloriesList, calories)
			calories = 0
		} else {
			tmpCalories := 0
			fmt.Sscanf(line, "%d", &tmpCalories)
			calories += tmpCalories
		}
	}
	caloriesList = append(caloriesList, calories)
	sort.Slice(caloriesList, func(i, j int) bool {
		return caloriesList[i] > caloriesList[j]
	})
	fmt.Println(caloriesList[0] + caloriesList[1] + caloriesList[2])
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
