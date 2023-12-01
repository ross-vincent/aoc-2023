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
		lineDigits := []string{}
		str := ""
		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				lineDigits = append(lineDigits, getDigitsFromString(str)...)
				lineDigits = append(lineDigits, string(char))
				str = ""
				continue
			}

			str += string(char)
		}
		lineDigits = append(lineDigits, getDigitsFromString(str)...)

		if len(lineDigits) == 0 {
			fmt.Printf("invalid line (%d): %s", i+1, line)
			continue
		}

		// fmt.Printf("line %d digits: %+v\n", i+1, lineDigits)

		lineVal, err := strconv.Atoi(lineDigits[0] + lineDigits[len(lineDigits)-1])
		if err != nil {
			fmt.Printf("could not get line val: %v", err)
			return
		}

		val += lineVal
	}

	fmt.Printf("%d", val)
}

type stringIndex struct {
	index int
	val   string
}

func getDigitsFromString(str string) []string {
	var digits []stringIndex
	for _, digitStr := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		digits = append(digits, getIndexesOfStrings(str, digitStr)...)
	}

	if len(digits) == 0 {
		return []string{}
	}

	slices.SortFunc(digits, func(a, b stringIndex) int {
		return a.index - b.index
	})

	out := make([]string, len(digits))
	for i := range digits {
		out[i] = getNumericDigit(digits[i].val)
	}

	return out
}

func getIndexesOfStrings(str string, searchString string) []stringIndex {
	var stringIndexes []stringIndex
	haystack := strings.Clone(str)
	for {
		index := strings.Index(haystack, searchString)
		if index > -1 {
			stringIndexes = append(stringIndexes, stringIndex{
				index: index,
				val:   searchString,
			})
			haystack = strings.Replace(haystack, searchString, strings.Repeat("x", len(searchString)), 1)
			continue
		}

		break
	}

	return stringIndexes
}

func getNumericDigit(val string) string {
	switch val {

	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	}

	return ""
}
