package main

import (
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

func parseInput(input string) (rules, updates [][]int) {
	parts := strings.Split(input, "\n\n")
	part1 := strings.Split(parts[0], "\n")
	part2 := strings.Split(parts[1], "\n")
	for _, line := range part1 {
		value1, value2 := 0, 0
		fmt.Sscanf(line, "%d|%d", &value1, &value2)
		rules = append(rules, []int{value1, value2})
	}
	for _, line := range part2 {
		pages := strings.Split(line, ",")
		var update []int
		for _, page := range pages {
			value := 0
			fmt.Sscanf(page, "%d", &value)
			update = append(update, value)
		}
		updates = append(updates, update)
	}
	return rules, updates
}

func isValidOrder(update []int, rules map[int][]int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}
	
	for index, rule := range rules {
		posX, exist := position[index];
		if exist {
			for _, page := range rule {
				posY, existY := position[page];
				if existY && posX > posY {
					return false
				}
			}
		}
	}
	
	return true
}

func correctOrder(update []int, rules map[int][]int) (result []int) {
	pagesPriority := make(map[int]int)
	prevPages := make(map[int][]int)
	queue := []int{}
	for _, page := range update {
		pagesPriority[page] = 0
		prevPages[page] = []int{}
	}
	for index, rule := range rules {
		_, ok := pagesPriority[index];
		if ok {
			for _, page := range rule {
				_, ok := pagesPriority[page];
				if ok {
					prevPages[index] = append(prevPages[index], page)
					pagesPriority[page]++
				}
			}
		}
	}
	for page, index := range pagesPriority {
		if index == 0 {
			queue = append(queue, page)
			break
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		for _, neighbor := range prevPages[current] {
			pagesPriority[neighbor]--
			if pagesPriority[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

func main() {
	input := readFileContent("./input.txt")
	rules, updates := parseInput(input)
	ruleMap := make(map[int][]int)
	result := 0
	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}
	for _, update := range updates {
		if !isValidOrder(update, ruleMap) {
			correctedUpdate := correctOrder(update, ruleMap)
			result += correctedUpdate[len(correctedUpdate)/2]
		}
	}
	fmt.Println(result)
}
