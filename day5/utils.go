package main

import (
	"log"
	"strconv"
)

func numberToDigits(number int) []int {
	result := []int{}
	digitList := []rune(strconv.Itoa(number))
	for _, digitRune := range digitList {
		digit, err := strconv.Atoi(string(digitRune))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, digit)
	}

	return result
}

func digitsToNumber(digits []int) int {
	rawDigits := ""
	for _, digit := range digits {
		rawDigits = rawDigits + strconv.Itoa(digit)
	}
	number, err := strconv.Atoi(rawDigits)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
