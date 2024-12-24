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

func findHighestZ(file []string) string {
	result := "z00"

	for _, line := range file {
		parts := strings.Split(line, " ")
		if parts[4][0] == 'z' {
			curr, _ := strconv.Atoi(parts[4][1:])
			ref, _ := strconv.Atoi(result[1:])
			if curr > ref {
				result = parts[4]
			}
		}
	}
	return result
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	index := 0
	wires := make(map[string]uint8)
	for file[index] != "" {
		parts := strings.Split(file[index], " ")
		value, _ := strconv.Atoi(parts[1])
		wires[parts[0][:len(parts[0]) -1]] = uint8(value)
		index++
	}
	index++
	file = file[index:]
	highestZ := findHighestZ(file)
	swaps := make(map[string]bool)
	for _, line := range file {
		parts := strings.Split(line, " ")
		a, op, b, dest := parts[0], parts[1], parts[2], parts[4]
		if dest[0] == 'z' && op != "XOR" && dest != highestZ  {
			swaps[dest] = true
		}
		if op == "XOR" && dest[0] != 'x' && dest[0] != 'y' && dest[0] != 'z' && a[0] != 'x' && a[0] != 'y' && a[0] != 'z' && b[0] != 'x' && b[0] != 'y' && b[0] != 'z' {
			swaps[dest] = true
		}
		if op == "AND" && a != "x00" && b != "x00" {
			for _, subLine := range file {
				subParts := strings.Split(subLine, " ")
				subA, subOp, subB := subParts[0], subParts[1], subParts[2]
				if (dest == subA || dest == subB) && subOp != "OR" {
					swaps[dest] = true
				}
			}
		}
		if op == "XOR" {
			for _, subLine := range file {
				subParts := strings.Split(subLine, " ")
				subA, subOp, subB := subParts[0], subParts[1], subParts[2]
				if (dest == subA || dest == subB) && subOp == "OR" {
					swaps[dest] = true
				}
			}
		}
	}
	keys := make([]string, 0, len(swaps))
	for key := range swaps {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(strings.Join(keys, ","))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
