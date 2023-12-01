package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	in, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("could not read file: %v", err)
		return
	}

	val := 0
	lines := strings.Split(string(in), "\r\n")
	for i, line := range lines {
		lineDigits := getDigitsFromString(line)
		if len(lineDigits) == 0 {
			fmt.Printf("invalid line (%d): %s", i+1, line)
			continue
		}

		lineVal, err := strconv.Atoi(lineDigits[0] + lineDigits[len(lineDigits)-1])
		if err != nil {
			fmt.Printf("could not get line val: %v", err)
			return
		}

		val += lineVal
	}

	fmt.Printf("%d", val)
}

type digitIndex struct {
	index int
	digit string
}

func getDigitsFromString(str string) []string {
	var digitIndexes []digitIndex
	for _, digits := range []struct {
		searchString string
		numericDigit string
	}{
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
		{"4", "4"},
		{"5", "5"},
		{"6", "6"},
		{"7", "7"},
		{"8", "8"},
		{"9", "9"},
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	} {
		digitIndexes = append(digitIndexes, getDigitIndexes(str, digits.searchString, digits.numericDigit)...)
	}

	if len(digitIndexes) == 0 {
		return []string{}
	}

	slices.SortFunc(digitIndexes, func(a, b digitIndex) int {
		return a.index - b.index
	})

	out := make([]string, len(digitIndexes))
	for i := range digitIndexes {
		out[i] = digitIndexes[i].digit
	}

	return out
}

func getDigitIndexes(str, searchString, digit string) []digitIndex {
	var digitIndexes []digitIndex
	for {
		index := strings.Index(str, searchString)
		if index > -1 {
			digitIndexes = append(digitIndexes, digitIndex{
				index: index,
				digit: digit,
			})
			str = strings.Replace(str, searchString, strings.Repeat("x", len(searchString)), 1)
			continue
		}

		break
	}

	return digitIndexes
}
