package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id     int
	rounds []round
}

func (g *game) minCubeNumbers() cubeNumbers {
	minNumbers := cubeNumbers{}
	for _, round := range g.rounds {
		if round.cubeNumbers.red > minNumbers.red {
			minNumbers.red = round.cubeNumbers.red
		}
		if round.cubeNumbers.blue > minNumbers.blue {
			minNumbers.blue = round.cubeNumbers.blue
		}
		if round.cubeNumbers.green > minNumbers.green {
			minNumbers.green = round.cubeNumbers.green
		}
	}

	return minNumbers
}

type round struct {
	num         int
	cubeNumbers cubeNumbers
}

func (r *round) validRound(maxCubeNumbers cubeNumbers) bool {
	if r.cubeNumbers.red > maxCubeNumbers.red {
		return false
	}

	if r.cubeNumbers.green > maxCubeNumbers.green {
		return false
	}

	if r.cubeNumbers.blue > maxCubeNumbers.blue {
		return false
	}

	return true
}

type cubeNumbers struct {
	red   int
	green int
	blue  int
}

var (
	gameRegex  = regexp.MustCompile(`^Game (\d+)$`)
	roundRegex = regexp.MustCompile(`^(?: (\d+) (red|green|blue),)?(?: (\d+) (red|green|blue),)? (\d+) (red|green|blue)$`)
)

func main() {
	in, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("could not read file: %v", err)
		return
	}

	lines := strings.Split(string(in), "\r\n")

	val, err := part1(lines)
	if err != nil {
		fmt.Printf("error in part 1: %v", err)
		return
	}

	fmt.Printf("part 1: %d\n", val)

	val2, err := part2(lines)
	if err != nil {
		fmt.Printf("error in part 2: %v", err)
		return
	}

	fmt.Printf("part 2: %d", val2)
}

func part1(lines []string) (int, error) {
	maxCubeNumbers := cubeNumbers{
		red:   12,
		green: 13,
		blue:  14,
	}

	val := 0
	for i, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return 0, fmt.Errorf("invalid game line (%d): %v", i+1, err)
		}

		validGame := true
		for _, round := range game.rounds {
			if !round.validRound(maxCubeNumbers) {
				validGame = false
			}
		}

		if validGame {
			val += game.id
		}
	}

	return val, nil
}

func part2(lines []string) (int, error) {
	val := 0

	for i, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return 0, fmt.Errorf("invalid game line (%d): %v", i+1, err)
		}

		minCubeNumbers := game.minCubeNumbers()
		gamePower := minCubeNumbers.red * minCubeNumbers.green * minCubeNumbers.blue
		val += gamePower
	}

	return val, nil
}

func parseGame(line string) (*game, error) {
	parts := strings.Split(line, ":")

	match := gameRegex.FindStringSubmatch(parts[0])
	if match == nil {
		return nil, errors.New("invalid game format")
	}

	id, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}

	game := game{
		id: id,
	}

	roundStrs := strings.Split(parts[1], ";")
	for i, roundStr := range roundStrs {
		round := round{
			num: i + 1,
		}

		roundVals := roundRegex.FindStringSubmatch(roundStr)
		if roundVals == nil {
			return nil, errors.New("invalid round format")
		}

		currentNum := 0
		for i, val := range roundVals {
			if i == 0 || val == "" {
				continue
			}

			if i%2 == 1 {
				num, err := strconv.Atoi(val)
				if err != nil {
					return nil, err
				}

				currentNum = num
				continue
			}

			switch val {

			case "red":
				round.cubeNumbers.red = currentNum
			case "green":
				round.cubeNumbers.green = currentNum
			case "blue":
				round.cubeNumbers.blue = currentNum

			}
		}

		game.rounds = append(game.rounds, round)
	}

	return &game, nil
}
