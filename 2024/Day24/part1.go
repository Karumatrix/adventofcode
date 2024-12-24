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

func reverseStr(str string) (string) {
	bytes := []byte(str)
	for i, j := 0, len(bytes) - 1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
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
	for len(file) > 0 {
		newFile := []string{}
		for _, line := range file {
			parts := strings.Split(line, " ")
			a, b, dest := parts[0], parts[2], parts[4]
			op := parts[1]

			_, existsA := wires[a]
			_, existsB := wires[b]
			if existsA && existsB {
				switch op {
				case "AND":
					wires[dest] = wires[a] & wires[b]
				case "OR":
					wires[dest] = wires[a] | wires[b]
				case "XOR":
					wires[dest] = wires[a] ^ wires[b]
				}
			} else {
				newFile = append(newFile, line)
			}
		}
		file = newFile
	}
	keys := make([]string, 0, len(wires))
	for key := range wires {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	binary := ""
	for _, key := range keys {
		if key[0] == 'z' {
			binary += strconv.Itoa(int(wires[key]))
		}
	}
	binary = reverseStr(binary)
	result, _ := strconv.ParseInt(binary, 2, 64)
	fmt.Println(result)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
