package main

import (
	"bufio"
	"fmt"
	"os"
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

func isPossible(design string, patterns []string, memo map[string]bool) bool {
	if result, exists := memo[design]; exists {
		return result
	}
	if design == "" {
		return true
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			if isPossible(design[len(pattern):], patterns, memo) {
				memo[design] = true
				return true
			}
		}
	}
	memo[design] = false
	return false
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	patterns := strings.Split(file[0], ", ")
	memo := make(map[string]bool)
	result := 0
	for i := 2; i < len(file); i++ {
		if isPossible(file[i], patterns, memo) {
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
