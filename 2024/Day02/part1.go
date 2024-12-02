package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFileContent(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func stringToLines(content string) []string {
    var lines []string
    scanner := bufio.NewScanner(strings.NewReader(content))

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}

func safeIncreasing(numbers []int) bool {
	for i := 0; i < len(numbers) - 1; i++ {
		diff := numbers[i + 1] - numbers[i]
		if diff < 1 || diff > 3 {
			return false
		}
		if numbers[i] > numbers[i + 1] || numbers[i] == numbers[i + 1] {
			return false
		}
	}
	return true
}

func safeDecreasing(numbers []int) bool {
	for i := 0; i < len(numbers) - 1; i++ {
		diff := numbers[i] - numbers[i + 1]
		if diff < 1 || diff > 3 {
			return false
		}
		if numbers[i] < numbers[i + 1] || numbers[i] == numbers[i + 1] {
			return false
		}
	}
	return true
}

func part1() int {
	reports := stringToLines(readFileContent("./input.txt"))
	result := 0

	for _, levels := range reports {
		strNumbers := strings.Split(levels, " ")
		var intNumbers []int
		for _, number := range strNumbers {
			value := 0
			fmt.Sscanf(number, "%d", &value)
			intNumbers = append(intNumbers, value)
		}
		if safeIncreasing(intNumbers) == true {
			result += 1
		}
		if safeDecreasing(intNumbers) == true {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(part1())
}
