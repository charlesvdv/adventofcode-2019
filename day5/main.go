package main

import (
	"bufio"
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

func main() {
	program := parseInput(bufio.NewReader(os.Stdin))

	processor := NewProcessor(program)
	output, err := processor.Run(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", output)

	output, err = processor.Run(5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 2: %v\n", output)
}
