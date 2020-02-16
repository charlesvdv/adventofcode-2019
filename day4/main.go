package main

import (
	"fmt"
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

func hasSuccessiveDigits(digits []int, numberOfSuccessions int) bool {
	lastDigit := -1
	successions := 0

	for _, digit := range digits {
		if lastDigit == digit {
			successions += 1
		} else {
			successions = 0
		}
		lastDigit = digit

		if successions == (numberOfSuccessions - 1) {
			return true
		}
	}

	return false
}

func hasExactSuccessiveDigits(digits []int, numberOfSuccessions int) bool {
	lastDigit := -1
	successions := 0

	for index, digit := range digits {
		if lastDigit == digit {
			successions += 1
		} else {
			successions = 0
		}
		lastDigit = digit

		if successions == (numberOfSuccessions - 1) {
			if index == len(digits)-1 {
				return true
			}
			if digit != digits[index+1] {
				return true
			}
		}
	}

	return false
}

func neverDecrease(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i-1] > digits[i] {
			return false
		}
	}
	return true
}

func main() {
	lowerBound := 256310
	higherBound := 732736

	validPasswords := 0
	for number := lowerBound; number <= higherBound; number++ {
		digits := numberToDigits(number)
		if hasSuccessiveDigits(digits, 2) && neverDecrease(digits) {
			validPasswords += 1
		}
	}

	fmt.Printf("Part 1: %v\n", validPasswords)

	validPasswords2 := 0
	for number := lowerBound; number <= higherBound; number++ {
		digits := numberToDigits(number)
		if hasExactSuccessiveDigits(digits, 2) && neverDecrease(digits) {
			validPasswords2 += 1
		}
	}

	fmt.Printf("Part 2: %v\n", validPasswords2)
}
