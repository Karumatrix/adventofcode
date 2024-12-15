package main

import (
	"bufio"
	"fmt"
	"os"
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

func checkPairs(line string) bool {
	for i := 0; i < len(line) - 3; i++ {
		for j := i + 2; j < len(line) - 1; j++ {
			if line[i] == line[j] && line[i + 1] == line[j + 1] {
				return true
			}
		}
	}
	return false
}

func checkPalindrome(line string) bool {
	for i := 0; i < len(line) - 2; i++ {
		if line[i] == line[i + 2] {
			return true
		}
	}
	return false
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for _, line := range file {
		if checkPairs(line) && checkPalindrome(line) {
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
