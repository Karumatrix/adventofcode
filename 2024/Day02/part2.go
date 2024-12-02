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

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		newSlice := make([]int, len(slice))
		copy(newSlice, slice)
		return newSlice
	}
	return append(append([]int{}, slice[:i]...), slice[i+1:]...)
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
		if safeIncreasing(intNumbers) == true || safeDecreasing(intNumbers) == true {
			result += 1
		} else {
			for i := 0; i < len(intNumbers); i++ {
				numbers := removeElement(intNumbers, i)
				if safeIncreasing(numbers) == true {
					result += 1
					break
				}
				if safeDecreasing(numbers) == true {
					result += 1
					break
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(part1())
}
