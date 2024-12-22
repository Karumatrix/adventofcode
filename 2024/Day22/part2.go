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

func updateSequence(sequence [4]int, newNb int) [4]int {
	sequence[0] = sequence[1]
	sequence[1] = sequence[2]
	sequence[2] = sequence[3]
	sequence[3] = newNb
	return sequence
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	diffs := make(map[[4]int]int)

	for _, line := range file {
		value, _ := strconv.Atoi(line)
		oldDiff, currDiff := 0, 0
		sequence := [4]int{0, 0, 0, 0}
		dif := make(map[[4]int]int)
		for i := 0; i < 2000; i++ {
			currDiff = value % 10
			sequence = updateSequence(sequence, currDiff - oldDiff)
			oldDiff = currDiff
			value = getNextSecret(value)
			if i > 3 {
				if _, exists := dif[sequence]; !exists {
					dif[sequence] = currDiff
				}
			}
		}
		for seq, val := range dif {
			diffs[seq] += val
		}
	}
	result := 0
	for _, value := range diffs {
		if value > result {
			result = value
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
