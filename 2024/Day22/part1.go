package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getNextSecret(secret int) int {
	tmp := secret * 64
	result := (secret ^ tmp) % 16777216
	tmp = result / 32
	result = (result ^ tmp) % 16777216
	tmp = result * 2048
	return (result ^ tmp) % 16777216
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for _, line := range file {
		value, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			value = getNextSecret(value)
		}
		result += value
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
