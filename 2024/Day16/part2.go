package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type Path struct {
	score int
	dir [2]int
	tiles [][2]int
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)

	start := [2]int{0, 0}
	end := [2]int{0, 0}
	maze := make(map[[2]int]rune)
	currDir := [2]int{0, 1}

	for i, line := range file {
		for j, char := range line {
			maze[[2]int{i, j}] = char
			if char == 'S' {
				start = [2]int{i, j}
			}
			if char == 'E' {
				end = [2]int{i, j}
			}
		}
	}

	visited := make(map[[2]int]map[[2]int]int)
	paths := []Path{{score: 0, dir: currDir, tiles: [][2]int{start}}}
	allPaths := []Path{}

	for iteration := 0; iteration < 1000; iteration++ {
		tempPaths := []Path{}
		sort.Slice(paths, func(i, j int) bool {
			return paths[i].score < paths[j].score
		})
		for _, path := range paths {
			score := path.score
			currDir := path.dir
			x, y := path.tiles[len(path.tiles) - 1][0], path.tiles[len(path.tiles) - 1][1]

			for _, delta := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nx, ny := x + delta[0], y + delta[1]
				if _, exists := maze[[2]int{nx, ny}]; !exists {
					continue
				}
				if maze[[2]int{nx, ny}] == '#' {
					continue
				}

				temp := Path{
					score: score + 1,
					dir: delta,
					tiles: append([][2]int{}, path.tiles...),
				}
				temp.tiles = append(temp.tiles, [2]int{nx, ny})

				if currDir[0] != delta[0] && currDir[1] != delta[1] {
					temp.score += 1000
				}

				if _, visitedDir := visited[[2]int{nx, ny}]; !visitedDir {
					visited[[2]int{nx, ny}] = make(map[[2]int]int)
				}
				if visitedScore, visistedExists := visited[[2]int{nx, ny}][delta]; visistedExists && visitedScore < temp.score {
					continue
				}
				visited[[2]int{nx, ny}][delta] = temp.score
				tempPaths = append(tempPaths, temp)
				if nx == end[0] && ny == end[1] {
					allPaths = append(allPaths, temp)
				}
			}
		}
		paths = tempPaths
	}
	minCost := math.MaxInt
	for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		if score, exists := visited[end][dir]; exists && score < minCost {
			minCost = score
		}
	}
	bestPaths := []Path{}
	for _, path := range allPaths {
		if path.score == minCost {
			bestPaths = append(bestPaths, path)
		}
	}
	bestTiles := make(map[[2]int]struct{})
	for _, bestPath := range bestPaths {
		for _, tile := range bestPath.tiles {
			bestTiles[tile] = struct{}{}
		}
	}
	result := len(bestTiles)
	fmt.Println(result)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
