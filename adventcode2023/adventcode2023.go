package adventcode2023

import (
	"errors"

	"alex952.com/advent2023/adventcode2023/day1"
)

type AdventOfCodeChallengRunner interface {
	GetInputFilename() string
	SetInputFilename(filename string)
	RunChallenge(first bool) (string, error)
}

var DayChallengesRunners = map[int]AdventOfCodeChallengRunner{
	1: &day1.Day1Runner{Input: "./adventcode2023/day1/input.txt"},
}

func MakeChallengeRunner(day int) (AdventOfCodeChallengRunner, error) {
	if runner, exists := DayChallengesRunners[day]; exists {
		return runner, nil
	} else {
		return nil, errors.New("Day not available!")
	}
}
