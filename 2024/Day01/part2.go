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

func part2() int {
	lists := stringToLines(readFileContent("./input.txt"))
	leftList := []int{}
	rightList := map[int]int{}
	sumSimilarityScore := 0

	for _, numbers := range lists {
		value := 0
		parts := strings.Split(numbers, "   ")
		fmt.Sscanf(parts[0], "%d", &value)
		leftList = append(leftList, value)
		fmt.Sscanf(parts[1], "%d", &value)
		rightList[value] += 1
	}
	for i := 0; i < len(leftList); i++ {
		sumSimilarityScore += leftList[i] * rightList[leftList[i]]
	}
	return sumSimilarityScore
}

func main() {
	fmt.Println(part2())
}
