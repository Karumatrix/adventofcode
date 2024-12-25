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

func getSchematics(file []string) (schematics [][]string) {
	index := 0
	schematics = append(schematics, []string{})

	for _, line := range file {
		if line == "" {
			schematics = append(schematics, []string{})
			index++
		} else {
			schematics[index] = append(schematics[index], line)
		}
	}
	return
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	schematics := getSchematics(file)
	var locks [][5]int
	var keys [][5]int

	for _, schema := range schematics {
		var pins [5]int
		for col := 0; col < len(schema[0]); col++ {
			pin := 0
			for row := 0; row < len(schema); row++ {
				if schema[row][col] == '#' {
					pin++
				}
			}
			pins[col] = pin - 1
		}
		if schema[0] == "#####" {
			locks = append(locks, pins)
		} else if schema[len(schema) - 1] == "#####" {
			keys = append(keys, pins)
		}
	}
	result := 0
	for _, lock := range locks {
		CheckKey:
		for _, key := range keys {
			for i := 0; i < 5; i++ {
				if lock[i] + key[i] > 5 {
					continue CheckKey
				}
			}
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
