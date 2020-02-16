package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"
)

type command struct {
	direction string
	distance  int
}

func parseInput(reader io.Reader) [][]command {
	scanner := bufio.NewScanner(reader)

	result := [][]command{}
	for scanner.Scan() {
		result = append(result, parseLine(scanner.Text()))
	}

	return result
}

func parseLine(line string) []command {
	result := []command{}

	rawCommands := strings.Split(line, ",")
	for _, rawCommand := range rawCommands {
		rawCommandRune := []rune(rawCommand)
		direction := string(rawCommandRune[0])
		distance, err := strconv.Atoi(string(rawCommandRune[1:]))
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, command{
			direction: direction,
			distance:  distance,
		})
	}

	return result
}

type coord struct {
	x int
	y int
}

type WireMap = map[coord]map[int]int

func buildWireMap(wires [][]command) WireMap {
	wireMap := WireMap{}
	for wireIndex, wireCommands := range wires {
		buildWireMapForOneWire(wireMap, wireIndex, wireCommands)
	}
	return wireMap
}

func buildWireMapForOneWire(wireMap WireMap, wireIndex int, commands []command) {
	currentCoord := coord{x: 0, y: 0}
	distanceDone := 0
	for _, command := range commands {
		for ; command.distance > 0; command.distance-- {
			switch command.direction {
			case LEFT:
				currentCoord.x -= 1
			case RIGHT:
				currentCoord.x += 1
			case UP:
				currentCoord.y += 1
			case DOWN:
				currentCoord.y -= 1
			default:
				log.Fatal("unknown command")
			}
			distanceDone += 1

			cell, ok := wireMap[currentCoord]
			if !ok {
				cell = map[int]int{}
			}
			cell[wireIndex] = distanceDone
			wireMap[currentCoord] = cell
		}
	}
}

func findClosestCrossPathManhattan(wireMap WireMap) int {
	bestDistance := math.MaxInt32
	for coord := range wireMap {
		if wireMap[coord][0] == 0 || wireMap[coord][1] == 0 {
			// we don't have an intersection
			continue
		}

		distanceToCenter := manhattanDistanceToCenter(coord)
		if distanceToCenter < bestDistance {
			bestDistance = distanceToCenter
		}
	}
	return bestDistance
}

func manhattanDistanceToCenter(coord coord) int {
	return intAbs(coord.x) + intAbs(coord.y)
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func findClosestCrossPathSignalDelay(wireMap WireMap) int {
	bestDistance := math.MaxInt32
	for coord := range wireMap {
		if wireMap[coord][0] == 0 || wireMap[coord][1] == 0 {
			// we don't have an intersection
			continue
		}

		calculatedDistance := wireMap[coord][0] + wireMap[coord][1]
		if calculatedDistance < bestDistance {
			bestDistance = calculatedDistance
		}
	}
	return bestDistance
}

func main() {
	wires := parseInput(bufio.NewReader(os.Stdin))
	if len(wires) != 2 {
		log.Fatal("invalid input")
	}
	wireMap := buildWireMap(wires)

	bestDistance := findClosestCrossPathManhattan(wireMap)
	fmt.Printf("Part 1: %v\n", bestDistance)

	bestDistanceSignal := findClosestCrossPathSignalDelay(wireMap)
	fmt.Printf("Part 2: %v\n", bestDistanceSignal)
}
