package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(reader io.Reader) []int {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	result := []int{}
	strContent := strings.TrimSpace(string(content))
	for _, rawNumber := range strings.Split(strContent, ",") {
		num, err := strconv.Atoi(rawNumber)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, num)
	}

	return result
}

func executeProgram(instructions []int) error {
	pc := 0 // program counter

	for pc < len(instructions) {
		switch instructions[pc] {
		case 1:
			opAdd(instructions, instructions[pc+1], instructions[pc+2], instructions[pc+3])
			pc = pc + 4
		case 2:
			opMultiply(instructions, instructions[pc+1], instructions[pc+2], instructions[pc+3])
			pc = pc + 4
		case 99:
			return nil
		default:
			return errors.New("unexpected op code")
		}
	}

	return errors.New("unexpected end of program")
}

func opAdd(instructions []int, lhs, rhs, output int) {
	instructions[output] = instructions[lhs] + instructions[rhs]
}

func opMultiply(instructions []int, lhs, rhs, output int) {
	instructions[output] = instructions[lhs] * instructions[rhs]
}

func bruteForceResult(instructions []int, expectedValue int) (int, int) {
	for nounIndex := 0; nounIndex < 100; nounIndex++ {
		for verbIndex := 0; verbIndex < 100; verbIndex++ {
			instructionsCopy := append([]int{}, instructions...)
			instructionsCopy[1] = nounIndex
			instructionsCopy[2] = verbIndex

			err := executeProgram(instructionsCopy)
			if err == nil && instructionsCopy[0] == expectedValue {
				return nounIndex, verbIndex
			}
		}
	}

	return -1, -1
}

func main() {
	instructions := parseInput(bufio.NewReader(os.Stdin))

	patchedInstructions := append([]int{}, instructions...)
	patchedInstructions[1] = 12
	patchedInstructions[2] = 2

	err := executeProgram(patchedInstructions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", patchedInstructions[0])

	noun, verb := bruteForceResult(instructions, 19690720)
	fmt.Printf("Part 2: %v\n", (100*noun)+verb)
}
