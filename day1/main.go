package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func parseInput(reader io.Reader) []int {
	result := []int{}

	scanner := bufio.NewScanner(reader)	
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, val)
	}

	return result
}

func calculateFuelRequired(mass int) int {
	return int(mass / 3) - 2
}

func calculateFuelRequiredForFuel(mass int) int {
	totalFuelForFuel := 0
	lastFuelRequired := mass
	for true {
		lastFuelRequired = calculateFuelRequired(lastFuelRequired)	
		if lastFuelRequired <= 0 {
			break
		}
		totalFuelForFuel = totalFuelForFuel + lastFuelRequired
	}
	return totalFuelForFuel
}

func main() {
	masses := parseInput(bufio.NewReader(os.Stdin))

	totalFuel := 0
	for _, mass := range masses {
		moduleFuel := calculateFuelRequired(mass)
		totalFuel = totalFuel + moduleFuel
	}

	fmt.Printf("Part 1: %v\n", totalFuel)

	totalFuel = 0
	for _, mass := range masses {
		moduleOnlyFuel := calculateFuelRequired(mass)
		fuelForFuel := calculateFuelRequiredForFuel(moduleOnlyFuel)
		totalFuel = totalFuel + moduleOnlyFuel + fuelForFuel
	}

	fmt.Printf("Part 2: %v\n", totalFuel)
}