package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Tree struct {
	Name     string
	Children []Tree
}

func parseInput(input io.Reader) map[string][]string {
	result := map[string][]string{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		rawInput := strings.Split(line, ")")
		parent, child := rawInput[0], rawInput[1]

		if _, ok := result[parent]; !ok {
			result[parent] = []string{}
		}
		result[parent] = append(result[parent], child)
	}

	return result
}

func getTotalDepthRecursive(pairs map[string][]string, parent string, depth int) int {
	recursiveResult := 0
	for _, child := range pairs[parent] {
		recursiveResult += getTotalDepthRecursive(pairs, child, depth+1)
	}
	return depth + recursiveResult
}

func getTotalDepth(pairs map[string][]string) int {
	return getTotalDepthRecursive(pairs, "COM", 0)
}

func getOrbitPathRecursive(pairs map[string][]string, target string, currentPath []string) []string {
	currentPos := currentPath[len(currentPath)-1]
	// exit condition
	if len(pairs[currentPos]) == 0 {
		if currentPos == target {
			return currentPath
		} else {
			return nil
		}
	}

	// recursion
	for _, child := range pairs[currentPos] {
		if possiblePath := getOrbitPathRecursive(pairs, target, append(currentPath, child)); possiblePath != nil {
			return possiblePath
		}
	}
	return nil
}

func getOrbitPath(pairs map[string][]string, target string) []string {
	return getOrbitPathRecursive(pairs, target, []string{"COM"})
}

func getShortestPath(pairs map[string][]string, start, finish string) int {
	startOrbitPath := getOrbitPath(pairs, start)
	endOrbitPath := getOrbitPath(pairs, finish)

	minSize := minInt(len(startOrbitPath), len(endOrbitPath))
	commonPartIndex := 0
	for ; commonPartIndex < minSize; commonPartIndex++ {
		if startOrbitPath[commonPartIndex] != endOrbitPath[commonPartIndex] {
			break
		}
	}

	// -2 to remove the start and finish value that should not be taken into account
	return len(startOrbitPath) + len(endOrbitPath) - 2*commonPartIndex - 2
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	pairs := parseInput(bufio.NewReader(os.Stdin))

	fmt.Printf("Part 1: %v\n", getTotalDepth(pairs))
	fmt.Printf("Part 2: %v\n", getShortestPath(pairs, "YOU", "SAN"))
}
