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

func isStoneEven(stone int) bool {
	return (len(strconv.Itoa(stone)) % 2) == 0
}

func aoc(filepath string) {
	file := filepathToString(filepath)
	parts := strings.Split(file, " ")
	stones := make(map[int]int)

	for _, part := range parts {
		stone, _ := strconv.Atoi(part)
		stones[stone]++
	}
	for i := 0; i < 75; i++ {
		newStones := make(map[int]int)
		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
			} else if isStoneEven(stone) {
				rockStr := strconv.Itoa(stone)
				mid := len(rockStr) / 2
				left, _ := strconv.Atoi(rockStr[:mid])
				right, _ := strconv.Atoi(rockStr[mid:])
				newStones[left] += count
				newStones[right] += count
			} else {
				newStones[stone * 2024] += count
			}
		}
		stones = newStones
	}

	result := 0
	for _, count := range stones {
		result += count
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
