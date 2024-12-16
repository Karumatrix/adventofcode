package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	result := 0

	for _, line := range file {
		parts := strings.Fields(line)
		value1, _ := strconv.Atoi(parts[0])
		value2, _ := strconv.Atoi(parts[1])
		value3, _ := strconv.Atoi(parts[2])
		sides := []int{value1, value2, value3}
		sort.Slice(sides, func(i, j int) bool {
			return sides[i] < sides[j]
		})
		if sides[0] + sides[1] > sides[2] {
			result++
		}
	}
	fmt.Println(result)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
