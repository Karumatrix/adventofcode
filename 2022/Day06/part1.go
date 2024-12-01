package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readFileContent(filepath string) string {
    content, err := ioutil.ReadFile(filepath)
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

func part1() int {
    buffer := readFileContent("input.txt")

    for i := 3; i < len(buffer); i++ {
        chars := buffer[i - 3 : i + 1]
        unique := true
        for j, c1 := range chars {
            for _, c2 := range chars[j+1:] {
                if c1 == c2 {
                    unique = false
                    break
                }
            }
            if !unique {
                break
            }
        }
        if unique {
            return i + 1
        }
    }
    return 0
}

func main() {
    fmt.Println(part1())
}
