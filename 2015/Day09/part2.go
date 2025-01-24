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

func combineCities(cities []string) (result [][]string) {
	var permute func(int)
	permute = func(i int) {
		if i == len(cities) - 1 {
			tmp := append([]string{}, cities...)
			result = append(result, tmp)
			return
		}
		for j := i; j < len(cities); j++ {
			cities[i], cities[j] = cities[j], cities[i]
			permute(i + 1)
			cities[i], cities[j] = cities[j], cities[i]
		}
	}
	permute(0)
	return
}

func safeAppend(cities []string, city string) []string {
	for _, line := range cities {
		if line == city {
			return cities
		}
	}
	cities = append(cities, city)
	return cities
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	cities := []string{}
	distancesMap := make(map[string]map[string]int)

	for _, line := range file {
		parts := strings.Split(line, " ")
		if _, exists := distancesMap[parts[0]]; !exists {
			distancesMap[parts[0]] = make(map[string]int)
		}
		distancesMap[parts[0]][parts[2]], _ = strconv.Atoi(parts[4])
		cities = safeAppend(cities, parts[0])
		cities = safeAppend(cities, parts[2])
	}
	citiesCombined := combineCities(cities)
	result := 0
	for _, comb := range citiesCombined {
		currDistance := 0
		for i := 0; i < len(comb) - 1; i++ {
			currDistance += distancesMap[comb[i]][comb[i + 1]]
			currDistance += distancesMap[comb[i + 1]][comb[i]]
		}
		if currDistance > result {
			result = currDistance
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
