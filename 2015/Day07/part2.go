package main

import (
	"bufio"
	"fmt"
	"os"
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

var instructions = make(map[string]string)
var cache = make(map[string]uint16)

func evaluate(wire string) uint16 {
	if val, exists := cache[wire]; exists {
		return val
	}
	if val, err := strconv.Atoi(wire); err == nil {
		return uint16(val)
	}
	instruction := instructions[wire]
	parts := strings.Split(instruction, " ")
	var result uint16

	switch len(parts) {
	case 1:
		result = evaluate(parts[0])
	case 2:
		result = ^evaluate(parts[1])
	case 3:
		left := evaluate(parts[0])
		op := parts[1]
		right := evaluate(parts[2])

		switch op {
		case "AND":
			result = left & right
		case "OR":
			result = left | right
		case "LSHIFT":
			result = left << right
		case "RSHIFT":
			result = left >> right
		}
	}
	cache[wire] = result
	return result
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)

	for _, line := range file {
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0]
	}
	signal := evaluate("a")
	cache = make(map[string]uint16)
	instructions = make(map[string]string)
	for _, line := range file {
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0]
	}
	cache["b"] = signal
	fmt.Println(evaluate("a"))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
